package admin

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	user_service "user-edge-service/app/component/user-service"
	"user-edge-service/app/config"
	"user-edge-service/app/helper"
	"user-edge-service/app/loger"
	"user-edge-service/app/protocol"
	"time"
)

type Admins struct {
}

func (*Admins) getLogTitle() string {
	return "service-admin-login-"
}

//用户信息
type ReturnUserInfo struct {
	AdminId  int32  `json:"admin_id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

//登录
func (this *Admins) Login(username string, password string, code uint32) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	//校验谷歌验证码
	ga_code, err := helper.MkGaCode(config.Configs.GaSecret)
	if err != nil {
		loger.Default().Error(this.getLogTitle(), "Login-error-MkGaCode:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}

	if code != ga_code {
		resp.Msg = "谷歌验证码错误"
		return resp
	}

	userInfo,err := user_service.GetUserByName(username)
	if err != nil {
		loger.Default().Error(this.getLogTitle(), "Login-error-GetUserByName:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	if userInfo == nil {
		resp.Msg = "用户不存在"
		return resp
	}

	//检测密码是否正确
	if helper.MkMd5(password) != userInfo.Password {
		loger.Default().Info(this.getLogTitle(), "username:", username)
		loger.Default().Info(this.getLogTitle(), "服务端密码:", userInfo.Password)
		loger.Default().Info(this.getLogTitle(), "客户端密码:", helper.MkMd5(password))
		resp.Msg = "密码错误"
		return resp
	}

	//生成token
	token, err := helper.JwtEncode(jwt.MapClaims{"admin_id": fmt.Sprintf("%d", userInfo.ID), "username": userInfo.Username, "expr_time": fmt.Sprintf("%d", time.Now().Unix())}, []byte(config.Configs.JwtSecret))
	if err != nil {
		loger.Default().Error(this.getLogTitle(), "Login-error2:", err.Error())
		resp.Ret = -999
		resp.Msg = "系统错误"
		return resp
	}
	user_info := ReturnUserInfo{
		AdminId:  userInfo.ID,
		UserName: userInfo.Username,
		Token:    token,
	}
	resp.Data = user_info
	resp.Ret = 0
	return resp
}
