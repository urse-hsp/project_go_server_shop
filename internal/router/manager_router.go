package router

import (
	"go-server/internal/controller"
	"go-server/internal/dao"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func InitManagerRouter(deps RouterDeps, r *gin.RouterGroup) {
	// ================= 管理员模块 =================

	managerRepository := dao.NewManagerRepository(deps.Repository)               // dao
	managerService := service.NewManagerService(deps.Service, managerRepository) // service
	managerController := controller.NewManagerController(managerService)         // controller
	noAuthRouter := r.Group("/")
	{
		// ✅ 不需要登录
		noAuthRouter.POST("login", managerController.Login)
	}
	// ✅ 需要登录
	StrictAuthRouter := noAuthRouter.Group("/users") //.Use(middleware.StrictAuth(deps.JWT, deps.Logger))
	{
		StrictAuthRouter.POST("", managerController.Create)
		StrictAuthRouter.GET("", managerController.GetLists)
		StrictAuthRouter.DELETE("/:id", managerController.Delete)
		StrictAuthRouter.PUT("/:id", managerController.Update)
	}
}
