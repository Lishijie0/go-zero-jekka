package thirdconst

// 三方类型
const (
	ThirdTypeShopify         = 1
	ThirdTypeShopee          = 2
	ThirdTypeTikTok          = 3
	ThirdTypeCustomWebsite   = 4
	ThirdTypeTikTokSellerAPI = 5
	ThirdTypeInstagram       = 6
	ThirdTypeFacebook        = 7
	ThirdTypeAmazon          = 8
	ThirdTypeLazada          = 9
)

// ThirdTypeDescMap 三方类型描述映射
var ThirdTypeDescMap = map[int]string{
	ThirdTypeShopify:         "Shopify",
	ThirdTypeShopee:          "Shopee",
	ThirdTypeTikTok:          "TikTok",
	ThirdTypeCustomWebsite:   "Custom Website",
	ThirdTypeTikTokSellerAPI: "TikTok Seller API",
	ThirdTypeInstagram:       "Instagram",
	ThirdTypeAmazon:          "Amazon",
	ThirdTypeLazada:          "Lazada",
}

// ModelChannelTypeMap 传入模型的渠道类型映射
var ModelChannelTypeMap = map[int]string{
	ThirdTypeShopify:         "shopify",
	ThirdTypeShopee:          "shopeelib",
	ThirdTypeTikTok:          "tiktoklib",
	ThirdTypeCustomWebsite:   "custom_website",
	ThirdTypeTikTokSellerAPI: "tiktoklib",
	ThirdTypeInstagram:       "instagram",
	ThirdTypeAmazon:          "amazon",
	ThirdTypeLazada:          "lazada",
}

// TikTokSellerAPIThirdBaseID TikTok 卖家 API 创始店铺标记
const TikTokSellerAPIThirdBaseID = "BASE"

// 商品库存映射值
const (
	InventoryOutOfStock    = "Out of stock"
	InventoryAdequateStock = "Adequate stock"
)

// 订单取消状态值
const (
	CancelReasonOverdue = "Buyer payment overdue"
	CancelReasonReview  = "Failed to pass risk review"
	CancelReasonOthers  = "others"
)

// SatisfactionEvaluationType 满意度评价支持的类型
var SatisfactionEvaluationType = []int{
	ThirdTypeShopify,
	ThirdTypeCustomWebsite,
}

// AggregateShopType 服务端消息聚合的店铺类型
var AggregateShopType = []int{
	ThirdTypeTikTok,
	ThirdTypeShopee,
	ThirdTypeLazada,
}

// AggregateShopTypeByAmazon 按 Amazon 聚合的店铺类型
var AggregateShopTypeByAmazon = []int{
	ThirdTypeAmazon,
}

// ProductCardIDLanguageMap 商品卡片 `product_id` 对应多语言
var ProductCardIDLanguageMap = map[string]string{
	"EN": "product_id",
	"FR": "identifiant_produit",
	"DE": "produkt_id",
	"ES": "id_producto",
	"AR": "معرف_المنتج",
	"IT": "id_prodotto",
	"PT": "id_produto",
	"VI": "mã_sản_phẩm",
	"ID": "id_produk",
	"TH": "รหัสผลิตภัณฑ์",
}

// ProductCardTitleLanguageMap 商品卡片 `title` 对应多语言
var ProductCardTitleLanguageMap = map[string]string{
	"EN": "product_title",
	"FR": "titre du produit",
	"DE": "Produkttitel",
	"ES": "título del producto",
	"AR": "عنوان المنتج",
	"IT": "titolo del prodotto",
	"PT": "título do produto",
	"VI": "tiêu đề sản phẩm",
	"ID": "judul produk",
	"TH": "ชื่อผลิตภัณฑ์",
}

// OrderCardIDLanguageMap 订单卡片 `order_id` 对应多语言
var OrderCardIDLanguageMap = map[string]string{
	"EN": "order_id",
	"FR": "Identifiant de commande",
	"DE": "Bestell-ID",
	"ES": "ID de orden",
	"AR": "معرف الطلب",
	"IT": "ID ordine",
	"PT": "ID do pedido",
	"VI": "ID đơn hàng",
	"ID": "ID pesanan",
	"TH": "รหัสคำสั่งซื้อ",
}

// VoucherCardIDLanguageMap 优惠券卡片 `voucher_id` 对应多语言
var VoucherCardIDLanguageMap = map[string]string{
	"EN": "voucher_id",
	"FR": "ID du coupon",
	"DE": "Gutschein-ID",
	"ES": "ID de cupón",
	"AR": "معرف القسيمة",
	"IT": "ID del coupon",
	"PT": "ID do cupão",
	"VI": "ID phiếu giảm giá",
	"ID": "ID kupon",
	"TH": "รหัสคูปอง",
}

// OrderReturnRefundCardLanguageMap 订单售后卡片 对应多语言
var OrderReturnRefundCardLanguageMap = map[string]string{
	"EN": "You can request a return/refund, order_id",
	"FR": "Vous pouvez demander un retour/remboursement, Identifiant de commande",
	"DE": "Sie können eine Rückgabe/Rückerstattung beantragen, Bestell-ID",
	"ES": "Puede solicitar una devolución/reembolso, ID de orden",
	"AR": "يمكنك تقديم طلب استرجاع / استرداد، معرف الطلب",
	"IT": "Puoi richiedere un reso/rimborso, ID ordine",
	"PT": "Você pode solicitar uma devolução/reembolso, ID do pedido",
	"VI": "Bạn có thể yêu cầu trả hàng/hoàn tiền, ID đơn hàng",
	"ID": "Anda dapat mengajukan pengembalian/pengembalian dana, ID pesanan",
	"TH": "คุณสามารถขอคืนสินค้า/คืนเงินได้ รหัสคำสั่งซื้อ",
}

// HistoryMessageSyncDays 历史会话同步天数
const HistoryMessageSyncDays = 14

// HistoryMessageLimitDays 历史会话消息同步天数
const HistoryMessageLimitDays = 30

// VoucherNationCurrency 各国家 货币符号 | 小数位数
var VoucherNationCurrency = map[string]map[string]interface{}{
	"MY": {"symbol": "RM", "decimals": 2},  // 马来西亚
	"ID": {"symbol": "Rp", "decimals": 0},  // 印度尼西亚
	"TH": {"symbol": "฿", "decimals": 2},   // 泰国
	"PH": {"symbol": "₱", "decimals": 2},   // 菲律宾
	"SG": {"symbol": "S$", "decimals": 2},  // 新加坡
	"VN": {"symbol": "₫", "decimals": 0},   // 越南
	"BR": {"symbol": "R$", "decimals": 2},  // 巴西
	"MX": {"symbol": "$", "decimals": 2},   // 墨西哥
	"CO": {"symbol": "$", "decimals": 2},   // 哥伦比亚
	"CL": {"symbol": "$", "decimals": 0},   // 智利
	"TW": {"symbol": "NT$", "decimals": 0}, // 台湾
}

// MinSyncConversationCount 最小同步会话数量
const MinSyncConversationCount = 300
