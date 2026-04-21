// model 集中注册
package model

func GetModels() []interface{} {
	return []interface{}{
		&Manager{},
		&UserCart{},
		&Type{},
		&Role{},
		&report_1{},
		&report_2{},
		&report_3{},
		&Permission{},
		&permissionApi{},
		&orderGoods{},
		&Order{},
		&Goods{},
		&GoodsPics{},
		&GoodsCats{},
		&GoodsAttr{},
		&Category{},
		&Attribute{},
		&Consignee{},
		&Express{},
	}
}
