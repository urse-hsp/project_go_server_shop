// model 集中注册
package model

func GetModels() []interface{} {
	return []interface{}{
		&User{},
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
		&Manager{},
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
