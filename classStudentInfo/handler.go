package classstudentinfo

import (
	"acmicpc_checker_v2_backend/db"
	"acmicpc_checker_v2_backend/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateHandler(c echo.Context) error {
	cs := new(model.ClassStudent)

	if err := c.Bind(cs); err != nil {
		return c.String(http.StatusBadRequest, "binding error")
	}

	db := db.Dbconnect()
	if err := db.Create(cs); err != nil {
		return c.String(http.StatusBadRequest, "assignment create db error")
	}

	return c.JSON(http.StatusCreated, cs)
}

func ListHandler(c echo.Context) error {
	condition := new(model.ClassStudent)
	if err := c.Bind(condition); err != nil {
		return c.String(http.StatusBadRequest, "binding error")
	}

	db := db.Dbconnect()
	var classStudents []model.ClassStudent

	if err := db.Where(&condition).Find(&classStudents).Error; err != nil {
		return c.String(http.StatusBadRequest, "query error")
	}

	return c.JSON(http.StatusOK, classStudents)
}
