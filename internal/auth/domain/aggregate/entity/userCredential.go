package entity

import (
	"time"

	"github.com/ant-02/myreel-plus/internal/auth/domain/aggregate/value"
	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
)

type UserCredential struct {
	id             value.UserID
	email          value.Email
	phone          value.Phone
	password       value.Password
	status         constants.UserStatus
	failedAttempts int64
	lockedUntil    *time.Time
	lastLoginAt    *time.Time
}

func (u *UserCredential) ID() value.UserID {
	return u.id
}

func (u *UserCredential) Email() value.Email {
	return u.email
}

func (u *UserCredential) Phone() value.Phone {
	return u.phone
}

func (u *UserCredential) Password() value.Password {
	return u.password
}

func (u *UserCredential) Status() constants.UserStatus {
	return u.status
}

func (u *UserCredential) FailedAttempts() int64 {
	return u.failedAttempts
}

// 安全的指针类型Getter，返回副本避免外部修改
func (u *UserCredential) LockedUntil() *time.Time {
	if u.lockedUntil == nil {
		return nil
	}
	t := *u.lockedUntil // 创建副本
	return &t
}

func (u *UserCredential) LastLoginAt() *time.Time {
	if u.lastLoginAt == nil {
		return nil
	}
	t := *u.lastLoginAt // 创建副本
	return &t
}

func (u *UserCredential) calculateLockDuration() time.Duration {
	multiplier := 1 << (u.failedAttempts - constants.MaxFailedAttempts)
	return constants.UserLockBaseDuration * time.Duration(multiplier)
}

func (u *UserCredential) RecordFailedAttempt() {
	u.failedAttempts++

	// 如果失败次数超过阈值，锁定账户
	if u.failedAttempts >= constants.MaxFailedAttempts {
		u.calculateLockDuration()
		lockDuration := u.calculateLockDuration()
		lockedUntil := time.Now().Add(lockDuration)
		u.lockedUntil = &lockedUntil
	}
}

func (u *UserCredential) ResetFailedAttempts() {
	u.failedAttempts = 0
	u.lockedUntil = nil
	now := time.Now()
	u.lastLoginAt = &now
}

func (u *UserCredential) isLocked() bool {
	if u.lockedUntil == nil {
		return false
	}
	return time.Now().Before(*u.lockedUntil)
}

func (u *UserCredential) VerifyAccountStatus() error {
	if u.status != constants.UserStatusActive {
		return errno.UserNotActive
	}

	if u.isLocked() {
		return errno.UserIsLocked
	}

	return nil
}

func (u *UserCredential) VerifyPassword(password value.Password) bool {
	// 检查账户是否被锁定
	if u.isLocked() {
		return false
	}

	// 比较密码
	return u.password.Compare(password.Value())
}

type UserCredentialBuilder struct {
	id             value.UserID
	email          value.Email
	phone          value.Phone
	password       value.Password
	status         constants.UserStatus
	failedAttempts int64
	lockedUntil    *time.Time
	lastLoginAt    *time.Time
}

func NewUserCredentialBuilder() *UserCredentialBuilder {
	return &UserCredentialBuilder{}
}

func (ucb *UserCredentialBuilder) WithID(id value.UserID) *UserCredentialBuilder {
	ucb.id = id
	return ucb
}

func (ucb *UserCredentialBuilder) WithEmail(email value.Email) *UserCredentialBuilder {
	ucb.email = email
	return ucb
}

func (ucb *UserCredentialBuilder) WithPhone(phone value.Phone) *UserCredentialBuilder {
	ucb.phone = phone
	return ucb
}

func (ucb *UserCredentialBuilder) WithPassword(password value.Password) *UserCredentialBuilder {
	ucb.password = password
	return ucb
}

func (ucb *UserCredentialBuilder) WithStatus(status constants.UserStatus) *UserCredentialBuilder {
	ucb.status = status
	return ucb
}

func (ucb *UserCredentialBuilder) WithFailedAttempts(failedAttempts int64) *UserCredentialBuilder {
	ucb.failedAttempts = failedAttempts
	return ucb
}

func (ucb *UserCredentialBuilder) WithLockedUntil(lockedUntil *time.Time) *UserCredentialBuilder {
	ucb.lockedUntil = lockedUntil
	return ucb
}

func (ucb *UserCredentialBuilder) WithLastLoginAt(lastLoginAt *time.Time) *UserCredentialBuilder {
	ucb.lastLoginAt = lastLoginAt
	return ucb
}

func (ucb *UserCredentialBuilder) Build() (*UserCredential, error) {
	return &UserCredential{
		id:             ucb.id,
		email:          ucb.email,
		phone:          ucb.phone,
		password:       ucb.password,
		status:         ucb.status,
		failedAttempts: ucb.failedAttempts,
		lastLoginAt:    ucb.lastLoginAt,
		lockedUntil:    ucb.lockedUntil,
	}, nil
}
