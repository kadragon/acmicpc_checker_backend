package classstudentinfo

import (
	"acmicpc_checker_v2_backend/model"

	"gorm.io/gorm"
)

func GetClassStudentIDList(db *gorm.DB, classInfoID uint) []uint {
	var classStudent []model.ClassStudent
	condition := &model.ClassStudent{ClassInfoID: classInfoID}
	db.Where(condition).Find(&classStudent)

	studentIDList := make([]uint, 0)
	for _, t := range classStudent {
		studentIDList = append(studentIDList, t.StudentID)
	}

	return studentIDList
}
