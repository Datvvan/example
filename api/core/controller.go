package core

import (
	"net/http"

	"github.com/datvvan/sample/db"
	"github.com/datvvan/sample/util"
	"github.com/labstack/echo/v4"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) GetModule(c echo.Context) error {
	modules, err := db.ModuleFindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, util.SuccessResponse(modules))
}

func (controller *Controller) Detail(c echo.Context) error {
	return c.JSON(http.StatusOK, util.SuccessResponse("Body of core menu 1"))

}
