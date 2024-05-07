package services

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal("Error loading .env.test file")
	}
	var connectionString = fmt.Sprintf("user=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
	testDB, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		log.Fatalf("cannot connect to database")
	}
	db = testDB
}

func TestGetStudents(t *testing.T) {

	studentService := NewStudentService(db)

	var students = studentService.GetStudents()
	var student = students[0]
	var expected = Student{
		Id:  1,
		Age: 8,
	}
	if student.Id != expected.Id || student.Age == expected.Age {
		t.Errorf("student not as expected")
	}
}
