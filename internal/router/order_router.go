package router

import (
	"go-server/internal/controller"
	"go-server/internal/dao"
	"go-server/internal/middleware"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func InitOrderRouter(deps RouterDeps, r *gin.RouterGroup) {
	// ================= 订单模块 =================
	// 初始化依赖
	Repository := dao.NewOrderRepository(deps.Repository)              // dao
	Service := service.NewOrderService(deps.Service, Repository)       // service
	Controller := controller.NewOrderController(deps.Handler, Service) // controller

	noAuthRouter := r.Group("orders")
	// ✅ 需要登录
	strictAuthRouter := noAuthRouter.Group("").Use(middleware.StrictAuth(deps.JWT, deps.Logger))
	{
		// strictAuthRouter.POST("", Controller.Create)       // create
		strictAuthRouter.GET("", Controller.GetPageList) // get
		// strictAuthRouter.PUT("/:id", Controller.Update)    // edit
		// strictAuthRouter.DELETE("/:id", Controller.Delete) // delete
		strictAuthRouter.GET("/:id", Controller.GetDetail) // detail
	}
}
