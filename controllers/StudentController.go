package controllers

import (
	"gordonlau/gin-example/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// go:generate mockery --name StudentController
type StudentController interface {
	GetStudents(c *gin.Context)
}

type studentController struct {
	studentService services.StudentService
}

func (s *studentController) GetStudents(c *gin.Context) {
	students := s.studentService.GetStudents()
	c.JSON(http.StatusOK, students)
}

func NewStudentController(s services.StudentService) StudentController {
	return &studentController{
		studentService: s,
	}
}
