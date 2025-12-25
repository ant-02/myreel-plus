package constants

import "time"

// service name
const (
	GatewayServiceName = "gateway"
	AuthServiceName    = "auth"
)

// auth
type IdentityType int64

const (
	IdentityTypeEmail IdentityType = iota
	IdentityTypePhone
)

type UserStatus int64

const (
	UserStatusUnActive UserStatus = iota
	UserStatusActive
)

const (
	UserMaximumPasswordLength      = 72 // DO NOT EDIT (ref: bcrypt.GenerateFromPassword)
	UserMinimumPasswordLength      = 5
	UserDefaultEncryptPasswordCost = 10
	MaxFailedAttempts              = 10
	UserLockBaseDuration           = 5 * time.Minute
)
