package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BalanceSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	or       *mock.MockOperationRepository
	wr       *mock.MockWalletRepository
	Database *database.Database
	Balance  *Balance
}

func TestBalanceSuite(t *testing.T) {
	suite.Run(t, new(BalanceSuite))
}

func (b *BalanceSuite) SetupTest() {
	b.ctrl = gomock.NewController(b.T())
	b.or = mock.NewMockOperationRepository(b.ctrl)
	b.wr = mock.NewMockWalletRepository(b.ctrl)
	b.Database = database.NewDatabase(&gorm.DB{}, b.wr, b.or)
	b.Balance = NewBalance(b.Database)
}

func (b *BalanceSuite) TearDownTest() {
	defer b.ctrl.Finish()
}

func (b *BalanceSuite) TestShortGetCorrectValues() {
	username := "user_name_test"
	valueExpect := map[string]Bal{
		"btc": Bal{
			Input:  1.2,
			Output: 0.2,
			Total:  1,
		},
	}

	b.or.EXPECT().FindByUserName(username).Return([]domain.Operation{
		domain.Operation{
			UserName:  username,
			Type:      "deposit",
			Coin:      "btc",
			Value:     float64(1.2),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		domain.Operation{
			UserName:  username,
			Type:      "withdraw",
			Coin:      "btc",
			Value:     float64(0.2),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil)

	balaco, err := b.Balance.Execute(username)

	assert.Empty(b.T(), err)
	assert.NotEmpty(b.T(), balaco)
	assert.Equal(b.T(), valueExpect, balaco)
}

func (b *BalanceSuite) TestShortReturnError() {
	username := "user_name_test"
	valueExpect := map[string]Bal{}
	errExpect := errors.New("Error expected")

	b.or.EXPECT().FindByUserName(username).Return([]domain.Operation{}, errExpect)

	balaco, err := b.Balance.Execute(username)

	assert.NotEmpty(b.T(), err)
	assert.Empty(b.T(), balaco)
	assert.Equal(b.T(), valueExpect, balaco)
	assert.EqualError(b.T(), err, errExpect.Error())
}
