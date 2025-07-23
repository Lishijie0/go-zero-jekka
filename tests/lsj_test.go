package tests

import (
	"github.com/joho/godotenv"
	"jekka-api-go/pkg/config"
	"os"
	"testing"
)

func Test(t *testing.T) {
	cfg := config.NewConfig()
	//print(cfg.AppConfig.Mode)
	//
	return
	// 加载 .env 文件
	err := godotenv.Load("../.env") // 注意路径是否正确
	if err != nil {
		t.Fatalf("加载 .env 文件失败: %v", err)
	}
	host := os.Getenv("REDIS_HOST")
	pass := os.Getenv("REDIS_PASSWORD")

	print("host:", host, "pass:", pass) //rds, err := redis.NewRedis(conf)
	//if err != nil {
	//	t.Fatalf("创建 Redis 客户端失败: %v", err)
	//}
	//
	//// 创建 FreqUtil 实例
	//freq := tools.NewFreqUtil()
	//ctx := context.Background()
	//msg := "test-error-key"
	//
	//// 调用 Control 方法
	//count, err := freq.Control(ctx, msg, 60)
	//if err != nil {
	//	t.Errorf("Control 方法出错: %v", err)
	//}
	//
	//t.Logf("当前频率计数: %d", count)
}

//	func TestGenerateLogId(t *testing.T) {
//		//logId = fmt.Sprintf("%d%s", time.Now().UnixNano(), util.RandString(3))
//		//
//		//logId := fmt.Sprintf("%d%s", time.Now().UnixNano(), util.RandString(3))
//		var env = util.GetEnv()
//		fmt.Println(env)
//	}
//func TestFreqUtil_Control(t *testing.T) {
//	conf := redis.RedisConf{
//		Host: "127.0.0.1:6379",
//		Type: "node",
//		Pass: "MtpRedis011",
//	}
//
//	rds, err := redis.NewRedis(conf)
//	if err != nil {
//		t.Fatalf("创建 Redis 客户端失败: %v", err)
//	}
//
//	// 检查 Redis 是否连通
//	val, err := rds.Get("test")
//	if err != nil && !errors.Is(err, redis.Nil) {
//		t.Fatalf("Redis 连接失败: %v", err)
//	}
//	t.Logf("Redis 连接成功，test key 值为: %s", val)
//
//	ctx := context.Background()
//	freq := tools.NewFreqUtil(rds)
//	msg := "test-error-key"
//
//	for i := 0; i < 6; i++ {
//		count, err := freq.Control(ctx, msg, 60)
//		assert.NoError(t, err)
//		t.Logf("Call %d -> count: %d", i+1, count)
//	}
//}
