package assignment

import (
	"acmicpc_checker_v2_backend/db"
	"acmicpc_checker_v2_backend/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateHandler(c echo.Context) error {
	a := &model.Assignment{}

	if err := c.Bind(a); err != nil {
		return c.String(http.StatusBadRequest, "binding error")
	}

	db := db.Dbconnect()
	if err := db.Create(a); err.Error != nil {
		return c.String(http.StatusBadRequest, "assignment create db error")
	}

	return c.JSON(http.StatusCreated, a)
}

func GetHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param invaild")
	}

	db := db.Dbconnect()

	var a model.Assignment
	if err := db.Where(id).Find(&a).Error; err != nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, a)
}

func UpdateHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param no match")
	}

	var a model.Assignment
	if err := c.Bind(a); err != nil {
		return c.String(http.StatusBadRequest, "data binding error")
	}

	if a.ID != uint(id) {
		return c.String(http.StatusBadRequest, "param and query no match")
	}

	db := db.Dbconnect()
	if err := db.Save(&a); err != nil {
		return c.String(http.StatusBadRequest, "update fail")
	}

	return c.JSON(http.StatusOK, a)
}

func DeleteHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param error")
	}

	db := db.Dbconnect()
	if err := db.Delete(&model.Assignment{}, id).Error; err != nil {
		return c.String(http.StatusBadRequest, "delete query error")
	}

	return c.NoContent(http.StatusNoContent)
}

func ListHandler(c echo.Context) error {
	condition := new(model.Assignment)
	if err := c.Bind(condition); err != nil {
		return c.String(http.StatusBadRequest, "binding error")
	}

	db := db.Dbconnect()

	var assignments []model.Assignment

	if err := db.Where(&condition).Order("name").Find(&assignments).Error; err != nil {
		return c.String(http.StatusBadRequest, "query error")
	}

	return c.JSON(http.StatusOK, assignments)
}
