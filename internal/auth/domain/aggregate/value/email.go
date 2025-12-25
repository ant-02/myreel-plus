package value

import (
	"net/mail"

	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
)

// Email 实体
type Email struct {
	value        string
	identityType constants.IdentityType
}

func (e *Email) Type() constants.IdentityType {
	return e.identityType
}

func (e *Email) Value() string {
	return e.value
}

// Email 格式验证器
type EmailValidator interface {
	Validate(val string) bool
}

type DefaultEmailValidator struct{}

func (dev *DefaultEmailValidator) Validate(val string) bool {
	_, err := mail.ParseAddress(val)
	return err == nil
}

// Email Builder 构建器
type Emailbuilder struct {
	value        string
	validator    EmailValidator
	identityType constants.IdentityType
	isValidated  bool
}

func NewEmailBuilder(val string) *Emailbuilder {
	return &Emailbuilder{
		value:        val,
		validator:    &DefaultEmailValidator{},
		identityType: constants.IdentityTypeEmail,
		isValidated:  false,
	}
}

func (eb *Emailbuilder) WithValidator(validator EmailValidator) *Emailbuilder {
	eb.validator = validator
	return eb
}

func (eb *Emailbuilder) WithIsValidated() *Emailbuilder {
	eb.isValidated = true
	return eb
}

func (eb *Emailbuilder) Build() (*Email, error) {
	if !eb.isValidated && !eb.validator.Validate(eb.value) {
		return nil, errno.EmailFormatError
	}

	return &Email{
		value:        eb.value,
		identityType: eb.identityType,
	}, nil
}
