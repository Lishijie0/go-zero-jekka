package shop

// 三方应用ID
const (
	ThirdAppIDTTOld     = 1 // TT老应用
	ThirdAppIDTTNew     = 2 // TT新应用
	ThirdAppIDTTUS      = 3 // 美国市场
	ThirdAppIDShopeeERP = 4 // Shopee erp应用
)

// 三方插件状态
const (
	ThirdPluginStatusUnknown = 0 // 未知
	ThirdPluginStatusOpened  = 1 // 打开
	ThirdPluginStatusClosed  = 2 // 关闭
)

// 店铺授权状态
const (
	ShopStatusUnfinished = 0 // 未完成授权
	ShopStatusEffecting  = 1 // 授权生效中
	ShopStatusExpiring   = 2 // 授权即将过期
	ShopStatusExpired    = 3 // 授权已过期
	ShopStatusAbnormal   = 4 // 授权异常
)

// ShopExceptionStatusList 店铺异常状态列表
var ShopExceptionStatusList = []int{
	ShopStatusAbnormal,
}

// 店铺状态
const (
	StatusOn  = 1 // 可用
	StatusOff = 0 // 禁用
)

// 数据来源
const (
	SourceCreate = "create"
	SourceUpdate = "update"
)

// ShopExpiringTimeMinutes 店铺授权过期时间设置 (15天)
const ShopExpiringTimeMinutes = 60 * 24 * 15

// 显示设置
const (
	ViewPromptNo  = 0 // 否
	ViewPromptYes = 1 // 是
)

// AI工作时段设置
const (
	WorkingHoursClosed = 0 // 关闭
	WorkingHoursOpened = 1 // 打开
)

// MaxShopTicketCreationNotificationCount 每个店铺每天最多收到的工单创建通知邮件数量
const MaxShopTicketCreationNotificationCount = 10

// 是否KA店铺
const (
	IsKANo  = 0 // 否
	IsKAYes = 1 // 是
)

// 是否BPO店铺
const (
	IsBPONo  = 0 // 否
	IsBPOYes = 1 // 是
)

// 是否活跃店铺
const (
	IsActiveNo  = 0 // 否
	IsActiveYes = 1 // 是
)

// CountryCodeUS 国家代码
const CountryCodeUS = "US"

// 前端会话分析类型
const (
	ConversationAnalysisHandledByJekka        = 2 // jekka会话
	ConversationAnalysisTransferredToOperator = 3 // 人工介入
)

// 新店流程步骤数
const (
	TutorialStepsUnknown    = -1  // 未知
	TutorialStepsNotStarted = 0   // 未开始
	TutorialStepsCompleted  = 999 // 已完成
)

// Shopee店铺授权状态
const (
	ShopeeShopAuthChat     = 1 // 只授权了消息
	ShopeeShopAuthERP      = 2 // 只授权了erp
	ShopeeShopAuthComplete = 3 // 授权完成
)

// 特殊国家代码
const (
	CountryCodeVN = "VN"
	CountryCodeBR = "BR"
	CountryCodeID = "ID"
	CountryCodeMY = "MY"
	CountryCodePH = "PH"
	CountryCodeSG = "SG"
	CountryCodeTH = "TH"
)

// 消息撤回时间限制
const (
	CancelMessageTimeLimitTen   = 10 // 消息撤回时间限制10分钟
	CancelMessageTimeLimitThree = 3  // 消息撤回时间限制3分钟
)
