# 检查 tmux 是否存在
TMUX_EXISTS := $(shell command -v tmux)
# 当前架构
ARCH := $(shell uname -m)
PREFIX = "[Makefile]"
# 目录相关
DIR := $(shell pwd)
IDL_PATH := $(DIR)/idl
OUTPUT := $(DIR)/output
DOCKER_PATH := $(DIR)/docker
SCRIPT := $(DOCKER_PATH)/script
ENV_PATH := $(DOCKER_PATH)/env

# 项目 MODULE 名
MODULE := github.com/ant-02/myreel-plus

# tmux session名
SESSION := myreel

# 服务名
SERVICES := gateway auth

# hertz 生成
.PHONY: hz-%
hz-%:
	hz update -idl ${IDL_PATH}/api/$*.proto

# kitex 生成
.PHONY: kitex-gen-%
kitex-gen-%:
	@kitex -module "${MODULE}" \
		-I idl \
		${IDL_PATH}/$*.proto
	@go mod tidy

# 构建并运行服务
.PHONY: $(SERVICES)
$(SERVICES):
	@if [ -z "$(TMUX_EXISTS)" ]; then \
		echo "$(PREFIX) tmux is not installed. Please install tmux first."; \
		exit 1; \
	fi
	@if [ -z "$$TMUX" ]; then \
		echo "$(PREFIX) you are not in tmux, press ENTER to start tmux environment."; \
		read -r; \
		if tmux has-session -t $(SESSION) 2>/dev/null; then \
			echo "$(PREFIX) Tmux session '$(SESSION)' already exists. Attaching to session and running command."; \
			tmux attach-session -t $(SESSION); \
			tmux send-keys -t $(SESSION) "make $@" C-m; \
		else \
			echo "$(PREFIX) No tmux session found. Creating a new session."; \
			tmux new-session -s $(SESSION) "make $@; $$SHELL"; \
		fi; \
	else \
		echo "$(PREFIX) Build $@ target..."; \
		mkdir -p output; \
		bash $(SCRIPT)/build.sh $@; \
		echo "$(PREFIX) Build $@ target completed"; \
	fi
ifndef BUILD_ONLY
	@echo "$(PREFIX) Automatic run server"
	@if tmux list-windows -F '#{window_name}' | grep -q "^$@$$"; then \
		echo "$(PREFIX) Window '$@' already exists. Reusing the window."; \
		tmux select-window -t "$@"; \
	else \
		echo "$(PREFIX) Window '$@' does not exist. Creating a new window."; \
		tmux new-window -n "$@"; \
		tmux split-window -h ; \
		tmux select-layout -t "$@" even-horizontal; \
	fi
	@echo "$(PREFIX) Running $@ service in tmux..."
	@tmux send-keys -t $@.0 'export SERVICE=$@ && bash ./docker/script/entrypoint.sh' C-m
	@tmux select-pane -t $@.1
endif

.PHONY: env-up
env-up:
	@docker compose -f ./docker/compose.yml up -d