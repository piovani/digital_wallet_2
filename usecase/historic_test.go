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

type HistoricSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	or       *mock.MockOperationRepository
	wr       *mock.MockWalletRepository
	Database *database.Database
	Historic *Historic
}

func TestHistoricSuite(t *testing.T) {
	suite.Run(t, new(HistoricSuite))
}

func (h *HistoricSuite) SetupTest() {
	h.ctrl = gomock.NewController(h.T())
	h.or = mock.NewMockOperationRepository(h.ctrl)
	h.wr = mock.NewMockWalletRepository(h.ctrl)
	h.Database = database.NewDatabase(&gorm.DB{}, h.wr, h.or)
	h.Historic = NewHistoric(h.Database)
}

func (h *HistoricSuite) TearDownTest() {
	defer h.ctrl.Finish()
}

func (h *HistoricSuite) TestShortReturnHistoic() {
	username := "user_name_test"

	h.or.EXPECT().FindByUserName(username).Return(
		[]domain.Operation{
			domain.Operation{},
			domain.Operation{},
		},
		nil,
	)

	operations, err := h.Historic.Execute(username)

	assert.Empty(h.T(), err)
	assert.Len(h.T(), operations, 2)
}

func (h *HistoricSuite) TestShortReturnHistoicWithError() {
	username := "user_name_test"
	errExpect := errors.New("Error expected")

	h.or.EXPECT().FindByUserName(username).Return(
		[]domain.Operation{},
		errExpect,
	)

	operations, err := h.Historic.Execute(username)

	assert.NotEmpty(h.T(), err)
	assert.EqualError(h.T(), err, errExpect.Error())
	assert.Len(h.T(), operations, 0)
}
