package thirdtype

//type ExtraPaymentByTT struct {
//	Tax                         string `json:"tax"`
//	Currency                    string `json:"currency"`
//	SubTotal                    string `json:"sub_total"`
//	ProductTax                  string `json:"product_tax"`
//	ShippingFee                 string `json:"shipping_fee"`
//	TotalAmount                 string `json:"total_amount"`
//	SellerDiscount              string `json:"seller_discount"`
//	ShippingFeeTax              string `json:"shipping_fee_tax"`
//	PlatformDiscount            string `json:"platform_discount"`
//	OriginalShippingFee         string `json:"original_shipping_fee"`
//	OriginalTotalProductPrice   string `json:"original_total_product_price"`
//	RetailDeliveryFee           string `json:"retail_delivery_fee"`
//	ShippingFeePlatformDiscount string `json:"shipping_fee_platform_discount"`
//	ShippingFeeSellerDiscount   string `json:"shipping_fee_seller_discount"`
//	SmallOrderFee               string `json:"small_order_fee"`
//}
//
//type ExtraRecipientAddressByTT struct {
//	AddressDetail string `json:"address_detail"`
//	FullAddress   string `json:"full_address"`
//	Name          string `json:"name"`
//	PhoneNumber   string `json:"phone_number"`
//	PostalCode    string `json:"postal_code"`
//	RegionCode    string `json:"region_code"`
//}

// ThirdOrderExtraTypeByTT 三方订单扩展字段
type ThirdOrderExtraTypeByTT struct {
	SellerNote            string `json:"seller_note"`
	CancellationInitiator string `json:"cancellation_initiator"`
	ShippingProvider      string `json:"shipping_provider"`
	Payment               struct {
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
	} `json:"payment"`
	RecipientAddress struct {
		AddressDetail string `json:"address_detail"`
		FullAddress   string `json:"full_address"`
		Name          string `json:"name"`
		PhoneNumber   string `json:"phone_number"`
		PostalCode    string `json:"postal_code"`
		RegionCode    string `json:"region_code"`
	} `json:"recipient_address"`
}
