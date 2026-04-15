package controller

import (
	v1 "go-server/api/v1"
	roledto "go-server/internal/dto/role"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func NewColeController(handler *Handler, s service.RoleService) *roleController {
	return &roleController{
		Handler:     handler,
		roleService: s,
	}
}

type roleController struct {
	*Handler
	roleService service.RoleService // 依赖注入
}

// ================= 创建 =================

// @Summary 创建
// @Tags 角色
// @Accept json
// @Produce json
// @Param data body roledto.LoginRequest true "创建参数"
// @Success 201 {object} roledto.RolePublicDTO
// @Router /api/private/v1/roles [post]
func (u *roleController) Create(c *gin.Context) {
	var req roledto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误")
		return
	}

	user, err := u.roleService.Create(c, req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Created(c, roledto.ToRolePublicDTO(user))
}

// ================= 删除 =================

// @Summary 删除用户
// @Tags 角色
// @Produce json
// @Param id path int true "用户ID"
// @Success 204 {string} string "No Content"
// @Router /api/private/v1/roles/{id} [delete]
func (u *roleController) Delete(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	if err := u.roleService.Delete(c, uint(id)); err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.NoContent(c)
}

// ================= 更新 =================

// @Summary 更新当前用户
// @Tags 角色
// @Accept json
// @Produce json
// @Param data body roledto.LoginRequest true "更新参数"
// @Success 200 {object} userdto.UserPrivateDTO
// @Router /api/private/v1/roles/info [put]
func (u *roleController) Update(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	var req roledto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误"+err.Error())
		return
	}

	user, err := u.roleService.Update(c, uint(id), req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Success(c, roledto.ToRolePublicDTO(user))
}

// ================= 根据ID查询 =================

// @Summary 获取用户详情
// @Tags 角色
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} userdto.UserPublicDTO
// @Router /api/private/v1/roles/{id} [get]
func (u *roleController) GetDetail(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	user, err := u.roleService.GetDetail(c, uint(id))
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Success(c, roledto.ToRolePublicDTO(user))
}

// ================= 列表 =================
// @Summary 角色列表
// @Tags 角色
// @Produce json
// @Success 200 {object} []roledto.RolePublicDTO
// @Router /api/private/v1/roles [get]
func (u *roleController) GetList(c *gin.Context) {
	users, err := u.roleService.GetList(c)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := roledto.RoleListToPublic(users)

	v1.Success(c, list)
}

// ================= 分页列表 =================

// @Summary 用户列表 分页
// @Tags 角色
// @Produce json
// @Success 200 {object} v1.PageResponse
// @Router /api/private/v1/roles/lists [get]
func (u *roleController) GetLists(c *gin.Context) {
	page, pageSize := v1.GetPage(c)

	users, total, err := u.roleService.GetLists(c, page, pageSize)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := roledto.RoleListToPublic(users)

	v1.List(c, list, int(total), page, pageSize)
}
