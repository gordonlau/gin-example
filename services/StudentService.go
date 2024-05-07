package services

import "github.com/jmoiron/sqlx"

type Student struct {
	Id        int    `db:"id"`
	Age       int    `db:"age"`
	CreatedAt string `db:"created_at"`
}

// go:generate mockery --name StudentService
type StudentService interface {
	GetStudents() []Student
}

type studentService struct {
	db *sqlx.DB
}

// GetStudents implements StudentService.
func (s *studentService) GetStudents() []Student {
	students := []Student{}
	s.db.Select(&students, "SELECT * from students")
	return students
}

func NewStudentService(db *sqlx.DB) StudentService {
	return &studentService{
		db: db,
	}
}
