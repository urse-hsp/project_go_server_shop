package controller

import (
	v1 "go-server/api/v1"
	orderdto "go-server/internal/dto/order"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func NewOrderController(handler *Handler, s service.OrderService) *orderController {
	return &orderController{
		Handler: handler,
		Service: s,
	}
}

type orderController struct {
	*Handler
	Service service.OrderService // 依赖注入
}

// ================= 创建 =================

// @Summary 订单创建
// @Tags 订单
// @Accept json
// @Produce json
// @Param data body orderdto.CreateRequest true "注册参数"
// @Success 201 {object} orderdto.PrivateDTO
// @Router /api/private/v1/orders [post]

func (u *orderController) Create(c *gin.Context) {
	var req orderdto.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误")
		return
	}

	user, err := u.Service.Create(c, req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Created(c, orderdto.ToPrivateDTO(user))
}

// ================= 删除id信息 =================

// @Summary 订单删除
// @Tags 订单
// @Produce json
// @Param id path int true "ID"
// @Success 204 {string} string "No Content"
// @Router /api/private/v1/orders/{id} [delete]

func (u *orderController) Delete(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	currentUserID := GetUserIdFromCtx(c)

	// 只允许删除自己（可扩展管理员）
	if currentUserID != uint(id) {
		v1.Forbidden(c, "无权限删除他人")
		return
	}

	if err := u.Service.Delete(c, uint(id)); err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.NoContent(c)
}

// ================= 更新当前id信息 =================

// @Summary 订单更新
// @Tags 订单
// @Accept json
// @Produce json
// @Param data body orderdto.UpdateRequest true "更新参数"
// @Success 200 {object} orderdto.PrivateDTO
// @Router /api/private/v1/orders [put]

func (u *orderController) Update(c *gin.Context) {
	id, idErr := ParseUintParam(c, "id")
	if idErr != nil {
		v1.BadRequest(c, idErr.Error())
		return
	}

	var req orderdto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误"+err.Error())
		return
	}

	user, err := u.Service.Update(c, id, req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Success(c, orderdto.ToPrivateDTO(user))
}

// ================= 获取id详情 =================

// @Summary 获取详情
// @Tags 订单
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} orderdto.PublicDTO
// @Router /api/private/v1/orders/{id} [get]
func (u *orderController) GetDetail(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	data, err := u.Service.GetDetail(c, uint(id))
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Success(c, orderdto.ToPublicDTO(data))
}

// ================= 列表 =================

// @Summary 订单列表
// @Tags 订单
// @Produce json
// @Param data query orderdto.RequestQuery false "查询参数"
// @Success 200 {object} []orderdto.PublicDTO
// @Router /api/private/v1/orders [get]

func (u *orderController) GetList(c *gin.Context) {
	var q orderdto.RequestQuery

	if err := c.ShouldBindQuery(&q); err != nil {
		v1.BadRequest(c, "参数错误")
		return
	}

	users, err := u.Service.GetList(c, q)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := orderdto.ListToPublic(users)

	v1.Success(c, list)
}

// ================= 分页列表 =================

// @Summary 订单列表-分页
// @Tags 订单
// @Produce json
// @Param data query orderdto.RequestPageQuery false "查询参数"
// @Success 200 {object} orderdto.PageResponse
// @Router /api/private/v1/orders/lists [get]
func (u *orderController) GetPageList(c *gin.Context) {
	var q orderdto.RequestPageQuery

	if err := c.ShouldBindQuery(&q); err != nil {
		v1.BadRequest(c, "参数错误")
		return
	}

	users, total, err := u.Service.GetPageList(c, q)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := orderdto.ListToPublic(users)

	v1.List(c, list, int(total), q.Page, q.PageSize)
}
