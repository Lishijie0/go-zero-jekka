package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"jekka-api-go/app/third/cmd/mq/internal/svc"
	"jekka-api-go/app/third/cmd/mq/jobtype"
	"jekka-api-go/pkg/db/model"
)

type SyncShopBasicDataHandler struct {
	svcCtx *svc.ServiceContext
}

// NewSyncShopOrderHandler 路由调用的方法
func NewSyncShopOrderHandler(svcCtx *svc.ServiceContext) *SyncShopBasicDataHandler {
	return &SyncShopBasicDataHandler{
		svcCtx: svcCtx,
	}
}

func (l *SyncShopBasicDataHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	// 同步店铺开始
	var shop model.JkShop
	// 查询店铺信息
	fmt.Println("开始test")
	JkShop := l.svcCtx.DB.Table("jk_shop").Select("id", "name", "third_session").Where("id = ?", 11815).First(&shop)
	if JkShop.RowsAffected == 0 {
		fmt.Println("record not found")
		return nil
	}

	//fmt.Println(JkShop.RowsAffected) // 返回找到的记录数
	//fmt.Println(JkShop.Error)        // returns error or nil
	// 同步订单
	//asynq.NewTask(jobtype.syncOrderSingle, payload),
	//asynq.Queue(constant.QueueHigh), // 投递到队列 “high”
	//asynq.ProcessIn(time.Second*2),  // 10s秒后执行
	// service tt/se/lzd
	// 同步退款

	// 同步优惠券

	// 同步表情

	//asynq.NewTask(jobtype2.SyncShopOrder, payload),
	//asynq.Queue(constant.QueueHigh), // 投递到队列 “high”
	//asynq.ProcessIn(time.Second*2),  // 10s秒后执行
	var user model.JkUser //引入表模型
	logx.WithContext(ctx).Infof("%s: started", t.Type())
	var payload jobtype.SyncShopBasicPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("解析任务数据失败: %v", err)
	}

	//var jkUserId = util.GetDatabaseName("test", payload.JkUserId)
	//fmt.Println(jkUserId)

	// 查询用户信息
	JkUser := l.svcCtx.DB.Model(&model.JkUser{}).Select("id", "username").Where("id = ?", payload.JkUserId).First(&user)
	if JkUser.Error != nil {
		return fmt.Errorf("查询用户失败: %v", JkUser.Error)
	}

	fmt.Println(user.Username)
	// 将结构体转换为 JSON 字符串
	//userJSON, err := json.Marshal(user)
	//if err != nil {
	//	log.Fatalf("转换为 JSON 失败: %v", err)
	//}
	//
	//// 进一步解析 JSON 对象
	//var jsonData map[string]interface{}
	//if err := json.Unmarshal(user, &jsonData); err != nil {
	//	fmt.Println("解析 JSON 失败:", err)
	//	return nil
	//}

	// 查询店铺信息
	//JkShop := l.svcCtx.DB.Table("jk_shop").Select("id", "name").Where("jk_user_id = ?", payload.JkUserId).First(&shop)
	//if JkShop.Error != nil {
	//	return fmt.Errorf("查询店铺失败: %v", JkShop.Error)
	//}
	//fmt.Println(shop.Name)

	// 查询订单信息
	//dbName := util.GetDatabaseName("jk_third_order", payload.JkUserId)
	//fmt.Println(dbName)

	var orders []model.JkThirdOrder
	l.svcCtx.DBThirdV2.Model(&model.JkThirdOrder{}).Where("jk_user_id", int64(3)).Debug().Find(&orders)
	fmt.Printf("1%#v\n", orders)
	l.svcCtx.DBThirdV2.Model(&model.JkThirdOrder{}).Where("jk_user_id", int64(2)).Debug().Find(&orders)
	fmt.Printf("2%#v\n", orders)
	//thirdOrder := l.svcCtx.DBThirdV2.Model(&model.JkThirdOrder{}).Select("id").Where("jk_user_id", int64(2)).First(&order)
	//if thirdOrder.Error != nil {
	//	return fmt.Errorf("查询订单失败: %v", thirdOrder.Error)
	//}
	//fmt.Println(order.ID)
	//thirdOrder := l.svcCtx.DBThirdV2.Where("jk_user_id = ?", payload.JkUserId).Where("jk_shop_id = ?", payload.JkShopId).First(&order)
	//if thirdOrder.Error != nil {
	//	return fmt.Errorf("查询订单失败: %v", thirdOrder.Error)
	//}
	//fmt.Println(order.ID)
	// 根据店铺类型再次放入不同的队列

	// 查询店铺
	fmt.Printf("队列消费成功-----%v\n", payload.JkShopId)
	return nil
}
