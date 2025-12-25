package value

import (
	"regexp"

	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
)

// Phone 实体
type Phone struct {
	value        string
	validator    PhoneValidator
	identityType constants.IdentityType
}

func (p *Phone) Type() constants.IdentityType {
	return p.identityType
}

func (p *Phone) Value() string {
	return p.value
}

// Phone 验证器
type PhoneValidator interface {
	Validate(val string) bool
}

type DefaultPhoneValidator struct{}

func (dpv *DefaultPhoneValidator) Validate(val string) bool {
	if len(val) != 11 {
		return false
	}

	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(val)
}

// Phone Builder 构建器
type PhoneBuilder struct {
	value        string
	validator    PhoneValidator
	identityType constants.IdentityType
	isValidated  bool
}

func NewPhoneBuilder(val string) *PhoneBuilder {
	return &PhoneBuilder{
		value:        val,
		validator:    &DefaultPhoneValidator{},
		identityType: constants.IdentityTypePhone,
		isValidated:  false,
	}
}

func (pb *PhoneBuilder) WithValidator(validator PhoneValidator) *PhoneBuilder {
	pb.validator = validator
	return pb
}

func (pb *PhoneBuilder) WithIsValidated() *PhoneBuilder {
	pb.isValidated = true
	return pb
}

func (pb *PhoneBuilder) Build() (*Phone, error) {
	if !pb.isValidated && !pb.validator.Validate(pb.value) {
		return nil, errno.PhoneFormatError
	}
	
	return &Phone{
		value:        pb.value,
		identityType: pb.identityType,
	}, nil
}
