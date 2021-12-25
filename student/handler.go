package student

import (
	"acmicpc_checker_v2_backend/db"
	"acmicpc_checker_v2_backend/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateHandler(c echo.Context) error {
	s := &model.Student{}

	if err := c.Bind(s); err != nil {
		return c.String(http.StatusBadRequest, "binding error")
	}

	// database input
	db := db.Dbconnect()
	if err := createUnique(db, s); err != nil {
		return c.JSON(http.StatusNotAcceptable, err)
	}

	return c.JSON(http.StatusCreated, s)
}

func GetHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param no match")
	}

	db := db.Dbconnect()

	s, err := get(db, id)
	if err != nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, &s)
}

func UpdateHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param no match")
	}

	var stu model.Student
	if err := c.Bind(stu); err != nil {
		return c.String(http.StatusBadRequest, "data binding error")
	}

	if stu.ID != uint(id) {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := update(db.Dbconnect(), &stu); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, stu)
}

func DeleteHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param error")
	}

	db := db.Dbconnect()
	if err := db.Delete(&model.Student{}, id).Error; err != nil {
		return c.String(http.StatusBadRequest, "delete query error")
	}

	return c.NoContent(http.StatusNoContent)
}

func ListHandler(c echo.Context) error {
	student := new(model.Student)
	if err := c.Bind(student); err != nil {
		return c.String(http.StatusBadRequest, "binding error")
	}

	db := db.Dbconnect()

	var students []model.Student

	if err := db.Where(&student).Find(&students).Error; err != nil {
		return c.String(http.StatusBadRequest, "query error")
	}

	return c.JSON(http.StatusOK, students)
}
