package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type DepositSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	or       *mock.MockOperationRepository
	wr       *mock.MockWalletRepository
	Database *database.Database
	Deposit  *Deposit
}

func TestDepositSuite(t *testing.T) {
	suite.Run(t, new(DepositSuite))
}

func (d *DepositSuite) SetupTest() {
	d.ctrl = gomock.NewController(d.T())
	d.or = mock.NewMockOperationRepository(d.ctrl)
	d.wr = mock.NewMockWalletRepository(d.ctrl)
	d.Database = database.NewDatabase(&gorm.DB{}, d.wr, d.or)
	d.Deposit = NewDeposit(d.Database)
}

func (d *DepositSuite) TearDownTest() {
	defer d.ctrl.Finish()
}

func (d *DepositSuite) TestShortDoDeposit() {
	username := "user_name_test"
	coin := "btc"
	value := float64(1.23)

	d.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{}, nil)
	d.wr.EXPECT().Save(gomock.Any()).Return(nil)
	d.or.EXPECT().Insert(gomock.Any()).Return(nil)

	err := d.Deposit.Execute(username, coin, value)

	assert.Empty(d.T(), err)
}

func (d *DepositSuite) TestShortReturnErrorInFindWallet() {
	username := "user_name_test"
	coin := "btc"
	value := float64(1.23)
	errExpect := errors.New("Error expected")

	d.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{}, errExpect)

	err := d.Deposit.Execute(username, coin, value)

	assert.NotEmpty(d.T(), err)
	assert.EqualError(d.T(), err, err.Error())
}

func (d *DepositSuite) TestShortReturnErrorNewOperatio() {
	username := "user_name_test"
	coin := "coin_not_exist"
	value := float64(1.23)
	errExpect := errors.New("invalid currency type")

	err := d.Deposit.Execute(username, coin, value)

	assert.NotEmpty(d.T(), err)
	assert.EqualError(d.T(), err, errExpect.Error())
}

func (d *DepositSuite) TestShortReturnErrorInSaveWallet() {
	username := "user_name_test"
	coin := "ada"
	value := float64(1.23)
	errExpect := errors.New("Error expected")

	d.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{}, nil)
	d.wr.EXPECT().Save(gomock.Any()).Return(errExpect)

	err := d.Deposit.Execute(username, coin, value)

	assert.NotEmpty(d.T(), err)
	assert.EqualError(d.T(), err, errExpect.Error())
}

func (d *DepositSuite) TestShortReturnErrorInInsertOperation() {
	username := "user_name_test"
	coin := "eth"
	value := float64(1.23)
	errExpect := errors.New("Error expected")

	d.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{}, nil)
	d.wr.EXPECT().Save(gomock.Any()).Return(nil)
	d.or.EXPECT().Insert(gomock.Any()).Return(errExpect)

	err := d.Deposit.Execute(username, coin, value)

	assert.NotEmpty(d.T(), err)
	assert.EqualError(d.T(), err, errExpect.Error())
}

func (d *DepositSuite) TestShortDoDepositWithXrp() {
	username := "user_name_test"
	coin := "xrp"
	value := float64(1.23)

	d.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{}, nil)
	d.wr.EXPECT().Save(gomock.Any()).Return(nil)
	d.or.EXPECT().Insert(gomock.Any()).Return(nil)

	err := d.Deposit.Execute(username, coin, value)

	assert.Empty(d.T(), err)
}

func (d *DepositSuite) TestShortDoDepositWithDoge() {
	username := "user_name_test"
	coin := "doge"
	value := float64(1.23)

	d.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{}, nil)
	d.wr.EXPECT().Save(gomock.Any()).Return(nil)
	d.or.EXPECT().Insert(gomock.Any()).Return(nil)

	err := d.Deposit.Execute(username, coin, value)

	assert.Empty(d.T(), err)
}