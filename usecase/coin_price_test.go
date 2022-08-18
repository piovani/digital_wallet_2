package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CoinPriceSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	CoinPrice *CoinPrice
}

func TestCoinPriceSuite(t *testing.T) {
	suite.Run(t, new(CoinPriceSuite))
}

func (c *CoinPriceSuite) SetupTest() {
	ctx := context.TODO()

	c.ctrl = gomock.NewController(c.T())
	c.CoinPrice = NewCoinPrice(ctx)
}

func (c *CoinPriceSuite) TearDownTest() {
	defer c.ctrl.Finish()
}

func (c *CoinPriceSuite) TestShortGetCorrectValuesPrice() {
	valueExpect := map[string]float64{
		"BTC_USD": 1,
		"BTC_EUR": 1,

		"ETH_USD": 1,
		"ETH_EUR": 1,

		"ADA_USD": 1,
		"ADA_EUR": 1,

		"XRP_USD": 1,
		"XRP_EUR": 1,

		"DOG_USD": 1,
		"DOG_EUR": 1,
	}

	price, err := c.CoinPrice.Execute()

	assert.Empty(c.T(), err)
	assert.NotEmpty(c.T(), price)
	assert.Equal(c.T(), valueExpect, price)
}
