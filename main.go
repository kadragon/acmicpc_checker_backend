package main

import (
	"acmicpc_checker_v2_backend/acmicpc"
	"acmicpc_checker_v2_backend/assignment"
	classinfo "acmicpc_checker_v2_backend/classInfo"
	classstudentinfo "acmicpc_checker_v2_backend/classStudentInfo"
	"acmicpc_checker_v2_backend/db"
	"acmicpc_checker_v2_backend/student"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.Init()

	// Web Server
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/student", student.ListHandler)
	e.POST("/student", student.CreateHandler)
	e.GET("/student/:id", student.GetHandler)
	e.PUT("/student/:id", student.UpdateHandler)
	e.DELETE("/student/:id", student.DeleteHandler)

	assignmentGroup := e.Group("/assignment")
	assignmentGroup.GET("", assignment.ListHandler)
	assignmentGroup.POST("", assignment.CreateHandler)
	assignmentGroup.GET("/:id", assignment.GetHandler)
	assignmentGroup.PUT("/:id", assignment.UpdateHandler)
	assignmentGroup.DELETE("/:id", assignment.DeleteHandler)

	e.GET("/acmicpc/convert/:id", acmicpc.GetProblemHandler)
	e.GET("/acmicpc/check/:id", acmicpc.GetCheckHandler)

	classInfoGroup := e.Group("/classInfo")
	classInfoGroup.GET("", classinfo.ListHandler)
	classInfoGroup.GET("/:id", classinfo.GetHandler)
	classInfoGroup.PUT("/:id", classinfo.UpdateHandler)
	classInfoGroup.GET("/student/:id", classinfo.ListStudentHandler)

	classStudentGroup := e.Group("/classStudent")
	classStudentGroup.GET("", classstudentinfo.ListHandler)
	classStudentGroup.POST("", classstudentinfo.CreateHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

// func main() {
// 	e := echo.New()

// 	// Middleware
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	// Routes
// 	e.GET("/users", getAllUsers)
// 	e.POST("/users", createUser)
// 	e.GET("/users/:id", getUser)
// 	e.PUT("/users/:id", updateUser)
// 	e.DELETE("/users/:id", deleteUser)

// 	// Start server
// 	e.Logger.Fatal(e.Start(":1323"))
// }
