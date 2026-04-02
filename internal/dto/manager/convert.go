package managerdto

import "go-server/internal/model"

// ================= DTO 转换 =================

// 👉 他人可见
func ToManagerPublicDTO(u *model.Manager) ManagerPublicDTO {
	return ManagerPublicDTO{
		ID:   u.MgID,
		User: u.MgName,
		// Avatar:   u.Avatar,
	}
}

// 👉 自己可见
func ToManagerPrivateDTO(u *model.Manager) ManagerPrivateDTO {
	name := u.MgName
	if u.RoleID == 0 {
		name = "超级管理员"
	}
	return ManagerPrivateDTO{
		MgID:      u.MgID,
		MgName:    u.MgName,
		MgEmail:   u.MgEmail,
		MgMobile:  u.MgMobile,
		MgState:   u.MgState == 1,
		Role_name: name,
	}
}

// // 👉 自己可见
// func ToUserPrivateDTO(u *model.Manager) UserPrivateDTO {
// 	return UserPrivateDTO{
// 		ID:       u.ID,
// 		Username: u.Username,
// 		// Avatar:   u.Avatar,
// 	}
// }

func ManagerListToPublic(users []model.Manager) []ManagerPrivateDTO {
	list := make([]ManagerPrivateDTO, 0, len(users))
	for _, u := range users {
		list = append(list, ToManagerPrivateDTO(&u))
	}
	return list
}
