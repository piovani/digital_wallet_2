package infra

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/piovani/digital_wallet_2/infra/config"
)

func Heart(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("API running on the port %d", config.Env.ApiPort))
}
