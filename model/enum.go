package model

import "github.com/Fiber-Man/funplugin"

// 订单类型
type OrderInfoChildrenTypeEnum uint

const (
	PHOTO_ALBUM OrderInfoChildrenTypeEnum = 1 // 相册系统
)

func (s OrderInfoChildrenTypeEnum) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"PHOTO_ALBUM": funplugin.EnumValue{
			Value:       PHOTO_ALBUM,
			Description: "相册系统",
		},
	}
}

// 订单状态枚举类型
type OrderStatusEnumType uint

const (
	TO_BE_PAID    OrderStatusEnumType = 1 //待支付
	TO_BE_DELIVER OrderStatusEnumType = 2 //待发货
	HAS_DELIVER   OrderStatusEnumType = 3 //已发货
	HAS_RECEIVED  OrderStatusEnumType = 4 //已收货
	CANCELED      OrderStatusEnumType = 5 //订单取消
)

func (s OrderStatusEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"TO_BE_PAID": funplugin.EnumValue{
			Value:       TO_BE_PAID,
			Description: "待支付",
		},
		"TO_BE_DELIVER": funplugin.EnumValue{
			Value:       TO_BE_DELIVER,
			Description: "待发货",
		},
		"HAS_DELIVER": funplugin.EnumValue{
			Value:       HAS_DELIVER,
			Description: "已发货",
		},
		"HAS_RECEIVED": funplugin.EnumValue{
			Value:       HAS_RECEIVED,
			Description: "已收货",
		},
		"CANCELED": funplugin.EnumValue{
			Value:       CANCELED,
			Description: "订单取消",
		},
	}
}

// 订单支付方式枚举类型
type OrderPayWayEnumType uint

const (
	WECHAT OrderPayWayEnumType = 1 //微信
	ALIPAY OrderPayWayEnumType = 2 //支付宝
)

func (s OrderPayWayEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"WECHAT": funplugin.EnumValue{
			Value:       WECHAT,
			Description: "微信",
		},
		"ALIPAY": funplugin.EnumValue{
			Value:       ALIPAY,
			Description: "支付宝",
		},
	}
}

type TradeStatus uint

const (
	TradeStatusWaitBuyerPay TradeStatus = 1 //（交易创建，等待买家付款）
	TradeStatusClosed       TradeStatus = 2 //（未付款交易超时关闭，或支付完成后全额退款）
	TradeStatusSuccess      TradeStatus = 3 //（交易支付成功）
	TradeStatusFinished     TradeStatus = 4 //（交易结束，不可退款）
)

func (s TradeStatus) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"TradeStatusWaitBuyerPay": funplugin.EnumValue{
			Value:       TradeStatusWaitBuyerPay,
			Description: "交易创建，等待买家付款",
		},
		"TradeStatusClosed": funplugin.EnumValue{
			Value:       TradeStatusClosed,
			Description: "未付款交易超时关闭，或支付完成后全额退款",
		},
		"TradeStatusSuccess": funplugin.EnumValue{
			Value:       TradeStatusSuccess,
			Description: "交易支付成功",
		},
		"TradeStatusFinished": funplugin.EnumValue{
			Value:       TradeStatusFinished,
			Description: "交易结束，不可退款",
		},
	}
}
