package classinfo

import (
	classstudentinfo "acmicpc_checker_v2_backend/classStudentInfo"
	"acmicpc_checker_v2_backend/db"
	"acmicpc_checker_v2_backend/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param invaild")
	}

	db := db.Dbconnect()

	var class model.ClassInfo
	if err := db.Where(id).Find(&class).Error; err != nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, class)
}

func UpdateHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param no match")
	}

	var class model.ClassInfo
	if err := c.Bind(class); err != nil {
		return c.String(http.StatusBadRequest, "data binding error")
	}

	if class.ID != uint(id) {
		return c.String(http.StatusBadRequest, "param and query no match")
	}

	db := db.Dbconnect()
	if err := db.Save(&class); err != nil {
		return c.String(http.StatusBadRequest, "update fail")
	}

	return c.JSON(http.StatusOK, class)
}

func ListHandler(c echo.Context) error {
	condition := new(model.Assignment)
	if err := c.Bind(condition); err != nil {
		return c.String(http.StatusBadRequest, "binding error")
	}

	db := db.Dbconnect()
	var classInfos []model.ClassInfo

	if err := db.Where(&condition).Find(&classInfos).Error; err != nil {
		return c.String(http.StatusBadRequest, "query error")
	}

	return c.JSON(http.StatusOK, classInfos)
}

func ListStudentHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id param no match")
	}

	var students []model.Student

	db := db.Dbconnect()

	uid := uint(id)
	studentIDList := classstudentinfo.GetClassStudentIDList(db, uid)

	if len(studentIDList) == 0 {
		return c.JSON(http.StatusNoContent, students)
	}

	if err := db.Order("grade, class, rname").Where(studentIDList).Find(&students).Error; err != nil {
		return c.String(http.StatusBadRequest, "query error")
	}

	return c.JSON(http.StatusOK, students)
}
