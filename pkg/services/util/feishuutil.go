package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type FeiShuUtil struct {
	mu sync.Mutex
}

// 常量定义
const (
	AtAll                            = "all"
	AtUserIdJiangTao                 = "ou_d92c0d87605dac2c41ac3e9a75d49abd"
	URLExceptionNotify               = "https://open.feishu.cn/open-apis/bot/v2/hook/9347caaa-9e20-4b6b-b081-29fddd5433aa"
	URLExceptionNotifyBackstage      = "https://open.feishu.cn/open-apis/bot/v2/hook/c7a73414-1169-41bf-9b40-02c8d74d7b1d"
	URLExceptionNotifyThird          = "https://open.feishu.cn/open-apis/bot/v2/hook/52b56ba8-e2ea-4fa2-83fa-c16634ff2964"
	URLExceptionNotifyDev            = "https://open.feishu.cn/open-apis/bot/v2/hook/bc07c7ae-8462-47d7-b993-6596f2bbf15f"
	URLOptimizeNotify                = "https://open.feishu.cn/open-apis/bot/v2/hook/0ec0b839-f822-4e74-b14e-51f7cdff1a4f"
	URLNewShopNotify                 = "https://open.feishu.cn/open-apis/bot/v2/hook/60b550e0-080e-40ae-b770-d7e37ab745bf"
	URLDelShopNotify                 = "https://open.feishu.cn/open-apis/bot/v2/hook/490c31cf-1905-46c1-af94-ae1e8b261069"
	URLNewMerchantNotify             = "https://open.feishu.cn/open-apis/bot/v2/hook/b92c1f96-1260-4c54-a31a-a1eaeecca4ff"
	URLGoogleMapSearchResult         = "https://open.feishu.cn/open-apis/bot/v2/hook/8b2382f6-6933-42f2-a38b-50d794192ad0"
	URLOrderOrProductLimitNotify     = "https://open.feishu.cn/open-apis/bot/v2/hook/60ac8e30-508f-4ebb-914d-938c5f77fde7"
	URLShopCommentOverrunNotify      = "https://open.feishu.cn/open-apis/bot/v2/hook/111e51b5-b965-451a-a5f8-f791cff723f0"
	URLOrderFollowExceptionNotify    = "https://open.feishu.cn/open-apis/bot/v2/hook/9b38bed1-5965-433f-96c1-e443ff2e1e74"
	URLOrderFollowBizExceptionNotify = "https://open.feishu.cn/open-apis/bot/v2/hook/320d4f94-dfb5-41e4-99b5-791fd29e0461"
	URLShopAuthExpiredNotify         = "https://open.feishu.cn/open-apis/bot/v2/hook/a3914682-6155-4416-a0fe-792a9adfbcdf"
	URLMailAvgDownloadTimeNotify     = "https://open.feishu.cn/open-apis/bot/v2/hook/653a4fca-ff93-4777-a3fa-ec02aabb3aba"
	URLOfficialWebsiteNotify         = "https://open.feishu.cn/open-apis/bot/v2/hook/c045cb60-b23c-4785-953d-55a67d7dc20f"
	URLPlanOrderNotify               = "https://open.feishu.cn/open-apis/bot/v2/hook/ddec1625-01f9-4e8c-b171-f3a50e24397b"
	URLRequestActiveAiNotify         = "https://open.feishu.cn/open-apis/bot/v2/hook/bafdee0b-3e55-4648-8e4a-22fbada8b223"
	URLKaCustomerExceptionNotify     = "https://open.feishu.cn/open-apis/bot/v2/hook/696df77e-1b2f-4aaa-83d7-e7811773384e"
	URLSyncProductCompleteNotify     = "https://open.feishu.cn/open-apis/bot/v2/hook/ceba5be7-84fd-4b9b-9a97-bdc63f5ed518"
	URLSyncOrderCompleteNotify       = "https://open.feishu.cn/open-apis/bot/v2/hook/6239d818-f93b-4872-9757-336f9d652736"
)

// 获取 FeiShuUtil 的单例
var feiShuUtilInstance *FeiShuUtil

func GetFeiShuUtil() *FeiShuUtil {
	if feiShuUtilInstance == nil {
		feiShuUtilInstance = &FeiShuUtil{}
	}
	return feiShuUtilInstance
}

// SendTextMsg 发送文本消息
func (f *FeiShuUtil) SendTextMsg(url string, text string, envName string, atList []string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("发送飞书通知时发生异常: %v", r)
		}
	}()
	envName = GetEnvName(envName)
	text = fmt.Sprintf("【%s】%s", envName, text) //🔔
	content := map[string]interface{}{
		"text": text,
	}
	data := map[string]interface{}{
		"msg_type": "text",
		"content":  content,
	}

	for _, at := range atList {
		if at == AtAll {
			// 将 "text" 字段类型转换为 string，并进行拼接
			content["text"] = content["text"].(string) + `<at user_id="all">所有人</at>`
		} else {
			content["text"] = content["text"].(string) + fmt.Sprintf(
				`<at user_id="%s">%s</at>`, at, AtUserIdNameMap(at))
		}
	}

	options := &http.Client{}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("JSON 序列化失败: %v", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("创建请求失败: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := options.Do(req)
	if err != nil {
		log.Printf("飞书通知发送请求失败: %v", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	// 记录响应或错误
	log.Printf("飞书通知发送结果: %v", res)
}

// SendPostMsg 发送富文本信息
func (f *FeiShuUtil) SendPostMsg(params map[string]interface{}, url string, atList []string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("发送富文本消息时发生异常: %v", r)
		}
	}()

	// 处理 @ 列表
	var content []interface{}
	for _, at := range atList {
		if at == AtAll {
			content = append(content, map[string]interface{}{
				"tag":     "at",
				"user_id": "all",
			})
		} else {
			content = append(content, map[string]interface{}{
				"tag":     "at",
				"user_id": at,
			})
		}
	}

	// 将 @ 信息添加到参数中
	if _, ok := params["content"]; ok {
		if contentArray, ok := params["content"].([]interface{}); ok {
			content = append(contentArray, content...)
		}
	}
	params["content"] = content

	// 创建消息体
	messageBody := map[string]interface{}{
		"msg_type": "post",
		"content":  map[string]interface{}{"post": map[string]interface{}{"zh_cn": params}},
	}

	// JSON 序列化消息体
	jsonData, err := json.Marshal(messageBody)
	if err != nil {
		log.Printf("JSON 序列化失败: %v", err)
		return
	}

	// 发送请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("创建请求失败: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("飞书通知发送请求失败: %v", err)
		return
	}
	defer res.Body.Close()

	// 记录响应
	log.Printf("飞书通知发送结果: %v", res.Status)
}

// SendExceptionNotify 发送异常通知
func (f *FeiShuUtil) SendExceptionNotify(text, envName string, atList []string) {
	if envName != "pro" {
		f.SendTextMsg(URLExceptionNotifyDev, text, envName, atList)
		return
	}
	f.SendTextMsg(URLExceptionNotify, text, envName, atList)
}

// SendSyncOrderCompleteNotify 发送其他类型的通知
func (f *FeiShuUtil) SendSyncOrderCompleteNotify(text, envName string, atList []string) {
	if envName != "pro" {
		return
	}
	f.SendTextMsg(URLSyncOrderCompleteNotify, text, envName, atList)
}

// 示例功能继续添加...

// AtUserIdNameMap 辅助函数：用户 ID 名称映射
func AtUserIdNameMap(userId string) string {
	userMap := map[string]string{
		AtUserIdJiangTao: "江涛",
	}

	return userMap[userId]
}
