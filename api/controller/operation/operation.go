package operation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database/mysql/repositories"
)

func Create(c echo.Context) error {
	repo := repositories.NewOperationRepository()

	var input Input
	c.Bind(&input)

	operation := domain.NewOperation(input.UserName, "deposit", input.Value)
	repo.Insert(*operation)

	return c.JSON(http.StatusCreated, operation)
}
