package repository

import (
	"context"

	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/entity"
	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/value"
)

type UserCredentialRepository interface {
	Magrate() error
	FindByIdentity(ctx context.Context, identity value.Identity) (*entity.UserCredential, error)
	Save(ctx context.Context, user *entity.UserCredential) error
}
