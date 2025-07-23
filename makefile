.PHONY: all api api_gen rpc rpc_gen dockerfile up

# 解析参数
p1 := $(word 2, $(MAKECMDGOALS))
p2 := $(strip $(word 3, $(MAKECMDGOALS)))

# 兼容 make 目标解析，避免将参数解析为目标
%:
	@:

# 快速rebase同步最新master代码，遇到git push推送远端分支时报错(failed to push some refs to)，请执行 git pull --rebase origin $$branch
rebase:
	export branch=`git branch | grep \* | grep -Eo ' .+'` && \
		git checkout master && \
		git pull --rebase && \
		git checkout $$branch && \
		git rebase master;

# 测试环境，合并代码到dev分支
dev:
	- git branch -D dev;
	git fetch;
	export branch=`git branch | grep \* | grep -Eo ' .+'` && \
		echo "Current branch: $$branch" && \
		git checkout dev && \
		git pull --rebase && \
		git merge origin/master && \
		echo "merge: \033[0;31morigin/master\033[0m" && \
		git merge $$branch && \
		echo "merge: \033[0;31m$$branch\033[0m" && \
		git push && \
		git checkout $$branch;

# 创建 API 服务 .api 文件 (eg: make api user)
api:
	mkdir -p app/$(p1)/cmd/api/desc
	cd app/$(p1)/cmd/api/desc && \
	goctl api --o $(p1).api

# 根据 .api 文件生成 HTTP 服务代码 (eg: make api_gen user)
api_gen:
	cd app/$(p1)/cmd/api/desc && \
	goctl api go -api $(p1).api -dir ../ --style=gozero

# 创建 RPC 服务 .proto 文件 (eg: make rpc product)
rpc:
	mkdir -p app/$(p1)/cmd/rpc/pb
	cd app/$(p1)/cmd/rpc/pb && \
	goctl rpc --o $(p1).proto

# 根据 .proto 文件生成 RPC 代码 (eg: make rpc_gen product)
rpc_gen:
	cd app/$(p1)/cmd/rpc/pb && \
	goctl rpc protoc $(p1).proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=gozero && \
	cd ../$(p1) && \
	sed -i "" 's/,omitempty//g' $(p1).pb.go
    # linux: sed -i 's/,omitempty//g' *.pb.go\

# 根据表名生成model文件，如果要切换数据库，手动更换数据库链接信息 (eg: make model jk_third_product)
model:
	# 使用 gentool 生成模型文件
	gentool -dsn "admin:ZK9Vx49QSTs07u6P@tcp(54.213.111.192:3306)/jekka_third_v2?charset=utf8mb4&parseTime=True&loc=Local" \
		-tables "$(p1)" \
		-onlyModel \
		-modelPkgName "pkg/db/model" \
	# 使用 mv 命令重命名生成的文件
	# mv "pkg/db/model/$(p1).gen.go" "pkg/db/model/$(shell echo "$(p1)" | tr -d '_').gen.go"

# 以 Docker Compose 方式启动 (make up)
up:
	docker compose up -d
