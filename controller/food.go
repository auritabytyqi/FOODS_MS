package controller

import (
	"net/http"

	"FOODS_MS/model"
	"FOODS_MS/storage"

	"github.com/labstack/echo/v4"
)

func GetFood(c echo.Context) error {
	id := c.Param("id")
	foods, _ := GetRepoFoods()
	for i := 0; i < len(foods); i++ {
		if foods[i].Id == id {
			return c.JSON(http.StatusOK, foods[i])
		}
	}
	return c.JSON(http.StatusOK, "Food doesn't exist")
}

func GetFoods(c echo.Context) error {
	foods, _ := GetRepoFoods()
	return c.JSON(http.StatusOK, foods)
}

func GetRepoFoods() ([]model.Foods, error) {
	db := storage.GetDBInstance()
	foods := []model.Foods{}

	if err := db.Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

func AddFood(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	description := c.QueryParam("description")
	id_exists := FoodExists(id)
	if !id_exists {
		food := model.Foods{Id: id, Name: name, Description: description}
		storage.AddFoodRecord(food)
		return c.JSON(http.StatusOK, "Food is created")
	}
	return c.JSON(http.StatusOK, "Food product exists with this id. Try another one...")
}

func DeleteFood(c echo.Context) error {
	id := c.Param(("id"))
	id_exists := FoodExists(id)
	if id_exists {
		storage.DeleteFood(id)
		return c.JSON(http.StatusOK, "Food is deleted")
	}
	return c.JSON(http.StatusOK, "Food doesn't exist")
}

func UpdateFood(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	description := c.QueryParam("description")
	id_exists := FoodExists(id)
	if id_exists {
		storage.UpdateFood(id, name, description)
		return c.JSON(http.StatusOK, "Food is updated")
	}
	return c.JSON(http.StatusOK, "Food doesn't exist")
}

func FoodExists(id string) bool {
	foods, _ := GetRepoFoods()
	id_exists := false
	for i := 0; i < len(foods); i++ {
		if foods[i].Id == id {
			id_exists = true
		}
	}
	return id_exists
}
