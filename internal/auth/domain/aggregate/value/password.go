package value

import (
	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
	"golang.org/x/crypto/bcrypt"
)

// Password 实体
type Password struct {
	value   string
	compare func(hash, password string) bool
}

func (p *Password) Value() string {
	return p.value
}

func (p *Password) Compare(password string) bool {
	return p.compare(p.value, password)
}

// Passowrd 验证器
type PasswordValidator interface {
	Validate(val string) bool
}

type DefaultPasswordValidator struct{}

func (dpv *DefaultPasswordValidator) Validate(val string) bool {
	return len(val) >= constants.UserMinimumPasswordLength && len(val) <= constants.UserMaximumPasswordLength
}

// Password 加密器
type PasswordEncryptor interface {
	Encrypt(val string) (string, error)
	Compare(hash, password string) bool
}

type DefaultPasswordEncryptor struct{}

func (dpe *DefaultPasswordEncryptor) Encrypt(val string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(val), constants.UserDefaultEncryptPasswordCost)
	return string(hash), err
}

func (dpe *DefaultPasswordEncryptor) Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type PasswordBuilder struct {
	value       string
	validator   PasswordValidator
	encryptor   PasswordEncryptor
	isValidated bool
	encrypt     bool
}

func NewPasswordBuilder(val string) *PasswordBuilder {
	return &PasswordBuilder{
		value:       val,
		validator:   &DefaultPasswordValidator{},
		encryptor:   &DefaultPasswordEncryptor{},
		isValidated: false,
		encrypt:     false,
	}
}

func (pb *PasswordBuilder) WithValidator(validator PasswordValidator) *PasswordBuilder {
	pb.validator = validator
	return pb
}

func (pb *PasswordBuilder) WithEncryptor(encryptor PasswordEncryptor) *PasswordBuilder {
	pb.encryptor = encryptor
	return pb
}

func (pb *PasswordBuilder) WithIsValidated() *PasswordBuilder {
	pb.isValidated = true
	return pb
}

func (pb *PasswordBuilder) WithEncrypt() *PasswordBuilder {
	pb.encrypt = true
	return pb
}

func (pb *PasswordBuilder) Build() (*Password, error) {
	if !pb.isValidated && !pb.validator.Validate(pb.value) {
		return nil, errno.PasswordFormatError
	}

	if pb.encrypt {
		hash, err := pb.encryptor.Encrypt(pb.value)
		if err != nil {
			return nil, errno.Errorf(errno.InternalServiceErrorCode, "password.NewPassword: failed to encode password, err: %v", err)
		}
		return &Password{
			value:   hash,
			compare: pb.encryptor.Compare,
		}, nil
	}

	return &Password{
		value:   pb.value,
		compare: pb.encryptor.Compare,
	}, nil
}
