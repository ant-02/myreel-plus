package rpc

import "github.com/ant-02/myreel-plus/kitex_gen/auth/authservice"

var (
	authClient authservice.Client
)

func Init() {
	InitAuthClient()
}
