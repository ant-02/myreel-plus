package mysql

import (
	"context"
	"errors"

	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/entity"
	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/value"
	"github.com/ant-02/myreel-plus/internal/auth/domain/repository"
	"github.com/ant-02/myreel-plus/internal/auth/infrastructure/persistence/mapper"
	"github.com/ant-02/myreel-plus/internal/auth/infrastructure/persistence/model"
	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
	"gorm.io/gorm"
)

type userDB struct {
	client *gorm.DB
}

func (db *userDB) Magrate() error {
	if err := db.client.AutoMigrate(&model.UserCredential{}); err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to auto magrate user credential model, err: %v", err)
	}
	return nil
}

func NewUserDB(c *gorm.DB) repository.UserCredentialRepository {
	return &userDB{client: c}
}

func (ud *userDB) FindByIdentity(ctx context.Context, identity value.Identity) (*entity.UserCredential, error) {
	tx := ud.client.WithContext(ctx).Model(&model.UserCredential{})
	switch identity.Type() {
	case constants.IdentityTypeEmail:
		tx = tx.Where("email = ?", identity.Value())
	case constants.IdentityTypePhone:
		tx = tx.Where("phone = ?", identity.Value())
	default:
		return nil, errno.IdentityTypeError
	}
	var model model.UserCredential
	if err := tx.Find(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserCredentialNotFound
		}
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "userDB.FindByIdentity: failed to get user credential, err: %v", err)
	}

	um := mapper.NewUserCredentialMapper()
	result, err := um.ToDomainEntity(&model)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "userDB.FindByIdentity: failed to convert user credential").WithError(err)
	}

	return result, nil
}

func (ud *userDB) Save(ctx context.Context, user *entity.UserCredential) error {
	um := mapper.NewUserCredentialMapper()
	result, err := um.ToPersistenceModel(user)
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "userDB.Save: failed to convert user credential").WithError(err)
	}

	if err := ud.client.WithContext(ctx).Save(&result).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "userDB.Save: failed to save user credential, err: %v", err)
	}

	return nil
}
