package main

import (
	"net"

	"github.com/ant-02/myreel-plus/config"
	"github.com/ant-02/myreel-plus/internal/auth"
	"github.com/ant-02/myreel-plus/kitex_gen/auth/authservice"
	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/utils"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func init() {
	config.Init(constants.AuthServiceName)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		logger.Fatalf("Auth: new etcd registry failed, err: %v", err)
	}

	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		logger.Fatalf("Auth: get available port failed, err: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("Auth: resolve tcp addr failed, err: %v", err)
	}

	svr := authservice.NewServer(
		auth.InjectAuthHandler(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.AuthServiceName,
		}),
	)

	if err = svr.Run(); err != nil {
		logger.Fatalf("Auth: run server failed, err: %v", err)
	}
}
