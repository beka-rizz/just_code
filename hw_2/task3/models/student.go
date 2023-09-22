package models

import "fmt"

type Degree string

const (
	Bachelor Degree = "bachelor's degree"
	Master   Degree = "master's degree"
	PHD      Degree = "doctor's degree"
)

type Student struct {
	Person
	studentId string
	gpa       float32
	degree    Degree
	takes     []Course
}

func NewStudent(name, surname string, age int, isMarried bool, studentId string, gpa float32, degree Degree, takes []Course) *Student {
	return &Student{
		Person{
			name,
			surname,
			age,
			isMarried,
		},
		studentId,
		gpa,
		degree,
		takes,
	}
}

func (s *Student) SetStudentId(studentId string) {
	s.studentId = studentId
}
func (s *Student) GetStudentId() string {
	return s.studentId
}

func (s *Student) SetGpa(gpa float32) {
	s.gpa = gpa
}
func (s *Student) GetGpa() float32 {
	return s.gpa
}

func (s *Student) SetDegree(degree Degree) {
	s.degree = degree
}
func (s *Student) GetDegree() Degree {
	return s.degree
}

func (s *Student) SetTakes(courses ...Course) {
	s.takes = courses
}

func (s *Student) GetTakes() (courses []string) {
	for _, elem := range s.takes {
		courses = append(courses, elem.GetName())
	}
	return
}

func (s *Student) GetInfo() string {
	return fmt.Sprintf("%s, Student ID: %s, GPA: %f, degree: %v, takes: %s", s.Person.GetInfo(), s.GetStudentId(), s.GetGpa(), s.GetDegree(), s.GetTakes())
}
