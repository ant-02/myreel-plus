package utils

import (
	"errors"
	"net"
	"strings"

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
	if config.Service.AddrList == nil {
		return "", errno.Errorf(errno.InternalServiceErrorCode, "utils.GetAvailablePort: config.Service.AddrList is nil")
	}
	for _, addr := range config.Service.AddrList {
		if ok := AddrCheck(addr); ok {
			return addr, nil
		}
	}
	return "", errno.Errorf(errno.InternalServiceErrorCode, "utils.GetAvailablePort: not available port from config")
}

// GetMysqlDSN 会拼接 mysql 的 DSN
func GetMysqlDSN() (string, error) {
	if config.MySQL == nil {
		return "", errors.New("config not found")
	}

	dsn := strings.Join([]string{
		config.MySQL.Username, ":", config.MySQL.Password,
		"@tcp(", config.MySQL.Addr, ")/",
		config.MySQL.Database, "?charset=" + config.MySQL.Charset + "&parseTime=true",
	}, "")

	return dsn, nil
}
