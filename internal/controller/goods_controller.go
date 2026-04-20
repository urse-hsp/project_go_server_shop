package controller

import (
	v1 "go-server/api/v1"
	goodsdto "go-server/internal/dto/goods"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func NewGoodsController(handler *Handler, s service.GoodsService) *goodsController {
	return &goodsController{
		Handler: handler,
		Service: s,
	}
}

type goodsController struct {
	*Handler
	Service service.GoodsService // 依赖注入
}

// ================= 创建 =================

// @Summary 注册
// @Tags 商品
// @Accept json
// @Produce json
// @Param data body goodsdto.CreateRequest true "注册参数"
// @Success 201 {object} goodsdto.DetailPublicDTO
// @Router /api/private/v1/goods [post]
func (u *goodsController) Create(c *gin.Context) {
	var req goodsdto.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误"+err.Error())
		return
	}

	data, err := u.Service.Create(c, req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Created(c, goodsdto.ToDetailDTO(data))
}

// ================= 删除id信息 =================

// @Summary 删除
// @Tags 商品
// @Produce json
// @Param id path int true "ID"
// @Success 204 {string} string "No Content"
// @Router /api/private/v1/goods/{id} [delete]
func (u *goodsController) Delete(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	if err := u.Service.Delete(c, uint(id)); err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.NoContent(c)
}

// ================= 更新当前id信息 =================

// @Summary 更新
// @Tags 商品
// @Accept json
// @Produce json
// @Param data body goodsdto.UpdateRequest true "更新参数"
// @Success 200 {object} goodsdto.DetailPublicDTO
// @Router /api/private/v1/goods/info [put]
func (u *goodsController) Update(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	var req goodsdto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误"+err.Error())
		return
	}

	data, err := u.Service.Update(c, req, id)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Success(c, goodsdto.ToDetailDTO(data))
}

// ================= 获取id详情 =================

// @Summary 获取详情
// @Tags 商品
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} goodsdto.DetailPublicDTO
// @Router /api/private/v1/goods/{id} [get]
func (u *goodsController) GetDetail(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	user, err := u.Service.GetDetail(c, uint(id))
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Success(c, goodsdto.ToDetailDTO(user))
}

// ================= 分页列表 =================

// @Summary 列表 分页
// @Tags 商品
// @Produce json
// @Param data query goodsdto.RequestPageQuery false "查询参数"
// @Success 200 {object} goodsdto.PageResponse
// @Router /api/private/v1/goods [get]
func (u *goodsController) GetLists(c *gin.Context) {
	var q goodsdto.RequestPageQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		v1.BadRequest(c, "参数错误")
		return
	}
	q.Normalize()

	users, total, err := u.Service.GetLists(c, q)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := goodsdto.ListToPublic(users)

	v1.List(c, list, int(total), q.Page, q.PageSize)
}
