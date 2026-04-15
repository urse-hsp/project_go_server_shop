package category

import "go-server/internal/model"

// ================= DTO 转换 =================

// 他人可见
func ToPublicDTO(u *model.Category) PublicDTO {

	return PublicDTO{
		CatsID:     u.CatID,
		CatName:    u.CatName,
		CatPID:     u.CatPID,
		CatLevel:   u.CatLevel,
		CatDeleted: u.CatDeleted == 1,
		Children:   ListToPublic(u.Children),
	}
}

// 自己可见
func ToPrivateDTO(u *model.Category) PrivateDTO {
	return PrivateDTO{}
}

func ListToPublic(users []model.Category) []PublicDTO {
	list := make([]PublicDTO, 0, len(users))

	for i := range users {
		list = append(list, ToPublicDTO(&users[i]))
	}

	return list
}
