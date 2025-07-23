package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/easycb/easycb-go"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"jekka-api-go/app/third/cmd/mq/internal/svc"
	"jekka-api-go/app/third/cmd/mq/jobtype"
	"jekka-api-go/pkg/db/model"
	"jekka-api-go/pkg/services/lib/third/tiktoklib"
	"jekka-api-go/pkg/types/thirdtype"
	"strconv"
)

type SyncWebhookOrderHandler struct {
	svcCtx *svc.ServiceContext
}

// NewSyncWebhookOrderHandler 路由调用的方法
func NewSyncWebhookOrderHandler(svcCtx *svc.ServiceContext) *SyncWebhookOrderHandler {
	return &SyncWebhookOrderHandler{
		svcCtx: svcCtx,
	}
}

func (l *SyncWebhookOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	fmt.Println("SyncWebhookOrder")
	logx.WithContext(ctx).Infof("%s: started", t.Type())
	var payload jobtype.WebhookOrderPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("解析任务数据失败: %v", err)
	}
	//defer NewSyncWebhookOrderHandler(l.svcCtx)
	var shop model.JkShop
	// 查询店铺信息
	fmt.Println("开始test")
	JkShop := l.svcCtx.DB.Table("jk_shop").Select("id", "name", "third_session").Where("id = ?", payload.JkShopId).First(&shop)
	if JkShop.RowsAffected == 0 {
		fmt.Println("record not found")
		return nil
	}

	// 设置三方平台配置
	config := thirdtype.AppConfig{
		ClientID:     l.svcCtx.Config.ThirdPlatformConfig.Tiktok.ClientID,
		ClientSecret: l.svcCtx.Config.ThirdPlatformConfig.Tiktok.ClientSecret,
	}

	// 实例化client
	base := tiktoklib.NewBaseTTService(shop, config)
	// 设置查询订单
	query := easycb.AnyMap{}.Set("ids", []string{payload.OrderId})
	var thirdSession tiktoklib.ThirdSessionTT
	_ = json.Unmarshal([]byte(shop.ThirdSession), &thirdSession)
	// fmt.Println(thirdSession.AccessToken, thirdSession.Shop.ShopCipher)
	// 配置查询参数
	queryConfig := tiktoklib.ApiQueryConfig{
		AccessToken: thirdSession.AccessToken,
		ShopCipher:  thirdSession.Shop.ShopCipher,
	}
	info := tiktoklib.GetOrderInfo(base, query, queryConfig)

	order := info.Data.Orders[0]
	price, _ := strconv.ParseFloat(order.Payment.TotalAmount, 64)
	//转换为
	thirdOrder := model.JkThirdOrder{
		JkShopID:          shop.ID,
		ThirdType:         shop.ThirdType,
		JkUserID:          payload.JkUserId,
		OrderID:           order.Id,
		Price:             price,
		OrderStatus:       order.Status,
		BuyerNote:         order.BuyerMessage,
		BuyerUserID:       order.UserId,
		BuyerEmail:        order.BuyerEmail,
		PaymentMethod:     order.PaymentMethodName,
		TrackingNumber:    order.TrackingNumber,
		ShippingProvider:  order.ShippingProvider,
		JkOperatorRecords: "{}", //todo后期有其他工具代替gentool,可以赋值nil
		CancelReason:      order.CancelReason,
		CreateTime:        order.CreateTime,
		UpdateTime:        order.UpdateTime,
	}

	result := l.svcCtx.DBThirdV2.Table("jk_third_order").Create(&thirdOrder)
	fmt.Printf("订单保存成功-----%v\n", thirdOrder.ID)
	fmt.Println(result.Error)
	extra := thirdtype.ThirdOrderExtraTypeByTT{
		SellerNote:            order.SellerNote,
		CancellationInitiator: order.CancellationInitiator,
		ShippingProvider:      order.ShippingProvider,
		Payment: struct {
			Tax                         string `json:"tax"`
			Currency                    string `json:"currency"`
			SubTotal                    string `json:"sub_total"`
			ProductTax                  string `json:"product_tax"`
			ShippingFee                 string `json:"shipping_fee"`
			TotalAmount                 string `json:"total_amount"`
			SellerDiscount              string `json:"seller_discount"`
			ShippingFeeTax              string `json:"shipping_fee_tax"`
			PlatformDiscount            string `json:"platform_discount"`
			OriginalShippingFee         string `json:"original_shipping_fee"`
			OriginalTotalProductPrice   string `json:"original_total_product_price"`
			RetailDeliveryFee           string `json:"retail_delivery_fee"`
			ShippingFeePlatformDiscount string `json:"shipping_fee_platform_discount"`
			ShippingFeeSellerDiscount   string `json:"shipping_fee_seller_discount"`
			SmallOrderFee               string `json:"small_order_fee"`
		}{
			Tax:                         order.Payment.Tax,
			Currency:                    order.Payment.Currency,
			SubTotal:                    order.Payment.SubTotal,
			ProductTax:                  order.Payment.ProductTax,
			ShippingFee:                 order.Payment.ShippingFee,
			TotalAmount:                 order.Payment.TotalAmount,
			SellerDiscount:              order.Payment.SellerDiscount,
			ShippingFeeTax:              order.Payment.ShippingFeeTax,
			PlatformDiscount:            order.Payment.PlatformDiscount,
			OriginalShippingFee:         order.Payment.OriginalShippingFee,
			OriginalTotalProductPrice:   order.Payment.OriginalTotalProductPrice,
			RetailDeliveryFee:           order.Payment.RetailDeliveryFee,
			ShippingFeePlatformDiscount: order.Payment.ShippingFeePlatformDiscount,
			ShippingFeeSellerDiscount:   order.Payment.ShippingFeeSellerDiscount,
			SmallOrderFee:               order.Payment.SmallOrderFee,
		},
		RecipientAddress: struct {
			AddressDetail string `json:"address_detail"`
			FullAddress   string `json:"full_address"`
			Name          string `json:"name"`
			PhoneNumber   string `json:"phone_number"`
			PostalCode    string `json:"postal_code"`
			RegionCode    string `json:"region_code"`
		}{
			AddressDetail: order.RecipientAddress.AddressDetail,
			FullAddress:   order.RecipientAddress.FullAddress,
			Name:          order.RecipientAddress.Name,
			PhoneNumber:   order.RecipientAddress.PhoneNumber,
			PostalCode:    order.RecipientAddress.PostalCode,
			RegionCode:    order.RecipientAddress.RegionCode,
		},
	}

	// 将结构体序列化为 JSON 字符串
	extraJSON, err := json.Marshal(extra)
	if err != nil {
		fmt.Println("Failed to marshal extra:", err)
	}
	extraString := string(extraJSON) // 将 []byte 转换为 string
	orderExtra := model.JkThirdOrderExtra{
		JkShopID:  shop.ID,
		ThirdType: shop.ThirdType,
		JkUserID:  payload.JkUserId,
		OrderID:   order.Id,
		Extra:     extraString,
	}

	resultExtra := l.svcCtx.DBThirdV2.Table("jk_third_order_extra").Create(&orderExtra)
	fmt.Printf("订单扩展保存结果-----%v\n", orderExtra.ID)
	fmt.Println(resultExtra.Error)

	//查询店铺
	fmt.Printf("webhookOrder队列消费成功-----%v\n", payload.JkShopId)
	return nil
}
