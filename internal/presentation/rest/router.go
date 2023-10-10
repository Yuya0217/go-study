package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func setupRoutes(e *echo.Echo, h Handlers) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/v1")

	// 医薬
	v1.GET("/medicines", h.MedicineHandler.Find)
	v1.GET("/medicines/:id", h.MedicineHandler.GetByID)
	v1.POST("/medicines", h.MedicineHandler.Create)
	v1.PUT("/medicines/:id", h.MedicineHandler.Update)
	v1.DELETE("/medicines/:id", h.MedicineHandler.GetByID)

}
