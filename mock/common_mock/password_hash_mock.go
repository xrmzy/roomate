package commonmock

import "github.com/stretchr/testify/mock"

type PasswordHashCommonMock struct {
	mock.Mock
}

func (p *PasswordHashCommonMock) GeneratePasswordHash(password string) (string, error) {
	args := p.Called(password)
	return args.Get(0).(string), args.Error(1)
}

func (p *PasswordHashCommonMock) ComparePasswordHash(hashedPassword, password string) error {
	args := p.Called(hashedPassword, password)
	return args.Error(0)
}
