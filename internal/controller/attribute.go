package controller

import (
	v1 "go-server/api/v1"
	attributedto "go-server/internal/dto/attribute"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func NewAttributeController(handler *Handler, s service.AttributeService) *attributeController {
	return &attributeController{
		Handler: handler,
		Service: s,
	}
}

type attributeController struct {
	*Handler
	Service service.AttributeService // 依赖注入
}

// ================= 创建 =================

// @Summary 分类参数创建
// @Tags DEMO
// @Accept json
// @Produce json
// @Param data body attributedto.CreateRequest true "注册参数"
// @Success 201 {object} attributedto.UserPrivateDTO
// @Router /demo [post]

func (u *attributeController) Create(c *gin.Context) {
	id, ok := GetId(c)
	if !ok {
		return
	}

	var req attributedto.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误")
		return
	}

	user, err := u.Service.Create(c, id, req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Created(c, attributedto.ToPrivateDTO(user))
}

// ================= 删除id信息 =================

// @Summary 分类参数删除
// @Tags DEMO
// @Produce json
// @Param id path int true "ID"
// @Success 204 {string} string "No Content"
// @Router /categories/:id/attributes/{id} [delete]

func (u *attributeController) Delete(c *gin.Context) {
	_, ok := GetId(c)
	if !ok {
		return
	}

	ids, oks := GetUintID(c, "attrId")
	if !oks {
		return
	}

	if err := u.Service.Delete(c, ids); err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.NoContent(c)
}

// ================= 更新当前id信息 =================

// @Summary 分类参数更新
// @Tags DEMO
// @Accept json
// @Produce json
// @Param data body attributedto.UpdateRequest true "更新参数"
// @Success 200 {object} attributedto.UserPrivateDTO
// @Router /demo [put]

func (u *attributeController) Update(c *gin.Context) {
	ids, oks := GetUintID(c, "attrId")
	if !oks {
		return
	}

	var req attributedto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误"+err.Error())
		return
	}

	_, err := u.Service.Update(c, ids, req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.NoContent(c)
}

// ================= 获取id详情 =================

// @Summary 获取详情
// @Tags DEMO
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} attributedto.UserPublicDTO
// @Router /categories/:id/attributes/{id} [get]

func (u *attributeController) GetDetail(c *gin.Context) {
	id, ok := GetId(c)
	if !ok {
		return
	}

	user, err := u.Service.GetDetail(c, uint(id))
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	currentUserID := GetUserIdFromCtx(c)

	// 权限控制：自己 vs 他人
	if currentUserID == uint(id) {
		v1.Success(c, attributedto.ToPrivateDTO(user))
	} else {
		v1.Success(c, attributedto.ToPublicDTO(user))
	}
}

// ================= 列表 =================

// @Summary 分类参数列表
// @Tags DEMO
// @Produce json
// @Param data query attributedto.RequestQuery false "查询参数"
// @Success 200 {object} []attributedto.UserPublicDTO
// @Router /user [get]

func (u *attributeController) GetList(c *gin.Context) {
	id, ok := GetId(c)
	if !ok {
		return
	}

	var q attributedto.RequestQuery

	if err := c.ShouldBindQuery(&q); err != nil {
		v1.BadRequest(c, "参数错误")
		return
	}

	users, err := u.Service.GetList(c, id, q)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := attributedto.ListToPublic(users)

	v1.Success(c, list)
}

// ================= 分页列表 =================

// @Summary 分类参数列表-分页
// @Tags DEMO
// @Produce json
// @Param data query attributedto.RequestPageQuery false "查询参数"
// @Success 200 {object} v1.PageResponse
// @Router /categories/:id/attributes/lists [get]

func (u *attributeController) GetPageList(c *gin.Context) {
	id, ok := GetId(c)
	if !ok {
		return
	}

	var q attributedto.RequestPageQuery

	if err := c.ShouldBindQuery(&q); err != nil {
		v1.BadRequest(c, "参数错误")
		return
	}

	users, total, err := u.Service.GetPageList(c, id, q)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := attributedto.ListToPublic(users)

	v1.List(c, list, int(total), q.Page, q.PageSize)
}

// 取模块参数id
func GetId(c *gin.Context) (uint, bool) {
	return GetUintID(c, "id")
}
