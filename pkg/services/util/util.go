package util

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"os"
	"reflect"
	"strings"
)

// GetEnv 获取环境
func GetEnv() string {
	type Config struct {
		Mode string
	}
	flag.Parse()
	var f = flag.String("f", "config.yaml", "config file")
	var c Config
	conf.MustLoad(*f, &c)
	return c.Mode
}

// IsJSON isJSON 检查数据是否为有效的 JSON
func IsJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// GetEnvName 辅助函数：获取环境名称
func GetEnvName(envName string) string {
	envMap := map[string]string{
		"local": "本地",
		"dev":   "开发",
		"pre":   "预发",
		"prod":  "线上",
	}
	if name, exists := envMap[envName]; exists {
		return name
	}
	return envName
}

// Empty 判断一个值是否为空
func Empty(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Slice, reflect.Map, reflect.Chan:
		return v.IsNil() || v.Len() == 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return v.IsNil()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Struct:
		return v.NumField() == 0
	default:
		return value == nil
	}
}

// RandString 生成指定长度的 URL 安全的随机字符串（字母、数字）
func RandString(length int) string {
	if length <= 0 {
		return ""
	}

	var result strings.Builder
	result.Grow(length)

	for {
		// 每次生成至少 ceil(length / 3) * 3 字节的随机数据
		remaining := length - result.Len()
		bytesSize := (remaining + 2) / 3 * 3 // 确保是 3 的倍数
		randBytes := make([]byte, bytesSize)

		if _, err := rand.Read(randBytes); err != nil {
			panic(err) // 或者根据需要处理错误
		}

		// Base64 编码，并移除 URL 不安全字符
		encoded := base64.URLEncoding.EncodeToString(randBytes)
		encoded = strings.ReplaceAll(encoded, "-", "") // 移除 URL 不安全字符
		encoded = strings.ReplaceAll(encoded, "_", "") // 可根据需要扩展

		// 截取所需长度
		result.WriteString(encoded[:remaining])

		if result.Len() == length {
			break
		}
	}

	return result.String()
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true // 文件存在
	}
	if os.IsNotExist(err) {
		return false // 文件不存在
	}
	// 其他错误，如权限问题等
	fmt.Printf("检查文件出错: %v\n", err)
	return false
}
