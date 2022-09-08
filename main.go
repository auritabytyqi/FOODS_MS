package main

import (
	"net/http"

	"FOODS_MS/controller"
	"FOODS_MS/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()

	// Routes
	e.GET("/", start)
	e.GET("request/api/foods", controller.GetFoods)
	e.GET("request/api/foods/:id", controller.GetFood)
	e.GET("request/api/AddFood", controller.AddFood)
	e.GET("request/api/DeleteFood/:id", controller.DeleteFood)
	e.GET("request/api/UpdateFood/values", controller.UpdateFood)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func start(c echo.Context) error {
	return c.String(http.StatusOK, "Starting app........ ")
}
