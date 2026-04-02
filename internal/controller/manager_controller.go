package controller

import (
	v1 "go-server/api/v1"
	managerdto "go-server/internal/dto/manager"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ManagerController interface {
	Login(ctx *gin.Context)
	GetManagerLists(ctx *gin.Context)
	// Create(ctx *gin.Context)
	// DeleteManager(ctx *gin.Context)
	// GetManagerInfo(ctx *gin.Context)
	// UpdateManager(ctx *gin.Context)
}

type managerController struct {
	managerService service.ManagerService // 依赖注入
}

func NewManagerController(s service.ManagerService) ManagerController {
	return &managerController{
		managerService: s,
	}
}

// @Summary 管理员登录
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data body managerdto.LoginRequest true "登录参数"
// @Success 200 {object} managerdto.LoginResponse
// @Router /api/private/v1/login [post]
func (c *managerController) Login(ctx *gin.Context) {
	// 1. 绑定请求参数
	var req managerdto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// ctx.JSON(400, gin.H{"error": "Invalid request"})
		v1.BadRequest(ctx, err.Error())
		return
	}

	// 2. 调用 Service 进行登录逻辑
	user, token, err := c.managerService.Login(ctx, req.Username, req.Password)
	if err != nil {
		// ctx.JSON(401, gin.H{"error": "Authentication failed"})
		v1.Unauthorized(ctx, err.Error())
		return
	}

	v1.Success(ctx, managerdto.LoginResponse{
		Token:             token,
		ManagerPrivateDTO: managerdto.ToManagerPrivateDTO(user),
	})
}

func (c *managerController) GetManagerLists(ctx *gin.Context) {

	page, pageSize := v1.GetPage(ctx)

	users, total, err := c.managerService.GetManagerLists(ctx, page, pageSize)
	if err != nil {
		v1.BadRequest(ctx, err.Error())
		return
	}

	list := managerdto.ManagerListToPublic(users)

	v1.List(ctx, list, int(total), page, pageSize)
}
