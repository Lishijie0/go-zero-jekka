package third

import (
	"fmt"
	"jekka-api-go/pkg/db/model"
	"jekka-api-go/pkg/services/util"
	"log"
)

// ErrConst 存储常量错误
type ErrConst struct {
	ParamError    string
	NotExists     string
	ParamErrorMsg string
	NotExistsMsg  string
}

// ApiException 表示API异常
type ApiException struct {
	Code    string
	Message string
}

func (e *ApiException) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}

// LogUtil 用于日志记录
type LogUtil struct{}

func (lu *LogUtil) Error(method string, msg string, context interface{}) {
	log.Printf("Error in %s: %s, context: %+v", method, msg, context)
}

// ShopCache 用于缓存店铺信息
type ShopCache struct{}

func (sc *ShopCache) RememberShop(jkShopId int64) *model.JkShop {
	// 这里应实现从缓存获取店铺信息的逻辑
	// 返回 nil 表示未找到
	return nil
}

// BaseThirdService 表示第三方服务的基类
type BaseThirdService struct {
	Shop      model.JkShop
	ThirdType int
	AppConfig third.AppConfig
}

// NewBaseThirdService 创建一个新的 BaseThirdService 实例
func NewBaseThirdService(jkShopId int64, thirdType int, config third.AppConfig) (*BaseThirdService, error) {
	// 判断config
	if util.Empty(config) {
		return nil, &ApiException{Code: ErrConst{}.ParamError, Message: fmt.Sprintf(ErrConst{}.ParamErrorMsg, "config")}
	}

	service := &BaseThirdService{
		Shop: model.JkShop{}, // 创建 ShopModel 实例
	}

	if jkShopId == 0 {
		// 认证阶段，无店铺信息
		service.Shop.ID = jkShopId
		service.ThirdType = thirdType
	} else {
		if jkShopId < 0 {
			// 参数错误
			logUtil := &LogUtil{}
			logUtil.Error("NewBaseThirdService", "参数错误", jkShopId)
			return nil, &ApiException{Code: ErrConst{}.ParamError, Message: fmt.Sprintf(ErrConst{}.ParamErrorMsg, "jkShopId")}
		}

		// 认证后，可直接获取店铺信息
		shop := (&ShopCache{}).RememberShop(jkShopId)
		if shop == nil {
			logUtil := &LogUtil{}
			logUtil.Error("NewBaseThirdService", "店铺信息不存在", jkShopId)
			return nil, &ApiException{Code: ErrConst{}.NotExists, Message: fmt.Sprintf(ErrConst{}.NotExistsMsg, "Store")}
		}

		// 复制店铺信息
		service.Shop.ID = shop.ID
		service.Shop.JKUserID = shop.JKUserID
		service.Shop.ThirdShopID = shop.ThirdShopID
		service.Shop.ThirdSession = shop.ThirdSession
		service.Shop.ThirdType = shop.ThirdType
		service.Shop.LangCode = shop.LangCode
		service.Shop.CountryCode = shop.CountryCode
		service.Shop.ThirdAppID = shop.ThirdAppID
		service.Shop.CreatedAt = shop.CreatedAt
		service.ThirdType = shop.ThirdType
	}
	service.AppConfig = config
	return service, nil
}
