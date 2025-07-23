package util

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"bytes"
	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/logx"
	"jekka-api-go/pkg/services/util/tools"
	"net/http"
)

const timeFormat = "2006-01-02 15:04:05.000"

var logLevelMap = map[string]int{
	"DEBUG": 0,
	"INFO":  1,
	"WARN":  2,
	"ERROR": 3,
}

type Logger struct {
	id       string
	freqUtil *tools.FreqUtil
}

func init() {
	_ = godotenv.Load() // 加载.env
	logx.MustSetup(logx.LogConf{
		ServiceName: "jekka-service",
		Mode:        "file",
		Path:        "logs",
		Level:       os.Getenv("LOG_LEVEL"),
		KeepDays:    7,
	})
}

func NewLogger(id ...string) *Logger {
	loggerId := fmt.Sprintf("L%s", time.Now().Format("150405000"))
	if len(id) > 0 {
		loggerId = id[0]
	}
	return &Logger{
		id:       loggerId,
		freqUtil: tools.NewFreqUtil(),
	}
}

func (l *Logger) log(level, msg string, args map[string]interface{}) {
	if !l.checkLevel(level) {
		return
	}
	timestamp := time.Now().Format(timeFormat)
	sid := ""
	if s, ok := args["sid"]; ok {
		sid = fmt.Sprintf("[%v]", s)
		delete(args, "sid")
	}
	logStr := fmt.Sprintf("[%s] %s %s_%s%s", timestamp, level, l.id, sid, msg)
	if len(args) > 0 {
		argBytes, _ := json.Marshal(args)
		logStr += " " + string(argBytes)
	}

	fmt.Println(logStr)
	ctx := context.Background()
	logx.WithContext(ctx).Info(logStr)
}

func (l *Logger) Debug(msg string, args map[string]interface{}) { l.log("DEBUG", msg, args) }
func (l *Logger) Info(msg string, args map[string]interface{})  { l.log("INFO", msg, args) }
func (l *Logger) Warn(msg string, args map[string]interface{})  { l.log("WARN", msg, args) }

func (l *Logger) Error(msg string, err error, args map[string]interface{}) {
	l.log("ERROR", msg, args)
	if l.freqUtil != nil {
		key := fmt.Sprintf("log:alert:%s", msg)
		count, _ := l.freqUtil.Control(context.Background(), key, 60)
		if count > 5 {
			return
		}
	}
	l.sendFeishu(msg, err, args)
}

func (l *Logger) sendFeishu(msg string, err error, args map[string]interface{}) {
	trace := []string{msg, err.Error()}
	if args != nil {
		for k, v := range args {
			trace = append(trace, fmt.Sprintf("%s: %v", k, v))
		}
	}
	text := strings.Join(trace, "\n")
	hook := os.Getenv("FEISHU_HOOK")
	if hook == "" {
		return
	}
	payload := map[string]interface{}{
		"msg_type": "text",
		"content":  map[string]string{"text": text},
	}
	body, _ := json.Marshal(payload)
	http.Post(hook, "application/json", bytes.NewReader(body))
}

func (l *Logger) checkLevel(level string) bool {
	envLevel := os.Getenv("LOG_LEVEL")
	if envLevel == "" {
		envLevel = "INFO"
	}
	return logLevelMap[strings.ToUpper(level)] >= logLevelMap[strings.ToUpper(envLevel)]
}
