package roledto

import "go-server/internal/model"

func ToRolePublicDTO(u *model.Role) RolePublicDTO {
	return RolePublicDTO{
		ID:       u.RoleID,
		RoleName: u.RoleName,
		RoleDesc: u.RoleDesc,
	}
}

func RoleListToPublic(users []model.Role) []RolePublicDTO {
	list := make([]RolePublicDTO, 0, len(users))
	for _, u := range users {
		list = append(list, ToRolePublicDTO(&u))
	}
	return list
}
