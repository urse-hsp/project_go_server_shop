package service

import (
	"context"
	"errors"
	"fmt"
	"go-server/internal/dao"
	goodsdto "go-server/internal/dto/goods"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type GoodsService interface {
	Create(ctx context.Context, req goodsdto.CreateRequest) (*model.Goods, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, req goodsdto.UpdateRequest, id uint) (*model.Goods, error)
	GetLists(ctx context.Context, q goodsdto.RequestPageQuery) ([]model.Goods, int64, error)

	GetDetail(ctx context.Context, id uint) (*model.Goods, error)
}

func NewGoodsService(
	service *Service,
	Repo dao.GoodsRepository,
) GoodsService {
	return &goodsService{
		Repo:    Repo,
		Service: service,
	}
}

type goodsService struct {
	*Service
	Repo dao.GoodsRepository
}

// ================= 注册 =================

func (s *goodsService) Create(ctx context.Context, req goodsdto.CreateRequest) (*model.Goods, error) {
	// 判断是否已存在
	goods, err := s.Repo.GetByKeyWhere(ctx, req.GoodsName)
	if err == nil && goods != nil {
		return nil, fmt.Errorf("名称已存在")
	}

	ids, err := ParseToUintSlice(req.GoodsCat)
	if err != nil {
		return nil, fmt.Errorf("分类参数错误")
	}

	data := &model.Goods{
		GoodsName:      req.GoodsName,
		GoodsPrice:     req.GoodsPrice,
		GoodsWeight:    req.GoodsWeight,
		GoodsIntroduce: req.GoodsIntroduce,
		GoodsNumber:    req.GoodsNumber,
		CatOneID:       ids[0],
		CatTwoID:       ids[1],
		CatThreeID:     ids[2],

		Attrs: buildGoodsAttrs(req.Attrs),
		Pics:  buildGoodsPics(req.Pics),
	}

	return s.Repo.Create(ctx, data)
}

// ================= 删除 =================

func (s *goodsService) Delete(ctx context.Context, id uint) error {
	// 直接调用 DAO（DAO 已经判断是否存在）
	return s.Repo.Delete(ctx, id)
}

// ================= 更新 =================

func (s *goodsService) Update(ctx context.Context, req goodsdto.UpdateRequest, id uint) (*model.Goods, error) {
	data, err := s.Repo.GetDetail(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("商品不存在")
		}
		return nil, err
	}

	data.GoodsName = req.GoodsName
	data.GoodsPrice = req.GoodsPrice
	data.GoodsWeight = req.GoodsWeight
	data.GoodsIntroduce = req.GoodsIntroduce
	data.GoodsNumber = req.GoodsNumber

	data.Attrs = buildGoodsAttrs(req.Attrs)
	data.Pics = buildGoodsPics(req.Pics)

	return s.Repo.Update(ctx, data)
}

// ================= 获取 =================

func (s *goodsService) GetDetail(ctx context.Context, id uint) (*model.Goods, error) {
	return s.Repo.GetDetail(ctx, id)
}

// ================= 分页列表 =================

func (s *goodsService) GetLists(ctx context.Context, q goodsdto.RequestPageQuery) ([]model.Goods, int64, error) {
	return s.Repo.GetLists(ctx, q)
}

func buildGoodsAttrs(attrs []goodsdto.CreateAttr) []model.GoodsAttr {
	res := make([]model.GoodsAttr, 0, len(attrs))
	for _, a := range attrs {
		res = append(res, model.GoodsAttr{
			AttrId:    a.AttrId,
			AttrValue: a.AttrValue,
		})
	}
	return res
}

func buildGoodsPics(pics []goodsdto.CreatePics) []model.GoodsPics {
	res := make([]model.GoodsPics, 0, len(pics))
	for _, p := range pics {
		res = append(res, model.GoodsPics{
			PicsBig: p.Url,
			PicsMid: p.Url,
			PicsSma: p.Url,
		})
	}
	return res
}
