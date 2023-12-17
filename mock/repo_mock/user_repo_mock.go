// user_repository_mock.go

package repomock

// import (
// 	"roomate/model/entity"

// 	"github.com/stretchr/testify/mock"
// )

// // MockUserRepository adalah struct mock untuk UserRepository
// type MockUserRepository struct {
// 	mock.Mock
// }

// // Get adalah fungsi untuk meniru Get pada UserRepository
// func (m *MockUserRepository) Get(id string) (entity.User, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(entity.User), args.Error(1)
// }

// // GetAll adalah fungsi untuk meniru GetAll pada UserRepository
// func (m *MockUserRepository) GetAll(limit, offset int) ([]entity.User, error) {
// 	args := m.Called(limit, offset)
// 	return args.Get(0).([]entity.User), args.Error(1)
// }

// // GetByEmail adalah fungsi untuk meniru GetByEmail pada UserRepository
// func (m *MockUserRepository) GetByEmail(email string) (entity.User, error) {
// 	args := m.Called(email)
// 	return args.Get(0).(entity.User), args.Error(1)
// }

// // Create adalah fungsi untuk meniru Create pada UserRepository
// func (m *MockUserRepository) Create(user entity.User) (entity.User, error) {
// 	args := m.Called(user)
// 	return args.Get(0).(entity.User), args.Error(1)
// }

// // Update adalah fungsi untuk meniru Update pada UserRepository
// func (m *MockUserRepository) Update(id string, user entity.User) (entity.User, error) {
// 	args := m.Called(id, user)
// 	return args.Get(0).(entity.User), args.Error(1)
// }

// // UpdatePassword adalah fungsi untuk meniru UpdatePassword pada UserRepository
// func (m *MockUserRepository) UpdatePassword(id, password string) (entity.User, error) {
// 	args := m.Called(id, password)
// 	return args.Get(0).(entity.User), args.Error(1)
// }

// // Delete adalah fungsi untuk meniru Delete pada UserRepository
// func (m *MockUserRepository) Delete(id string) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }
