package util

import (
	"net"

	"github.com/ant-02/myreel-plus/config"
	"github.com/ant-02/myreel-plus/pkg/errno"
	"github.com/bytedance/gopkg/util/logger"
)

// AddrCheck 会检查当前的监听地址是否已被占用
func AddrCheck(addr string) bool {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	defer func() {
		if err := l.Close(); err != nil {
			logger.Errorf("utils.AddrCheck: failed to close listener: %v", err.Error())
		}
	}()
	return true
}

// GetAvailablePort 会尝试获取可用的监听地址
func GetAvailablePort() (string, error) {
	if AddrCheck(config.Service.Addr) {
		return config.Service.Addr, nil
	}
	return "", errno.Errorf(nil, "utils.GetAvailablePort: not available port from config")
}
