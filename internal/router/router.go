package router

import (
	"go-server/internal/bootstrap"
	"go-server/internal/controller"
	"go-server/internal/middleware"
	"go-server/internal/service"
	"go-server/pkg/jwt"
	"go-server/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RouterDeps struct {
	Logger     *log.Logger
	Config     *viper.Viper
	Repository *bootstrap.Repository // dao层工具包
	Service    *service.Service      // 业务层工具包
	Handler    *controller.Handler   // 控制层工具包
	JWT        *jwt.JWT
}

func SetupRouter(deps RouterDeps) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	// 静态资源
	r.Static("/web", "./web/dist")
	// SPA 兜底。解决前端路由刷新/直达访问 404
	r.NoRoute(func(c *gin.Context) {
		c.File("./web/dist/index.html")
	})

	// 全局中间件
	r.Use(
		middleware.CORSMiddleware(),
		middleware.RequestLogMiddleware(deps.Logger),  // 依赖注入日志组件
		middleware.ResponseLogMiddleware(deps.Logger), // 依赖注入日志组件
	)

	api := r.Group("/api/private/v1")

	// ================= 用户模块 =================
	// InitUserRouter(deps, api)

	// ================= 管理员模块 =================
	InitManagerRouter(deps, api)

	return r
}
