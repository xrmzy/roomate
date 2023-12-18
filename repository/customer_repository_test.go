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

type CustomerRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    CustomerRepository
}

func (suite *CustomerRepoTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.sqlmock = sqlmock
	suite.repo = NewCustomerRepository(suite.mockDB)
}

func TestCustomerRepoSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}

var dummyCustomer = entity.Customer{
	Id:          "23dfff3",
	Name:        "Invoker",
	Email:       "invoker@gmail.com",
	Address:     "Bandung",
	PhoneNumber: "082112288779",
	CreatedAt:   time.Now().Truncate(time.Second),
	UpdatedAt:   time.Now().Truncate(time.Second),
	IsDeleted:   false,
}

func (suite *CustomerRepoTestSuite) TestRepository_GetCustomer() {
	// test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetCustomer)).WithArgs(dummyCustomer.Id).WillReturnError(errors.New("Get by id failed"))

	_, err := suite.repo.Get(dummyCustomer.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Get by id failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "name", "email", "address", "phone_number", "created_at", "updated_at"}).AddRow(dummyCustomer.Id, dummyCustomer.Name, dummyCustomer.Email, dummyCustomer.Address, dummyCustomer.PhoneNumber, dummyCustomer.CreatedAt, dummyCustomer.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetCustomer)).WithArgs(dummyCustomer.Id).WillReturnRows(rows)

	actual, err := suite.repo.Get(dummyCustomer.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyCustomer.Id, actual.Id)
}

func (suite *CustomerRepoTestSuite) TestRepository_CreateCustomer() {
	//test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateCustomer)).WithArgs(dummyCustomer.Name, dummyCustomer.Email, dummyCustomer.Address, dummyCustomer.PhoneNumber, dummyCustomer.UpdatedAt).WillReturnError(errors.New("Create customer failed"))

	_, err := suite.repo.Create(dummyCustomer)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Create customer failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "name", "email", "address", "phone_number", "created_at", "updated_at"}).AddRow(dummyCustomer.Id, dummyCustomer.Name, dummyCustomer.Email, dummyCustomer.Address, dummyCustomer.PhoneNumber, dummyCustomer.CreatedAt, dummyCustomer.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateCustomer)).WithArgs(dummyCustomer.Name, dummyCustomer.Email, dummyCustomer.Address, dummyCustomer.PhoneNumber, dummyCustomer.UpdatedAt).WillReturnRows(rows)

	actual, err := suite.repo.Create(dummyCustomer)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyCustomer.Id, actual.Id)
}

func (suite *CustomerRepoTestSuite) TestRepository_GetAllCustomer() {

	// test failed

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllCustomers)).WithArgs(1, 0).WillReturnError(sql.ErrNoRows)

	_, err := suite.repo.GetAll(1, 0)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), sql.ErrNoRows, err)

	// test success
	rows := sqlmock.NewRows([]string{"id", "name", "email", "address", "phone_number", "created_at", "updated_at"}).AddRow(dummyCustomer.Id, dummyCustomer.Name, dummyCustomer.Email, dummyCustomer.Address, dummyCustomer.PhoneNumber, dummyCustomer.CreatedAt, dummyCustomer.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllCustomers)).WithArgs(1, 0).WillReturnRows(rows)

	actual, err := suite.repo.GetAll(1, 0)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyCustomer.Id, actual[0].Id)
	assert.Equal(suite.T(), dummyCustomer.Address, actual[0].Address)
	assert.Equal(suite.T(), dummyCustomer.PhoneNumber, actual[0].PhoneNumber)

}

func (suite *CustomerRepoTestSuite) TestRepository_UpdateCustomer() {

	updatedCustomer := dummyCustomer
	updatedCustomer.Name = "Udin"
	updatedCustomer.Email = "Udin@email.com"
	updatedCustomer.Address = "Jakarta"
	updatedCustomer.PhoneNumber = "082119944662"

	// mock error from QueryRow
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdateCustomer)).WithArgs(dummyCustomer.Id, updatedCustomer.Name, updatedCustomer.Email, updatedCustomer.Address, updatedCustomer.PhoneNumber, updatedCustomer.UpdatedAt).WillReturnError(errors.New("update failed"))

	_, err := suite.repo.Update(dummyCustomer.Id, updatedCustomer)

	// assertions
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update failed", err.Error())

	// test success
	rows := sqlmock.NewRows([]string{"id", "name", "email", "address", "phone_number", "created_at", "updated_at"}).AddRow(dummyCustomer.Id, updatedCustomer.Name, updatedCustomer.Email, updatedCustomer.Address, updatedCustomer.PhoneNumber, updatedCustomer.CreatedAt, updatedCustomer.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdateCustomer)).WithArgs(dummyCustomer.Id, updatedCustomer.Name, updatedCustomer.Email, updatedCustomer.Address, updatedCustomer.PhoneNumber, updatedCustomer.UpdatedAt).WillReturnRows(rows)

	actual, err := suite.repo.Update(dummyCustomer.Id, updatedCustomer)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyCustomer.Id, actual.Id)
	assert.Equal(suite.T(), updatedCustomer.Name, actual.Name)
	assert.Equal(suite.T(), updatedCustomer.Email, actual.Email)
	assert.Equal(suite.T(), updatedCustomer.Address, actual.Address)
	assert.Equal(suite.T(), updatedCustomer.PhoneNumber, actual.PhoneNumber)
}

func (suite *CustomerRepoTestSuite) TestRepository_DeleteCustomer() {
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteCustomer)).WithArgs(dummyCustomer.Id).
		WillReturnError(errors.New("delete failed"))

	err := suite.repo.Delete(dummyCustomer.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete failed", err.Error())

	// test success
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteCustomer)).WithArgs(dummyCustomer.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = suite.repo.Delete(dummyCustomer.Id)

	assert.NoError(suite.T(), err)
}
