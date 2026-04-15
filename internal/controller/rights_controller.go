package controller

import (
	v1 "go-server/api/v1"
	rightsdto "go-server/internal/dto/rights"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRightsController(handler *Handler, s service.RightsService) *rightsController {
	return &rightsController{
		Handler:       handler,
		rightsService: s,
	}
}

type rightsController struct {
	*Handler
	rightsService service.RightsService // 依赖注入
}

// ================= 列表 =================
// @Summary 角色列表
// @Tags 角色
// @Produce json
// @Success 200 {object} []roledto.RolePublicDTO
// @Router /api/private/v1/roles [get]
func (u *rightsController) GetList(c *gin.Context) {
	users, err := u.rightsService.GetList(c)
	if err != nil {
		v1.BadRequest(c, err.Error())
		return
	}

	list := rightsdto.ListToPublic(users)

	v1.Success(c, list)
}
