package student

import (
	"acmicpc_checker_v2_backend/model"
	"errors"

	"gorm.io/gorm"
)

func createUnique(db *gorm.DB, student *model.Student) error {
	var check model.Student
	res := db.Where("email = ?", student.Email).Find(&check)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected > 0 {
		return errors.New("student duplicated")
	}

	if create := db.Create(student); create.Error != nil {
		return create.Error
	}

	return nil
}

func get(db *gorm.DB, id int) (*model.Student, error) {
	var find model.Student
	res := db.Where(id).Find(&find)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("no find student")
	}

	return &find, nil
}

func update(db *gorm.DB, value *model.Student) error {
	if result := db.Save(value); result.Error != nil {
		return result.Error
	}

	return nil
}
