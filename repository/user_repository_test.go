package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"roomate/model/entity"
	"roomate/utils/common"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    UserRepository
}

func (suite *UserRepoTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.sqlmock = sqlmock
	suite.repo = NewUserRepository(suite.mockDB)
}

func TestUserRepoSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

var dummyUser = entity.User{
	Id:        "2121",
	Name:      "Invoker",
	Email:     "invoker@gmail.com",
	Password:  "secret123",
	RoleId:    "1",
	RoleName:  "admin",
	CreatedAt: time.Now().Truncate(time.Second),
	UpdatedAt: time.Now().Truncate(time.Second),
	IsDeleted: false,
}

func (suite *UserRepoTestSuite) TestRepository_CreateUser() {

	rows := sqlmock.NewRows([]string{"id", "name", "email", "role_id", "role_name", "created_at", "updated_at"}).AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.RoleId, dummyUser.RoleName, dummyUser.CreatedAt, dummyUser.UpdatedAt)

	// expected return success
	suite.sqlmock.ExpectQuery("INSERT INTO users").WithArgs(dummyUser.Name, dummyUser.Email, dummyUser.Password, dummyUser.RoleId, dummyUser.RoleName, dummyUser.UpdatedAt).WillReturnRows(rows)

	actual, err := suite.repo.Create(dummyUser)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyUser.Id, actual.Id)
	assert.Equal(suite.T(), dummyUser.RoleId, actual.RoleId)
	assert.Equal(suite.T(), dummyUser.RoleName, actual.RoleName)
}

func (suite *UserRepoTestSuite) TestRepository_CreateUserFail() {

	// expected return error
	suite.sqlmock.ExpectQuery("INSERT INTO users").WithArgs(dummyUser.Name, dummyUser.Email, dummyUser.Password, dummyUser.RoleId, dummyUser.RoleName, dummyUser.UpdatedAt).WillReturnError(errors.New("insert failed"))

	_, err := suite.repo.Create(dummyUser)
	assert.Error(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestRepository_GetUser() {

	suite.sqlmock.ExpectQuery("SELECT id, name, email, role_id, role_name, created_at, updated_at FROM users").WithArgs(dummyUser.Id).WillReturnError(errors.New("select failed"))

	_, err := suite.repo.Get(dummyUser.Id)
	assert.Error(suite.T(), err)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "role_id", "role_name", "created_at", "updated_at"}).AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.RoleId, dummyUser.RoleName, dummyUser.CreatedAt, dummyUser.UpdatedAt)

	suite.sqlmock.ExpectQuery("SELECT id, name, email, role_id, role_name, created_at, updated_at FROM users").WithArgs(dummyUser.Id).WillReturnRows(rows)

	actual, err := suite.repo.Get(dummyUser.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyUser.Id, actual.Id)
}

func (suite *UserRepoTestSuite) TestRepository_GetAllUser() {

	// test failed

	suite.sqlmock.ExpectQuery("SELECT id, name, email, role_id, role_name, created_at, updated_at FROM users WHERE is_deleted = false").WithArgs(1, 0).WillReturnError(sql.ErrNoRows)

	_, err := suite.repo.GetAll(1, 0)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), sql.ErrNoRows, err)

	// test success
	rows := sqlmock.NewRows([]string{"id", "name", "email", "role_id", "role_name", "created_at", "updated_at"}).AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.RoleId, dummyUser.RoleName, dummyUser.CreatedAt, dummyUser.UpdatedAt)

	suite.sqlmock.ExpectQuery("SELECT id, name, email, role_id, role_name, created_at, updated_at FROM users WHERE is_deleted = false ").WithArgs(1, 0).WillReturnRows(rows)

	actual, err := suite.repo.GetAll(1, 0)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyUser.Id, actual[0].Id)
	assert.Equal(suite.T(), dummyUser.RoleId, actual[0].RoleId)
	assert.Equal(suite.T(), dummyUser.RoleName, actual[0].RoleName)
	// assert.Empty(suite.T(), users) // Verify empty user slice after error
}

func (suite *UserRepoTestSuite) TestRepository_UpdateUser() {

	updatedUser := dummyUser
	updatedUser.Name = "Udin"
	updatedUser.Email = "Udin@email.com"
	updatedUser.RoleId = "2"
	updatedUser.RoleName = "employee"

	// mock error from QueryRow
	suite.sqlmock.ExpectQuery("UPDATE users").
		WithArgs(dummyUser.Id, updatedUser.Name, updatedUser.Email, updatedUser.RoleId, updatedUser.RoleName, updatedUser.UpdatedAt).
		WillReturnError(fmt.Errorf("update failed"))

	_, err := suite.repo.Update(dummyUser.Id, updatedUser)

	// assertions
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update failed", err.Error())

	// mock expected output user with updated values

	rows := sqlmock.NewRows([]string{"id", "name", "email", "role_id", "role_name", "created_at", "updated_at"}).AddRow(updatedUser.Id, updatedUser.Name, updatedUser.Email, updatedUser.RoleId, updatedUser.RoleName, updatedUser.CreatedAt, updatedUser.UpdatedAt)

	suite.sqlmock.ExpectQuery("UPDATE users").WithArgs(dummyUser.Id, updatedUser.Name, updatedUser.Email, updatedUser.RoleId, updatedUser.RoleName, dummyUser.UpdatedAt).WillReturnRows(rows)

	actual, err := suite.repo.Update(dummyUser.Id, updatedUser)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyUser.Id, actual.Id)
	assert.Equal(suite.T(), updatedUser.Name, actual.Name)
	assert.Equal(suite.T(), updatedUser.Email, actual.Email)
	assert.Equal(suite.T(), updatedUser.RoleId, actual.RoleId)
	assert.Equal(suite.T(), updatedUser.RoleName, actual.RoleName)
}

func (suite *UserRepoTestSuite) TestRepository_DeleteUser() {
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteUser)).WithArgs(dummyUser.Id).
		WillReturnError(errors.New("delete failed"))

	err := suite.repo.Delete(dummyUser.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete failed", err.Error())

	// test success
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteUser)).WithArgs(dummyUser.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = suite.repo.Delete(dummyUser.Id)

	assert.NoError(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestRepository_GetByEmail() {
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetByEmail)).WithArgs(dummyUser.Email).WillReturnError(sql.ErrNoRows)

	_, err := suite.repo.GetByEmail(dummyUser.Email)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), sql.ErrNoRows, err)

	rows := sqlmock.NewRows([]string{"id", "role_name", "password"}).AddRow(dummyUser.Id, dummyUser.RoleName, dummyUser.Password)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetByEmail)).WithArgs(dummyUser.Email).WillReturnRows(rows)

	user, err := suite.repo.GetByEmail(dummyUser.Email)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyUser.Id, user.Id)

}

func (suite *UserRepoTestSuite) TestRepository_UpdatePassword() {
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdatePassword)).WithArgs(dummyUser.Id, "kieukieu123").WillReturnError(errors.New("error change password: Update failed"))

	_, err := suite.repo.UpdatePassword(dummyUser.Id, "kieukieu123")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error change password: Update failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "name", "email", "role_id", "role_name", "created_at", "updated_at"}).AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.RoleId, dummyUser.RoleName, dummyUser.CreatedAt, dummyUser.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdatePassword)).WithArgs(dummyUser.Id, "kieukieu123").WillReturnRows(rows)

	user, err := suite.repo.UpdatePassword(dummyUser.Id, "kieukieu123")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyUser.Id, user.Id)

}
