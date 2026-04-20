package controller

import (
	v1 "go-server/api/v1"
	"go-server/internal/dto/category"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func NewCategoryController(handler *Handler, s service.CategoryService) *categoryController {
	return &categoryController{
		Handler: handler,
		Service: s,
	}
}

type categoryController struct {
	*Handler
	Service service.CategoryService // 依赖注入
}

// ================= 创建 =================

// @Summary 商品分类创建
// @Tags 商品分类
// @Accept json
// @Produce json
// @Param data body category.CreateRequest true "注册参数"
// @Success 201 {object} category.PrivateDTO
// @Router /api/private/v1/categories [post]
func (u *categoryController) Create(c *gin.Context) {
	var req category.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误")
		return
	}

	user, err := u.Service.Create(c, req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Created(c, category.ToPrivateDTO(user))
}

// ================= 删除id信息 =================

// @Summary 商品分类删除
// @Tags 商品分类
// @Produce json
// @Param id path int true "ID"
// @Success 204 {string} string "No Content"
// @Router /api/private/v1/categories/{id} [delete]
func (u *categoryController) Delete(c *gin.Context) {
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

// @Summary 商品分类更新
// @Tags 商品分类
// @Accept json
// @Produce json
// @Param data body category.UpdateRequest true "更新参数"
// @Success 200 {object} category.PrivateDTO
// @Router /api/private/v1/categories [put]
func (u *categoryController) Update(c *gin.Context) {
	id, ok := GetUintID(c, "id")
	if !ok {
		return
	}

	var req category.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		v1.BadRequest(c, "请求参数错误"+err.Error())
		return
	}

	user, err := u.Service.Update(c, uint(id), req)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	v1.Success(c, category.ToPrivateDTO(user))
}

// ================= 获取id详情 =================

// @Summary 获取详情
// @Tags 商品分类
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} category.ToPublicDTO
// @Router /api/private/v1/categories/{id} [get]

func (u *categoryController) GetDetail(c *gin.Context) {
	id, ok := GetUintID(c, "id")
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
		v1.Success(c, category.ToPrivateDTO(user))
	} else {
		v1.Success(c, category.ToPublicDTO(user))
	}
}

// ================= 列表 =================

// @Summary 商品分类列表
// @Tags 商品分类
// @Produce json
// @Param data query category.RequestQuery false "查询参数"
// @Success 200 {object} []category.PublicDTO
// @Router  /api/private/v1/categories [get]
func (u *categoryController) GetList(c *gin.Context) {
	var q category.RequestQuery

	if err := c.ShouldBindQuery(&q); err != nil {
		v1.BadRequest(c, "参数错误")
		return
	}

	if q.Page != nil || q.PageSize != nil {
		u.GetPageList(c)
		return
	}

	users, err := u.Service.GetList(c, q)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := category.ListToPublic(users)

	v1.Success(c, list)
}

// ================= 分页列表 =================

// @Summary 商品分类列表-分页
// @Tags 商品分类
// @Produce json
// @Param data query category.RequestPageQuery false "查询参数"
// @Success 200 {object} v1.PageResponse
// @Router /api/private/v1/categories/lists [get]

func (u *categoryController) GetPageList(c *gin.Context) {
	var q category.RequestPageQuery

	if err := c.ShouldBindQuery(&q); err != nil {
		v1.BadRequest(c, "参数错误")
		return
	}

	users, total, err := u.Service.GetPageList(c, q)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := category.ListToPublic(users)

	v1.List(c, list, int(total), q.Page, q.PageSize)
}
