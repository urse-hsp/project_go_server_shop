package router

import (
	"go-server/internal/controller"
	"go-server/internal/middleware"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func InitUploadRouter(deps RouterDeps, r *gin.RouterGroup) {
	// ================= **模块 =================
	// 初始化依赖
	Service := service.NewUploadService(deps.Service)                   // service
	Controller := controller.NewUploadController(deps.Handler, Service) // controller

	noAuthRouter := r.Group("/upload")
	// ✅ 需要登录
	strictAuthRouter := noAuthRouter.Group("").Use(middleware.StrictAuth(deps.JWT, deps.Logger))
	{
		strictAuthRouter.POST("", Controller.Upload) // create
	}
}
