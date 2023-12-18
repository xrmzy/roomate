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

type ServiceRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    ServiceRepository
}

func (suite *ServiceRepoTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.sqlmock = sqlmock
	suite.repo = NewServiceRepository(suite.mockDB)
}

func TestServiceRepoSuite(t *testing.T) {
	suite.Run(t, new(ServiceRepoTestSuite))
}

var dummyService = entity.Service{
	Id:        "2",
	Name:      "Breakfast",
	Price:     80000,
	CreatedAt: time.Now().Truncate(time.Second),
	UpdatedAt: time.Now().Truncate(time.Second),
	IsDeleted: false,
}

func (suite *ServiceRepoTestSuite) TestRepository_GetService() {
	// test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetService)).WithArgs(dummyService.Id).WillReturnError(errors.New("Get by id failed"))

	_, err := suite.repo.Get(dummyService.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Get by id failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "name", "price", "created_at", "updated_at"}).AddRow(dummyService.Id, dummyService.Name, dummyService.Price, dummyService.CreatedAt, dummyService.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetService)).WithArgs(dummyService.Id).WillReturnRows(rows)

	actual, err := suite.repo.Get(dummyService.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyService.Id, actual.Id)
}

func (suite *ServiceRepoTestSuite) TestRepository_GetAllServices() {
	// test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllServices)).WithArgs(1, 0).WillReturnError(errors.New("Get All failed"))

	_, err := suite.repo.GetAll(1, 0)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Get All failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "name", "price", "created_at", "updated_at"}).AddRow(dummyService.Id, dummyService.Name, dummyService.Price, dummyService.CreatedAt, dummyService.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllServices)).WithArgs(1, 0).WillReturnRows(rows)

	actual, err := suite.repo.GetAll(1, 0)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyService.Id, actual[0].Id)
}

func (suite *ServiceRepoTestSuite) TestRepository_CreateService() {
	//test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateService)).WithArgs(dummyService.Id, dummyService.Name, dummyService.Price, dummyService.UpdatedAt).WillReturnError(errors.New("Create service failed"))

	_, err := suite.repo.Create(dummyService)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Create service failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "name", "price", "created_at", "updated_at"}).AddRow(dummyService.Id, dummyService.Name, dummyService.Price, dummyService.CreatedAt, dummyService.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateService)).WithArgs(dummyService.Id, dummyService.Name, dummyService.Price, dummyService.UpdatedAt).WillReturnRows(rows)
	actual, err := suite.repo.Create(dummyService)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyService.Id, actual.Id)
}

func (suite *ServiceRepoTestSuite) TestRepository_UpdateService() {

	updatedService := dummyService
	updatedService.Name = "Extra beverages"
	updatedService.Price = 50000

	// mock error from QueryRow
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdateService)).WithArgs(dummyService.Id, updatedService.Name, updatedService.Price, updatedService.UpdatedAt).WillReturnError(errors.New("update failed"))

	_, err := suite.repo.Update(dummyService.Id, updatedService)

	// assertions
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update failed", err.Error())

	// test success
	rows := sqlmock.NewRows([]string{"id", "name", "price", "created_at", "updated_at"}).AddRow(dummyService.Id, updatedService.Name, updatedService.Price, updatedService.CreatedAt, updatedService.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdateService)).WithArgs(dummyService.Id, updatedService.Name, updatedService.Price, updatedService.UpdatedAt).WillReturnRows(rows)

	actual, err := suite.repo.Update(dummyService.Id, updatedService)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyService.Id, actual.Id)
	assert.Equal(suite.T(), updatedService.Name, actual.Name)
}

func (suite *ServiceRepoTestSuite) TestRepository_DeleteService() {
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteService)).WithArgs(dummyService.Id).
		WillReturnError(errors.New("delete failed"))

	err := suite.repo.Delete(dummyService.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete failed", err.Error())

	// test success
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteService)).WithArgs(dummyService.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = suite.repo.Delete(dummyService.Id)

	assert.NoError(suite.T(), err)
}
