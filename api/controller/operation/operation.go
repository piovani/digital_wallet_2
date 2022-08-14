package operation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/piovani/digital_wallet_2/usecase"
)

type OperationController struct {
	UsecaseDeposit *usecase.Deposit
}

func NewOperationController(d *usecase.Deposit) *OperationController {
	return &OperationController{
		UsecaseDeposit: d,
	}
}

func (o *OperationController) Create(c echo.Context) error {
	var input Input
	c.Bind(&input)

	if err := o.UsecaseDeposit.Execute(input.UserName, input.Value); err != nil {
		c.Error(err)
	}

	return c.NoContent(http.StatusCreated)
}
