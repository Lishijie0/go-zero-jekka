package monitor

import (
	"errors"
	"fmt"
	"jekka-api-go/app/third/cmd/mq/internal/config"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

func StartMonitor(c config.Config) error {
	if !c.Monitor.Enabled {
		return nil
	}

	// 创建监控处理器
	h := asynqmon.New(asynqmon.Options{
		RootPath: c.Monitor.Path,
		RedisConnOpt: asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
			DB:       c.Redis.DB,
		},
	})

	// 创建带认证的中间件处理器
	authHandler := basicAuth(h, c.Monitor.Name, c.Monitor.Pass) // 请修改用户名和密码

	// 启动 HTTP 服务
	mux := http.NewServeMux()
	mux.Handle(h.RootPath()+"/", authHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.Monitor.Port),
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	return nil
}

// 添加基本认证中间件
func basicAuth(next http.Handler, username, password string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != username || pass != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
