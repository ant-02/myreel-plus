package value

import (
	"github.com/ant-02/myreel-plus/pkg/constants"
	"github.com/ant-02/myreel-plus/pkg/errno"
)

type Identity interface {
	Type() constants.IdentityType
	Value() string
}

type IdentityFactory struct{}

func (f *IdentityFactory) Create(
	rawValue string,
	identityType constants.IdentityType,
) (Identity, error) {
	switch identityType {
	case constants.IdentityTypeEmail:
		builder := NewEmailBuilder(rawValue)
		return builder.Build()
	case constants.IdentityTypePhone:
		builder := NewPhoneBuilder(rawValue)
		return builder.Build()
	default:
		return nil, errno.IdentityTypeError
	}
}
