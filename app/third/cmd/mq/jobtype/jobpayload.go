package jobtype

// SyncShopBasicPayload 同步店铺数据
type SyncShopBasicPayload struct {
	JkShopId int64
	JkUserId int64
}

// WebhookOrderPayload 同步webhook订单
type WebhookOrderPayload struct {
	JkShopId int64
	JkUserId int64
	OrderId  string
}
