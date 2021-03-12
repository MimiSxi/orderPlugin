package model

import (
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

// 订单信息
type OrderInfo struct {
	ID           uint                      `gorm:"primary_key" gqlschema:"query!;querys;update!" description:"订单id"`
	UserId       uint                      `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"创建用户id" funservice:"employee"`
	ChildrenId   uint                      `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys;update;" description:"引用id"`
	ChildrenType OrderInfoChildrenTypeEnum `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys;update;" description:"引用类型"`
	Remark       string                    `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;update" description:"备注"`
	CreatedAt    time.Time                 `description:"创建时间" gqlschema:"querys"`
	UpdatedAt    time.Time                 `description:"更新时间" gqlschema:"querys"`
	DeletedAt    *time.Time
	v2           int `gorm:"-" exclude:"true"`
}

type OrderInfos struct {
	TotalCount int
	Edges      []OrderInfo
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
	o.ChildrenType = p["childrenType"].(OrderInfoChildrenTypeEnum)
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
	if p["childrenId"] != nil {
		v.ChildrenId = uint(p["childrenId"].(int))
	}
	if p["childrenType"] != nil {
		v.ChildrenType = p["childrenType"].(OrderInfoChildrenTypeEnum)
	}
	if p["remark"] != nil {
		v.Remark = p["remark"].(string)
	}
	err := db.Save(&v).Error
	return v, err
}
