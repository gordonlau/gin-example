package main

import (
	"fmt"
	"gordonlau/gin-example/controllers"
	"gordonlau/gin-example/services"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func InitApp(r *gin.Engine) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var connectionString = fmt.Sprintf("user=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
	db, error := sqlx.Connect("postgres", connectionString)
	if error != nil {
		log.Fatalln(error)
	}
	studentService := services.NewStudentService(db)
	studentController := controllers.NewStudentController(studentService)

	r.GET("/", studentController.GetStudents)
}

func main() {

	r := gin.Default()
	InitApp(r)
	r.Run()
}
