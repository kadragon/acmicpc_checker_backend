package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Email     string `gorm:"unique" json:"email" query:"email" form:"email"`
	AcmicpcID string `gorm:"unique" json:"acmicpc_id" form:"email"`
	Rname     string `json:"rname" query:"rname" form:"email"`
	Solved    int    `json:"solved" query:"solved" form:"solved"`
	Grade     int    `json:"grade" query:"grade" form:"grade"`
	Class     int    `json:"class" query:"class" form:"class"`
}

type ClassInfo struct {
	gorm.Model
	Year       int    `json:"year" query:"year" form:"year"`
	Semester   int    `gorm:"check:semester < 5" json:"semester" query:"semester" form:"semester"`
	Subject    string `gorm:"index:idx_class_info,unique" json:"subject" query:"subject" form:"subject"`
	ClassCount int    `gorm:"index:idx_class_info,unique" json:"class_count" query:"class_count" form:"class_count"`
}

type ClassStudent struct {
	gorm.Model
	ClassInfoID uint `gorm:"index:idx_class_student,unique" json:"class_info_id" query:"class_info_id" form:"class_info_id"`
	StudentID   uint `gorm:"index:idx_class_student,unique" json:"student_id" query:"student_id" form:"student_id"`
}

type Assignment struct {
	gorm.Model
	ClassInfoID     uint   `gorm:"index:idx_assignment,unique" json:"class_info_id" query:"class_info_id" form:"class_info_id"`
	Name            string `gorm:"index:idx_assignment,unique" json:"name" query:"name" form:"name"`
	AssignmentCount int    `json:"assignment_count" query:"assignment_count" form:"assignment_count"`
	StartTime       string `json:"start_time" query:"start_time" form:"start_time"`
	EndTime         string `json:"end_time" query:"end_time" form:"end_time"`
	ProblemList     string `json:"problem_list" query:"problem_list" form:"problem_list"`
}

type SolvedInfo struct {
	gorm.Model
	StudentAID string    `gorm:"index:idx_solvedInfo,unique" json:"student_a_id" query:"student_a_id" form:"student_a_id"`
	ProblemID  int       `gorm:"index:idx_solvedInfo,unique" json:"problem_id" query:"problem_id" form:"problem_id"`
	SolvedTime time.Time `json:"solved_time" query:"solved_time" form:"solved_time"`
}
