package value

type UserID struct {
	value int64
}

func (uid *UserID) Value() int64 {
	return uid.value
}

type UserIDBuilder struct {
	value int64
	// isGenerated bool
}

func NewUserIDBuilder() *UserIDBuilder {
	return &UserIDBuilder{}
}

func (uib *UserIDBuilder) WithValue(val int64) *UserIDBuilder {
	uib.value = val
	return uib
}

func (uib *UserIDBuilder) Build() (*UserID, error) {
	return &UserID{value: uib.value}, nil
}
