package controller

import (
	v1 "go-server/api/v1"
	managerdto "go-server/internal/dto/manager"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

type managerController struct {
	managerService service.ManagerService // 依赖注入
}

func NewManagerController(s service.ManagerService) *managerController {
	return &managerController{
		managerService: s,
	}
}

// ================= 登录 =================

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

// ================= 注册 =================

// @Summary 管理员创建
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data body managerdto.CreateRequest true "创建参数"
// @Success 201 {string} string "No Content"
// @Router /api/private/v1/users [post]
func (u *managerController) Create(c *gin.Context) {
	var req managerdto.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误")
		return
	}

	_, err := u.managerService.Create(c, req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Created(c, true)
}

// ================= 删除用户 =================

// @Summary 删除用户
// @Tags 管理员
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {string} string "No Content"
// @Router /api/private/v1/users/{id} [delete]
func (u *managerController) Delete(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	// currentUserID := GetUserIdFromCtx(c)
	// // 只允许删除自己（可扩展管理员）
	// if currentUserID != uint(id) {
	// 	v1.Forbidden(c, "无权限删除他人")
	// 	return
	// }

	if err := u.managerService.Delete(c, uint(id)); err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Success(c, true)
}

// ================= 更新用户 =================

// @Summary 更新用户
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data body managerdto.UpdateRequest true "更新参数"
// @Success 200 {object} string "No Content"
// @Router /api/private/v1/users/{id} [put]
func (u *managerController) Update(c *gin.Context) {
	id, idErr := ParseUintParam(c, "id")
	if idErr != nil {
		v1.BadRequest(c, idErr.Error())
		return
	}

	var req managerdto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误"+err.Error())
		return
	}
	if req.Email == nil && req.Mobile == nil && req.State == nil {
		v1.BadRequest(c, "至少需要传一个更新字段")
		return
	}

	_, err := u.managerService.Update(c, id, req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Success(c, true)
}

// ================= 分页列表  =================

// @Summary 管理员列表
// @Tags 管理员
// @Produce json
// @Param data query managerdto.ManagerQuery false "查询参数"
// @Success 200 {object} managerdto.PageResponse
// @Router /api/private/v1/users [get]
func (c *managerController) GetLists(ctx *gin.Context) {
	var q managerdto.ManagerQuery
	if err := ctx.ShouldBindQuery(&q); err != nil {
		v1.BadRequest(ctx, "参数错误"+err.Error())
		return
	}
	q.Normalize()

	users, total, err := c.managerService.GetLists(ctx, q)
	if err != nil {
		v1.BadRequest(ctx, err.Error())
		return
	}

	list := managerdto.ManagerListToPublic(users)

	v1.List(ctx, list, int(total), q.Page, q.PageSize)
}
