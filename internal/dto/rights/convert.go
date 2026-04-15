package rightsdto

import "go-server/internal/model"

func ToPublicDTO(u *model.Permission) PublicDTO {
	return PublicDTO{
		ID:      u.ID,
		PsName:  u.PsName,
		PsPid:   u.PsPid,
		PsC:     u.PsC,
		PsLevel: u.PsLevel,
	}
}

func ListToPublic(users []model.Permission) []PublicDTO {
	list := make([]PublicDTO, 0, len(users))
	for _, u := range users {
		list = append(list, ToPublicDTO(&u))
	}
	return list
}
