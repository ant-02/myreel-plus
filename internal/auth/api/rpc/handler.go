package handler

import (
	"context"

	"github.com/ant-02/myreel-plus/internal/auth/application"
	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/value"
	"github.com/ant-02/myreel-plus/kitex_gen/auth"
	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/response"
)

type AuthHandler struct {
	authAppService application.AuthAppService
}

func NewAuthController(aas application.AuthAppService) *AuthHandler {
	return &AuthHandler{authAppService: aas}
}

func (ah *AuthHandler) Login(ctx context.Context, req *auth.LoginRequest) (r *auth.LoginResponse, err error) {
	r = new(auth.LoginResponse)

	password, err := value.NewPasswordBuilder(req.Password).Build()
	if err != nil {
		r.Base = response.BuildBaseResp(err)
		return
	}

	err = ah.authAppService.Login(ctx, req.Identity, constants.IdentityType(req.IdentityType), *password)
	if err != nil {
		r.Base = response.BuildBaseResp(err)
	}

	r.Base = response.BuildSuccessResp()
	return
}
