package service

import (
	"context"

	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/value"
	"github.com/ant-02/myreel-plus/internal/auth/domain/repository"
	"github.com/ant-02/myreel-plus/pkg/errno"
)

type authenticationService struct {
	userCredentialRepo repository.UserCredentialRepository
}

type AuthenticationService interface {
	Authenticate(ctx context.Context, identity value.Identity, password value.Password) error
}

func NewAuthenticationService(userCredentialRepo repository.UserCredentialRepository) AuthenticationService {
	return &authenticationService{userCredentialRepo: userCredentialRepo}
}

func (s *authenticationService) Authenticate(
	ctx context.Context,
	identity value.Identity,
	password value.Password,
) error {

	// 1. 查找用户
	user, err := s.userCredentialRepo.FindByIdentity(ctx, identity)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "authentication.Authenticate: failed to get user by identity").WithError(err)
	}

	// 2. 检查账户状态
	if err := user.VerifyAccountStatus(); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "authentication.Authenticate: user is unactive").WithError(err)
	}

	// 3. 验证密码（调用领域对象的方法）
	if !user.VerifyPassword(password) {
		user.RecordFailedAttempt()
		return errno.NewErrNo(errno.InternalServiceErrorCode, "authentication.Authenticate: password is error")
	}

	// 4. 重置失败计数
	user.ResetFailedAttempts()
	if err := s.userCredentialRepo.Save(ctx, user); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "authentication.Authenticate: failed to save user").WithError(err)
	}

	return nil
}
