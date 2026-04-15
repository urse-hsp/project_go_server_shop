package router

import (
	"go-server/internal/controller"
	"go-server/internal/dao"
	"go-server/internal/middleware"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRightsRouter(deps RouterDeps, r *gin.RouterGroup) {
	// ================= 权限模块 =================

	rightsRepository := dao.NewRightsRepository(deps.Repository)                    // dao
	rightsService := service.NewRightsService(deps.Service, rightsRepository)       // service
	rightsController := controller.NewRightsController(deps.Handler, rightsService) // controller

	user := r.Group("/rights")

	strictAuthRouter := user.Group("").Use(middleware.StrictAuth(deps.JWT, deps.Logger))
	{
		strictAuthRouter.GET("", rightsController.GetList) // 当前toekn信息
	}
}
