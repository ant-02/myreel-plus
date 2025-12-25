package application

import (
	"context"

	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/value"
	"github.com/ant-02/myreel-plus/internal/auth/domain/service"
	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
)

type authAppService struct {
	as              service.AuthenticationService
	identityFactory *value.IdentityFactory
}

type AuthAppService interface {
	Login(ctx context.Context, rawValue string, identityType constants.IdentityType, password value.Password) error
}

func NewAuthAppService(as service.AuthenticationService, iF *value.IdentityFactory) AuthAppService {
	return &authAppService{
		as:              as,
		identityFactory: iF,
	}
}

func (aas *authAppService) Login(ctx context.Context, rawValue string, identityType constants.IdentityType, password value.Password) error {
	identity, err := aas.identityFactory.Create(rawValue, identityType)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "authAppService.Login: failed to create identity").WithError(err)
	}

	if err := aas.as.Authenticate(ctx, identity, password); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "authAppService.Login: failed to authenticate user").WithError(err)
	}

	return nil
}
