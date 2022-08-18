package operation

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/piovani/digital_wallet_2/api/controller"
	"github.com/piovani/digital_wallet_2/usecase"
)

type OperationController struct {
	UsecaseDeposit  *usecase.Deposit
	UsecaseWithdraw *usecase.Withdraw
	UsecaseHistoric *usecase.Historic
	UsecaseBalance  *usecase.Balance
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewOperationController(
	d *usecase.Deposit,
	w *usecase.Withdraw,
	h *usecase.Historic,
	b *usecase.Balance,
) *OperationController {
	return &OperationController{
		UsecaseDeposit:  d,
		UsecaseWithdraw: w,
		UsecaseHistoric: h,
		UsecaseBalance:  b,
	}
}

func (o *OperationController) Deposit(c echo.Context) error {
	var input Input
	c.Bind(&input)

	if err := ValidInput(input); err != nil {
		return c.JSON(http.StatusBadRequest, controller.NewResponseError(err))
	}

	if err := o.UsecaseDeposit.Execute(input.UserName, input.Coin, input.Value); err != nil {
		return c.JSON(http.StatusInternalServerError, controller.NewResponseError(err))
	}

	return c.NoContent(http.StatusCreated)
}

func (o *OperationController) Withdraw(c echo.Context) error {
	var input Input
	c.Bind(&input)

	if err := ValidInput(input); err != nil {
		return c.JSON(http.StatusBadRequest, controller.NewResponseError(err))
	}

	if err := o.UsecaseWithdraw.Execute(input.UserName, input.Coin, input.Value); err != nil {
		return c.JSON(http.StatusInternalServerError, controller.NewResponseError(err))
	}

	return c.NoContent(http.StatusCreated)
}

func (o *OperationController) Historic(c echo.Context) error {
	username := c.FormValue("user_name")

	operations, err := o.UsecaseHistoric.Execute(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller.NewResponseError(err))
	}

	return c.JSON(http.StatusOK, GetOpertationsOutput(operations))
}

func (o *OperationController) Balance(c echo.Context) error {
	username := c.FormValue("user_name")

	balance, err := o.UsecaseBalance.Execute(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller.NewResponseError(err))
	}

	return c.JSON(http.StatusOK, balance)
}
