package attributedto

import "go-server/internal/model"

// ================= DTO 转换 =================

// 👉 他人可见
func ToPublicDTO(u *model.Attribute) PublicDTO {
	return PublicDTO{
		AttrID:    u.AttrID,
		AttrName:  u.AttrName,
		CatID:     u.CatID,
		AttrSel:   u.AttrSel,
		AttrWrite: u.AttrWrite,
		AttrVals:  u.AttrVals,
	}
}

// 👉 自己可见
func ToPrivateDTO(u *model.Attribute) PrivateDTO {
	return PrivateDTO{}
}

func ListToPublic(users []model.Attribute) []PublicDTO {
	list := make([]PublicDTO, 0, len(users))
	for _, u := range users {
		list = append(list, ToPublicDTO(&u))
	}
	return list
}
