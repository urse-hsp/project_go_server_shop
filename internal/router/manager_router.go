package router

import (
	"go-server/internal/controller"
	"go-server/internal/dao"
	"go-server/internal/middleware"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func InitManagerRouter(deps RouterDeps, r *gin.RouterGroup) {
	// ================= 管理员模块 =================
	manager := r.Group("/")

	managerRepository := dao.NewManagerRepository(deps.Repository)               // dao
	managerService := service.NewManagerService(deps.Service, managerRepository) // service
	managerController := controller.NewManagerController(managerService)         // controller
	{
		// ✅ 不需要登录
		manager.POST("/login", managerController.Login)
		// manager.POST("/register", managerController.Create)

		// ✅ 需要登录
		auth := manager.Group("/")
		auth.Use(middleware.StrictAuth(deps.JWT, deps.Logger)) // 依赖注入 JWT 和 Logger
		{
			auth.GET("/users", managerController.GetManagerLists) // 分页用户列表
		}
	}

}
