#! /usr/bin/env bash
CURDIR=$(pwd)

: ${ETCD_ADDR:="localhost:2379"}
export ETCD_ADDR

# 此处只涉及 Kitex，但是 Hertz 使用这个没有影响，保留即可
export KITEX_RUNTIME_ROOT=$CURDIR

# 这个 SERVICE 环境变量会自动地由 Dockerfile/Makefile 设置
exec "$CURDIR/output/$SERVICE/$SERVICE"