package router

import (
	"go-server/internal/controller"
	"go-server/internal/dao"
	"go-server/internal/middleware"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRoleRouter(deps RouterDeps, r *gin.RouterGroup) {
	// ================= 角色模块 =================

	// r.GET("/roles", func(c *gin.Context) {
	// 	c.String(200, "123")
	// })

	// 初始化依赖
	roleRepository := dao.NewRoleRepository(deps.Repository)                  // dao
	roleService := service.NewRoleService(deps.Service, roleRepository)       // service
	roleController := controller.NewColeController(deps.Handler, roleService) // controller
	noAuthRouter := r.Group("/")

	{
		// ✅ 需要登录
		StrictAuthRouter := noAuthRouter.Group("roles").Use(middleware.StrictAuth(deps.JWT, deps.Logger))
		{
			StrictAuthRouter.POST("", roleController.Create)
			StrictAuthRouter.GET("", roleController.GetList)
			StrictAuthRouter.PUT("/:id", roleController.Update)
			StrictAuthRouter.DELETE("/:id", roleController.Delete)
		}
	}
}
