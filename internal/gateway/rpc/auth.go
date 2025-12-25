package rpc

import (
	"context"

	"github.com/ant-02/myreel-plus/kitex_gen/auth"
	"github.com/ant-02/myreel-plus/pkg/client"
	"github.com/ant-02/myreel-plus/pkg/errno"
	"github.com/ant-02/myreel-plus/pkg/response"
	"github.com/bytedance/gopkg/util/logger"
)

func InitAuthClient() {
	c, err := client.InitAuthRPC()
	if err != nil {
		logger.Fatalf("api.rpc.auth InitAuthRPC failed, err is %v", err)
	}

	authClient = *c
}

func LoginRPC(ctx context.Context, req *auth.LoginRequest) error {
	resp, err := authClient.Login(ctx, req)
	if err != nil {
		logger.Errorf("LoginRPC: RPC called failed: %v", err)
		return errno.InternalServiceError.WithError(err)
	}

	if !response.IsSuccess(resp.Base.Code) {
		return errno.InternalServiceError.WithError(err)
	}
	return nil
}
