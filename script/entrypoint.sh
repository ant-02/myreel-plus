#! /usr/bin/env bash
CURDIR=$(pwd)

if [ -n "$ENV_FILE" ] && [ -f "$ENV_FILE" ]; then
  echo "Loading env from $ENV_FILE"
  set -a
  source "$ENV_FILE"
  set +a
else
  echo "ENV_FILE not set or file not found, skip loading env"
fi

# 这个 SERVICE 环境变量会自动地由 Dockerfile/Makefile 设置
exec "$CURDIR/output/$SERVICE/$SERVICE"