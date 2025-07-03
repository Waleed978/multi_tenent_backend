package routes

import (
	"github.com/Waleed978/multi_tenent_backend/handlers"

	"github.com/gin-gonic/gin"
)

// SetupStudentRoutes sets up the API routes for student-related operations.
func SetupStudentRoutes(router *gin.Engine, studentHandler *handlers.StudentHandler) {
	// Create a group for student routes
	studentRoutes := router.Group("/students")
	{
		studentRoutes.POST("/", studentHandler.CreateStudent)
		studentRoutes.GET("/", studentHandler.GetStudents)
		studentRoutes.GET("/:id", studentHandler.GetStudent)
		studentRoutes.PUT("/:id", studentHandler.UpdateStudent)
		studentRoutes.DELETE("/:id", studentHandler.DeleteStudent)
	}
}
