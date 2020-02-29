package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"user-service/api/user"
)

type UserService struct {
	user.UserInfo
}


func NewUserInfo() *UserService {
	return &UserService{}
}

func (this *UserService) GetUserById(ctx context.Context, id int32) (r *user.UserInfo, err error) {
	db := MysqlDefault().GetConn()
	defer db.Close()
	userinfo := &user.UserInfo{}
	if err := db.Where("id = ?", id).Find(userinfo).Error; err != nil {
		return nil,fmt.Errorf("系统错误")
	}
	return userinfo, nil
}
// Parameters:
//  - ID
func (this *UserService) GetTeacherById(ctx context.Context, id int32) (r *user.UserInfo, err error) {
	db := MysqlDefault().GetConn()
	defer db.Close()
	userinfo := &user.UserInfo{}
	if err := db.Where("id = ?", id).Find(userinfo).Error; err != nil {
		return nil,fmt.Errorf("系统错误")
	}
	return userinfo, nil
}
// Parameters:
//  - Username
func (this *UserService)  GetUserByName(ctx context.Context, username string) (r *user.UserInfo, err error) {
	db := MysqlDefault().GetConn()
	defer db.Close()
	userinfo := &user.UserInfo{}
	if err := db.Where("username = ?", username).Find(userinfo).Error; err != nil {
		return nil,fmt.Errorf("系统错误")
	}
	return userinfo, nil
}
// Parameters:
//  - UserInfo
func (this *UserService)  RegiserUser(ctx context.Context, userInfo *user.UserInfo) (err error) {
	db := MysqlDefault().GetConn()
	defer db.Close()
	err = db.Model(user.UserInfo{}).Create(userInfo).Error
	if err != nil {
		return fmt.Errorf("系统错误")
	}
	return nil
}

func main() {

	transport,err := thrift.NewTServerSocket(":9001")
	if err != nil {
		panic(err)
	}
	handler := &UserService{}
	processor := user.NewUserServiceProcessor(handler)
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(processor, transport,transportFactory,protocolFactory)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}

