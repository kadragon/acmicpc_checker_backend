package acmicpc

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProblemHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param no match")
	}

	problemList := GetProblemListFromExerciseBook(id)

	return c.JSON(http.StatusOK, map[string]string{"problemList": problemList})
}

func GetCheckHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param no match")
	}

	checkData := SolvedData(uint(id))

	return c.JSON(http.StatusOK, checkData)
}
