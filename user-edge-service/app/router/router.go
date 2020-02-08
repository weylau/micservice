package router

import (
	"github.com/gin-gonic/gin"
	"user-edge-service/app/controller/admin"
	"user-edge-service/app/middleware"
	"user-edge-service/app/protocol"
	"net/http"
)

type Router struct {
	engine *gin.Engine
}

func Default() *Router {
	router := &Router{}
	router.engine = gin.Default()
	return router
}

func (this *Router) Run() {
	this.SetCors()
	this.setFront()
	this.setAdmin()
	this.set404()
}

func (this *Router) GetEngin() *gin.Engine {
	return this.engine
}

func (this *Router) SetCors() {
	this.engine.Use(middleware.Cors())
}

func (this *Router) setFront() {

}

func (this *Router) setAdmin() {
	//后台管理
	login_admin_ctrl := admin.Login{}
	user_admin_ctrl := admin.User{}
	this.engine.POST("/adapi/login", login_admin_ctrl.Login)
	authorized := this.engine.Group("/adapi")
	authorized.Use(middleware.CheckAuth())
	{
		authorized.GET("/user", user_admin_ctrl.Show)
	}
}

func (this *Router) set404() {
	this.engine.NoRoute(func(context *gin.Context) {
		resp := protocol.Resp{Ret: 404, Msg: "page not exists!", Data: ""}
		//返回404状态码
		context.JSON(http.StatusNotFound, resp)
	})
}
