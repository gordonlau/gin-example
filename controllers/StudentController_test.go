package controllers

import (
	"gordonlau/gin-example/services"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/ovechkin-dm/mockio/mock"
)

func TestGetStudents(t *testing.T) {
	SetUp(t)
	m := Mock[services.StudentService]()
	students := []services.Student{
		{
			Id:        10,
			Age:       30,
			CreatedAt: "2024-05-07",
		},
	}
	WhenSingle(m.GetStudents()).ThenReturn(students)
	studentController := NewStudentController(m)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	studentController.GetStudents(c)
	Verify(m, AtLeastOnce()).GetStudents()
}
