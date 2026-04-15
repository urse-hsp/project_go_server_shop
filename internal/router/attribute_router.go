package router

import (
	"go-server/internal/controller"
	"go-server/internal/dao"
	"go-server/internal/middleware"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func InitGoodsAttrRouter(deps RouterDeps, r *gin.RouterGroup) {
	// ================= 分类参数模块 =================
	// 初始化依赖
	Repository := dao.NewAttributeRepository(deps.Repository)              // dao
	Service := service.NewAttributeService(deps.Service, Repository)       // service
	Controller := controller.NewAttributeController(deps.Handler, Service) // controller

	noAuthRouter := r.Group("/:id/attributes")
	// ✅ 不需要登录
	{
	}
	// ✅ 需要登录
	strictAuthRouter := noAuthRouter.Group("").Use(middleware.StrictAuth(deps.JWT, deps.Logger))
	{
		strictAuthRouter.POST("", Controller.Create)           // create
		strictAuthRouter.GET("", Controller.GetList)           // get
		strictAuthRouter.PUT("/:attrId", Controller.Update)    // edit
		strictAuthRouter.DELETE("/:attrId", Controller.Delete) // delete
	}
	// // ✅ 不强制登录
	// noStrictAuth := noAuthRouter.Group("").Use(middleware.NoStrictAuth(deps.JWT, deps.Logger))
	// {
	// 	noStrictAuth.GET("/lists", Controller.GetPageList) // 分页列表
	// }
}
