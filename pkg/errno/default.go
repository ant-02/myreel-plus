package errno

var (
	Success              = NewErrNo(SuccessCode, "ok")
	InternalServiceError = NewErrNo(InternalServiceErrorCode, "internal server error")
)

// auth
var (
	ErrUserCredentialNotFound = NewErrNo(InternalServiceErrorCode, "user credential not found")
	EmailFormatError          = NewErrNo(InternalServiceErrorCode, "email format is error")
	PhoneFormatError          = NewErrNo(InternalServiceErrorCode, "phone format is error")
	PasswordFormatError       = NewErrNo(InternalServiceErrorCode, "password format is error")
	IdentityTypeError         = NewErrNo(InternalServiceErrorCode, "identity type is error")
	UserNotActive             = NewErrNo(InternalServiceErrorCode, "user not active")
	UserIsLocked              = NewErrNo(InternalServiceErrorCode, "user is locked")
)
