package router

import (
	"go-server/internal/controller"
	"go-server/internal/dao"
	"go-server/internal/middleware"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func InitGoodsRouter(deps RouterDeps, r *gin.RouterGroup) {
	// ================= 商品模块 =================
	// 初始化依赖
	goodsRepository := dao.NewGoodsRepository(deps.Repository)                   // dao
	goodsService := service.NewGoodsService(deps.Service, goodsRepository)       // service
	goodsController := controller.NewGoodsController(deps.Handler, goodsService) // controller

	user := r.Group("/goods")
	// ✅ 需要登录
	strictAuthRouter := user.Group("").Use(middleware.StrictAuth(deps.JWT, deps.Logger))
	{
		user.POST("", goodsController.Create)
		strictAuthRouter.GET("", goodsController.GetLists)      // 当前toekn信息
		strictAuthRouter.PUT("/:id", goodsController.Update)    // 修改
		strictAuthRouter.DELETE("/:id", goodsController.Delete) // 删除
		strictAuthRouter.GET("/:id", goodsController.GetDetail) // 详情
	}
}
