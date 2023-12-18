package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"roomate/model/entity"
	"roomate/utils/common"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoleRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    RoleRepository
}

func (suite *RoleRepoTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.sqlmock = sqlmock
	suite.repo = NewRoleRepository(suite.mockDB)
}

func TestRoleRepoSuite(t *testing.T) {
	suite.Run(t, new(RoleRepoTestSuite))
}

var dummyRole = entity.Role{
	Id:        "2",
	RoleName:  "admin",
	CreatedAt: time.Now().Truncate(time.Second),
	UpdatedAt: time.Now().Truncate(time.Second),
	IsDeleted: false,
}

func (suite *RoleRepoTestSuite) TestRepository_GetRole() {
	// test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetRole)).WithArgs(dummyRole.Id).WillReturnError(errors.New("Get by id failed"))

	_, err := suite.repo.Get(dummyRole.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Get by id failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "role_name", "created_at", "updated_at"}).AddRow(dummyRole.Id, dummyRole.RoleName, dummyRole.CreatedAt, dummyRole.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetRole)).WithArgs(dummyRole.Id).WillReturnRows(rows)

	actual, err := suite.repo.Get(dummyRole.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyRole.Id, actual.Id)
}

func (suite *RoleRepoTestSuite) TestRepository_GetAllRoles() {
	// test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllRoles)).WithArgs(1, 0).WillReturnError(errors.New("Get by id failed"))

	_, err := suite.repo.GetAll(1, 0)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Get by id failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "role_name", "created_at", "updated_at"}).AddRow(dummyRole.Id, dummyRole.RoleName, dummyRole.CreatedAt, dummyRole.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllRoles)).WithArgs(1, 0).WillReturnRows(rows)

	actual, err := suite.repo.GetAll(1, 0)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyRole.Id, actual[0].Id)
}

func (suite *RoleRepoTestSuite) TestRepository_CreateRole() {
	//test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateRole)).WithArgs(dummyRole.RoleName, dummyRole.UpdatedAt).WillReturnError(errors.New("Create role failed"))

	_, err := suite.repo.Create(dummyRole)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Create role failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "role_name", "created_at", "updated_at"}).AddRow(dummyRole.Id, dummyRole.RoleName, dummyRole.CreatedAt, dummyRole.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateRole)).WithArgs(dummyRole.RoleName, dummyRole.UpdatedAt).WillReturnRows(rows)
	actual, err := suite.repo.Create(dummyRole)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyRole.Id, actual.Id)
}

func (suite *RoleRepoTestSuite) TestRepository_UpdateRole() {

	updatedRole := dummyRole
	updatedRole.RoleName = "employee"

	// mock error from QueryRow
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdateRole)).WithArgs(dummyRole.Id, updatedRole.RoleName, updatedRole.UpdatedAt).WillReturnError(errors.New("update failed"))

	_, err := suite.repo.Update(dummyRole.Id, updatedRole)

	// assertions
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update failed", err.Error())

	// test success
	rows := sqlmock.NewRows([]string{"id", "role_name", "created_at", "updated_at"}).AddRow(dummyRole.Id, updatedRole.RoleName, updatedRole.CreatedAt, updatedRole.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdateRole)).WithArgs(dummyRole.Id, updatedRole.RoleName, updatedRole.UpdatedAt).WillReturnRows(rows)

	actual, err := suite.repo.Update(dummyRole.Id, updatedRole)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyRole.Id, actual.Id)
	assert.Equal(suite.T(), updatedRole.RoleName, actual.RoleName)
}

func (suite *RoleRepoTestSuite) TestRepository_DeleteRole() {
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteRole)).WithArgs(dummyRole.Id).
		WillReturnError(errors.New("delete failed"))

	err := suite.repo.Delete(dummyRole.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete failed", err.Error())

	// test success
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteRole)).WithArgs(dummyRole.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = suite.repo.Delete(dummyRole.Id)

	assert.NoError(suite.T(), err)
}
