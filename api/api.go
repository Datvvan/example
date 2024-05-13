package api

import (
	"github.com/datvvan/sample/api/core"
	"github.com/labstack/echo/v4"
)

func RegisterAPI(app *echo.Echo) {

	apiGroup := app.Group("api")
	userAPIs(apiGroup)
}

func userAPIs(app *echo.Group) {
	controller := core.NewController()
	group := app.Group("/core")
	group.GET("/get-modules", controller.GetModule)
	group.GET("/detail", controller.Detail)
}
