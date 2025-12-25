package mapper

import (
	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/entity"
	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/value"
	"github.com/ant-02/myreel-plus/internal/auth/infrastructure/persistence/model"
	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
)

type userCredentialMapper struct{}

func NewUserCredentialMapper() *userCredentialMapper {
	return &userCredentialMapper{}
}

func (ucm *userCredentialMapper) ToDomainEntity(uc *model.UserCredential) (*entity.UserCredential, error) {
	idBuilder := value.NewUserIDBuilder().WithValue(uc.ID)
	id, err := idBuilder.Build()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "userCredentialMapper.ToDomainEntity: failed to build user id").WithError(err)
	}

	emailBuilder := value.NewEmailBuilder(uc.Email).WithIsValidated()
	email, err := emailBuilder.Build()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "userCredentialMapper.ToDomainEntity: failed to build email").WithError(err)
	}

	phoneBuilder := value.NewPhoneBuilder(uc.Phone).WithIsValidated()
	phone, err := phoneBuilder.Build()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "userCredentialMapper.ToDomainEntity: failed to build phone").WithError(err)
	}

	passwordBuilder := value.NewPasswordBuilder(uc.PasswordHash).WithIsValidated().WithEncrypt()
	password, err := passwordBuilder.Build()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "userCredentialMapper.ToDomainEntity: failed to build password").WithError(err)
	}

	return entity.NewUserCredentialBuilder().
		WithID(*id).
		WithEmail(*email).
		WithPhone(*phone).
		WithPassword(*password).
		WithStatus(constants.UserStatus(uc.Status)).
		WithFailedAttempts(uc.FailedAttempts).
		WithLockedUntil(uc.LockedUntil).
		WithLastLoginAt(uc.LastLoginAt).
		Build()
}

func (ucm *userCredentialMapper) ToPersistenceModel(uc *entity.UserCredential) (*model.UserCredential, error) {
	if uc == nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "userCredentialMapper.ToPersistenceModel: entity is nil")
	}

	id := uc.ID()
	email := uc.Email()
	phone := uc.Phone()
	password := uc.Password()

	return &model.UserCredential{
		ID:             id.Value(),
		Email:          email.Value(),
		Phone:          phone.Value(),
		PasswordHash:   password.Value(),
		Status:         int64(uc.Status()),
		FailedAttempts: uc.FailedAttempts(),
		LockedUntil:    uc.LockedUntil(),
		LastLoginAt:    uc.LastLoginAt(),
	}, nil
}
