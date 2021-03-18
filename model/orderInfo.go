package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

// 订单信息
type OrderInfo struct {
	ID           uint                `gorm:"primary_key" gqlschema:"query!;querys;update!" description:"订单id"`
	UserId       uint                `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"创建用户id" funservice:"employee"`
	ChildrenId   uint                `gorm:"unique_refer;" gqlschema:"create!;querys" description:"引用id" funservice:"object_typekey:ChildrenType"`
	ChildrenType string              `gorm:"unique_refer;" gqlschema:"create!;querys" description:"引用类型"`
	Status       OrderStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create!;querys;update;" description:"订单状态"`
	//GoodsProp    PropJson                  `gorm:"Type:text;" gqlschema:"create!;update;" description:"商品属性"`
	Address      string              `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create;querys;update" description:"收货地址"`
	Remark       string              `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create;querys;update" description:"备注"`
	GoodsPrice   uint                `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys;update" description:"商品价格"`
	FreightPrice uint                `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys;update" description:"运费"`
	PaymentId    string              `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"支付id"`
	PayWay       OrderPayWayEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create;update;querys" description:"支付方式枚举类型"`
	PayTime      time.Time           `gorm:"DEFAULT:'1970-1-1 00:00:00';" description:"支付时间" gqlschema:"querys"`
	DeliveryId   string              `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"update;querys" description:"快递id"`
	CreatedAt    time.Time           `description:"创建时间" gqlschema:"querys"`
	UpdatedAt    time.Time           `description:"更新时间" gqlschema:"querys"`
	DeletedAt    *time.Time
	v2           int `gorm:"-" exclude:"true"`
}

type OrderInfos struct {
	TotalCount int
	Edges      []OrderInfo
}

type PropJson map[string]interface{}

func (c PropJson) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c PropJson) Scan(input interface{}) error {
	v, ok := input.([]byte)
	if !ok {
		v = []byte(input.(string))
	}
	err := json.Unmarshal(v, c)
	return err
}

func (o *OrderInfo) QueryByID(id uint) (err error) {
	return db.Where("id = ?", id).First(&o).Error
}

func (o OrderInfo) Query(params graphql.ResolveParams) (OrderInfo, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o OrderInfo) Querys(params graphql.ResolveParams) (OrderInfos, error) {
	var result OrderInfos

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o OrderInfo) Create(params graphql.ResolveParams) (OrderInfo, error) {
	p := params.Args
	o.UserId = uint(p["userId"].(int))
	o.ChildrenId = uint(p["childrenId"].(int))
	o.ChildrenType = p["childrenType"].(string)
	o.Status = p["status"].(OrderStatusEnumType)
	//o.GoodsProp = nil // todo 商品json
	if p["address"] != nil {
		o.Address = p["address"].(string)
	}
	if p["freightPrice"] != nil {
		o.FreightPrice = uint(p["freightPrice"].(int))
	}
	if p["goodsPrice"] != nil {
		o.GoodsPrice = uint(p["goodsPrice"].(int))
	}
	if p["paymentId"] != nil {
		o.PaymentId = p["paymentId"].(string)
	}
	if p["payWay"] != nil {
		o.PayWay = p["payWay"].(OrderPayWayEnumType)
	}
	if p["remark"] != nil {
		o.Remark = p["remark"].(string)
	}
	err := db.Create(&o).Error
	return o, err
}

func (o OrderInfo) Update(params graphql.ResolveParams) (OrderInfo, error) {
	v, ok := params.Source.(OrderInfo)
	if !ok {
		return o, errors.New("update param")
	}
	p := params.Args
	if p["status"] != nil {
		v.Status = p["status"].(OrderStatusEnumType)
	}
	//if p["goodsProp"] != nil {
	//	v.GoodsProp = nil // todo 商品json
	//}
	if p["address"] != nil {
		v.Address = p["address"].(string)
	}
	if p["remark"] != nil {
		v.Remark = p["remark"].(string)
	}
	if p["goodsPrice"] != nil {
		v.GoodsPrice = uint(p["goodsPrice"].(int))
	}
	if p["freightPrice"] != nil {
		v.FreightPrice = uint(p["freightPrice"].(int))
	}
	if p["paymentId"] != nil {
		v.PaymentId = p["paymentId"].(string)
	}
	if p["payWay"] != nil {
		v.PayWay = p["payWay"].(OrderPayWayEnumType)
	}
	if p["deliveryId"] != nil {
		v.DeliveryId = p["deliveryId"].(string)
	}
	err := db.Save(&v).Error
	return v, err
}
