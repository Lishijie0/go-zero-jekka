package cache

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	// "gorm.io/gorm"
)

const (
	// ShopInfo 店铺信息
	ShopInfo = "shop:info:%s"
	// ShopBindCode 认证绑定码
	ShopBindCode = "shop:bind:code:%s"
	// ShopInfoByThird 店铺信息
	ShopInfoByThird = "shop:info:by:thirdtype:%s:%s"
	// ShopIds 用户的店铺ID列表（用于店铺聚合）
	ShopIds = "shop:ids:%s:%s"
	// ShopOrderSubTaskRemainingCount 以下两个key均用于新店同步订单完成后发送通知用
	ShopOrderSubTaskRemainingCount = "shop:order:sync:sub_task:remaining:%s"
	ShopOrderSyncTotal             = "shop:order:sync:total:%s"
)

// ShopCache 结构体
type ShopCache struct {
	*BaseCache
	db *sqlx.SqlConn
}

// NewShopCache 创建一个新的 ShopCache 实例
func NewShopCache(redisAddr string, password string, db *sqlx.SqlConn) *ShopCache {
	return &ShopCache{
		BaseCache: NewBaseCache(redisAddr, password),
		db:        db,
	}
}

// GetKey 返回格式化的缓存键
func (sc *ShopCache) GetKey(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

/*// rememberShop 获取店铺信息
func (sc *ShopCache) rememberShop(id int64) (*model.JkShop, error) {
	if id == 0 {
		return nil, nil
	}

	key := sc.GetKey(ShopInfo, id)
	var shopInfo *model.JkShop
	err := sc.client.GetObj(sc.ctx, key, &shopInfo, func() (interface{}, error) {
		return shop.NewShopModel(sc.db).FindOne(id)
	}, time.Duration(TtlTenMinute)*time.Second)
	if err != nil {
		return nil, err
	}
	return shopInfo, nil
}

// forgetShop 删除店铺信息
func (sc *ShopCache) forgetShop(id int64) error {
	key := sc.GetKey(ShopInfo, id)
	return sc.client.Del(sc.ctx, key)
}*/
