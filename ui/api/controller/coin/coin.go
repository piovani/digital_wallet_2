package coin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/piovani/digital_wallet_2/usecase"
)

type CoinController struct {
	UsecaseCoinPrice *usecase.CoinPrice
}

func NewCoinController(usecaseCoinPrice *usecase.CoinPrice) *CoinController {
	return &CoinController{
		UsecaseCoinPrice: usecaseCoinPrice,
	}
}

// Price godoc
// @Summary     Price from Coins
// @Description Price show current values from coins
// @Tags        coin
// @Produce     json
// @Success 200 {object} Output
// @failure 400 {object} controller.ResponseError
// @Router /coin/price [get]
func (c *CoinController) Price(e echo.Context) error {
	res, err := c.UsecaseCoinPrice.Execute()
	if err != nil {
		return e.String(http.StatusInternalServerError, "Erro ao obter a contação")
	}

	return e.JSON(http.StatusOK, NewOutput(res))
}
