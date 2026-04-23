package router

import (
	"go-server/internal/controller"
	"go-server/internal/dao"
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
	Repository *dao.Repository     // dao层工具包
	Service    *service.Service    // 业务层工具包
	Handler    *controller.Handler // 控制层工具包
	JWT        *jwt.JWT
}

func SetupRouter(deps RouterDeps) *gin.Engine {
	r := gin.Default()

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

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.Static("/storage/uploads", "./storage/uploads") // 模拟上传图片。托管到服务中

	v1 := r.Group("/api/private/v1")

	// ================= 管理员模块 =================
	InitManagerRouter(deps, v1)

	// ================= 角色模块 =================
	InitRoleRouter(deps, v1)

	// ================= 权限模块 =================
	InitRightsRouter(deps, v1)

	// ================= 商品模块 =================
	InitGoodsRouter(deps, v1)

	// ================= 类别模块 =================
	categoryR := InitCategoryRouter(deps, v1)

	// ================= 分类属性模块 =================
	InitGoodsAttrRouter(deps, categoryR)

	// ================= 订单模块 =================
	InitOrderRouter(deps, v1)

	// ================= 上传模块 =================
	InitUploadRouter(deps, v1)

	return r
}
