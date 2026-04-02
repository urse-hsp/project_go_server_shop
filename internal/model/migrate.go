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
		&permission{},
		&permissionApi{},
		&orderGoods{},
		&order{},
		&Manager{},
		&Goods{},
		&GoodsPics{},
		&GoodsCats{},
		&goodsAttr{},
		&Category{},
		&Attribute{},
		&Consignee{},
		&Express{},
	}
}
