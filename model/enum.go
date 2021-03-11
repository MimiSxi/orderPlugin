package model

import "github.com/Fiber-Man/funplugin"

// 订单类型
type OrderInfoChildrenTypeEnum uint

const (
	PHOTO_ALBUM  OrderInfoChildrenTypeEnum = 1 // 相册系统
)

func (s OrderInfoChildrenTypeEnum) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"PHOTO_ALBUM": funplugin.EnumValue{
			Value:       PHOTO_ALBUM,
			Description: "相册系统",
		},
	}
}