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

type WithdrawSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	or       *mock.MockOperationRepository
	wr       *mock.MockWalletRepository
	Database *database.Database
	Withdraw *Withdraw
}

func TestWithdrawSuite(t *testing.T) {
	suite.Run(t, new(WithdrawSuite))
}

func (w *WithdrawSuite) SetupTest() {
	w.ctrl = gomock.NewController(w.T())
	w.or = mock.NewMockOperationRepository(w.ctrl)
	w.wr = mock.NewMockWalletRepository(w.ctrl)
	w.Database = database.NewDatabase(&gorm.DB{}, w.wr, w.or)
	w.Withdraw = NewWithdraw(w.Database)
}

func (w *WithdrawSuite) TearDownTest() {
	defer w.ctrl.Finish()
}

func (w *WithdrawSuite) TestShortDoWithdraw() {
	username := "user_name_test"
	coin := "btc"
	value := float64(1.23)

	w.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{Btc: float64(2)}, nil)
	w.wr.EXPECT().Save(gomock.Any()).Return(nil)
	w.or.EXPECT().Insert(gomock.Any()).Return(nil)

	err := w.Withdraw.Execute(username, coin, value)

	assert.Empty(w.T(), err)
}

func (w *WithdrawSuite) TestShortReturnErrorInFindWallet() {
	username := "user_name_test"
	coin := "btc"
	value := float64(1.23)
	errExpect := errors.New("Error expected")

	w.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{Btc: float64(2)}, errExpect)

	err := w.Withdraw.Execute(username, coin, value)

	assert.NotEmpty(w.T(), err)
	assert.EqualError(w.T(), err, err.Error())
}

func (w *WithdrawSuite) TestShortReturnErrorNewOperatio() {
	username := "user_name_test"
	coin := "coin_not_exist"
	value := float64(1.23)
	errExpect := errors.New("invalid currency type")

	w.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{Ada: float64(2)}, nil)

	err := w.Withdraw.Execute(username, coin, value)

	assert.NotEmpty(w.T(), err)
	assert.EqualError(w.T(), err, errExpect.Error())
}

func (w *WithdrawSuite) TestShortReturnErrorInsufficientValue() {
	username := "user_name_test"
	coin := "btc"
	value := float64(1.23)
	errExpect := errors.New("Insufficient portfolio value")

	w.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{Btc: float64(0)}, nil)

	err := w.Withdraw.Execute(username, coin, value)

	assert.NotEmpty(w.T(), err)
	assert.EqualError(w.T(), err, errExpect.Error())
}

func (w *WithdrawSuite) TestShortReturnErrorInSaveWallet() {
	username := "user_name_test"
	coin := "ada"
	value := float64(1.23)
	errExpect := errors.New("Error expected")

	w.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{Ada: float64(2)}, nil)
	w.wr.EXPECT().Save(gomock.Any()).Return(errExpect)

	err := w.Withdraw.Execute(username, coin, value)

	assert.NotEmpty(w.T(), err)
	assert.EqualError(w.T(), err, errExpect.Error())
}

func (w *WithdrawSuite) TestShortReturnErrorInInsertOperation() {
	username := "user_name_test"
	coin := "eth"
	value := float64(1.23)
	errExpect := errors.New("Error expected")

	w.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{Eth: float64(2)}, nil)
	w.wr.EXPECT().Save(gomock.Any()).Return(nil)
	w.or.EXPECT().Insert(gomock.Any()).Return(errExpect)

	err := w.Withdraw.Execute(username, coin, value)

	assert.NotEmpty(w.T(), err)
	assert.EqualError(w.T(), err, errExpect.Error())
}

func (w *WithdrawSuite) TestShortDoDepositWithXrp() {
	username := "user_name_test"
	coin := "xrp"
	value := float64(1.23)

	w.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{Xrp: float64(2)}, nil)
	w.wr.EXPECT().Save(gomock.Any()).Return(nil)
	w.or.EXPECT().Insert(gomock.Any()).Return(nil)

	err := w.Withdraw.Execute(username, coin, value)

	assert.Empty(w.T(), err)
}

func (w *WithdrawSuite) TestShortDoDepositWithDoge() {
	username := "user_name_test"
	coin := "doge"
	value := float64(1.23)

	w.wr.EXPECT().FindByUserName(username).Return(&domain.Wallet{Doge: float64(2)}, nil)
	w.wr.EXPECT().Save(gomock.Any()).Return(nil)
	w.or.EXPECT().Insert(gomock.Any()).Return(nil)

	err := w.Withdraw.Execute(username, coin, value)

	assert.Empty(w.T(), err)
}
