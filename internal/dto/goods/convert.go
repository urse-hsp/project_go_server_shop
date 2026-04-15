package goodsdto

import (
	"fmt"
	"go-server/internal/model"
)

// ================= DTO 转换 =================

// 他人可见
func ToPublicDTO(u *model.Goods) PublicDTO {
	return PublicDTO{
		GoodsID:     u.GoodsID,
		GoodsName:   u.GoodsName,
		GoodsNumber: u.GoodsNumber,
		GoodsPrice:  u.GoodsPrice,
		GoodsState:  u.GoodsState,
		GoodsWeight: u.GoodsWeight,
		IsPromote:   u.IsPromote,
		AddTime:     uint(u.AddTime),
	}
}

func ToDetailDTO(u *model.Goods) DetailPublicDTO {
	// 转换 attrs
	attrs := make([]AttrDTO, 0, len(u.Attrs))
	for _, a := range u.Attrs {
		attrs = append(attrs, AttrDTO{
			ID:        a.ID,
			GoodsId:   a.GoodsId,
			AttrId:    a.AttrId,
			AttrValue: a.AttrValue,
			AddPrice:  a.AddPrice,
		})
	}

	// 转换 pics
	pics := make([]PicDTO, 0, len(u.Pics))
	for _, p := range u.Pics {
		pics = append(pics, PicDTO{
			PicsID:  p.PicsID,
			GoodsID: p.GoodsId,
			PicsBig: p.PicsBig,
			PicsMid: p.PicsMid,
			PicsSma: p.PicsSma,
		})
	}
	return DetailPublicDTO{
		Attrs:     attrs,
		Pics:      pics,
		PublicDTO: ToPublicDTO(u),

		GoodsIntroduce: u.GoodsIntroduce,
		CatId:          u.CatID,
		CatOneID:       u.CatOneID,
		CatTwoID:       u.CatTwoID,
		CatThreeID:     u.CatThreeID,
		GoodsCat:       fmt.Sprintf("%d,%d,%d", u.CatOneID, u.CatTwoID, u.CatThreeID),
		GoodsBigLogo:   u.GoodsBigLogo,
		GoodsSmallLogo: u.GoodsSmallLogo,
		UpdTime:        uint(u.UpdTime),
	}
}

// 自己可见
func ToPrivateDTO(u *model.Goods) PrivateDTO {
	return PrivateDTO{}
}

func ListToPublic(users []model.Goods) []PublicDTO {
	list := make([]PublicDTO, 0, len(users))
	for _, u := range users {
		list = append(list, ToPublicDTO(&u))
	}
	return list
}
