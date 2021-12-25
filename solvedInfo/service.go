package solvedInfo

import (
	"acmicpc_checker_v2_backend/model"
	"errors"

	"gorm.io/gorm"
)

func Create(db *gorm.DB, s *model.SolvedInfo) error {
	if err := db.Create(s).Error; err != nil {
		return err
	}

	return nil
}

func Get(db *gorm.DB, s *model.SolvedInfo) (*model.SolvedInfo, error) {
	res := db.Where(s).Find(s)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no find solvedinfo")
	}

	return s, nil
}

func CheckedList(db *gorm.DB, acmicpc_id_list []string, problem_list []int) []model.SolvedInfo {
	var result []model.SolvedInfo
	db.Model(&model.SolvedInfo{}).Order("student_a_id, problem_id").Where("student_a_id in ? and problem_id in ?", acmicpc_id_list, problem_list).Find(&result)

	return result
}
