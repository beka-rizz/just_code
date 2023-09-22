package task3

import (
	"fmt"
	"homeworks/hw_2/task3/models"
)

type InformInterface interface {
	GetInfo() string
}

func Run() {
	p1 := models.NewPerson("Bekarys", "Nurgassym", 19, false)
	s1 := models.NewStudent("Neznakomyi", "Patsan", 20, false, "21B001", 3.92, models.Degree(models.Bachelor), []models.Course{})
	c1 := models.NewCourse(1, "Math", nil)
	c2 := models.NewCourse(2, "Web Programming", c1)
	fmt.Println(p1.GetInfo())
	fmt.Println(s1.GetInfo())
	fmt.Println(c1.GetInfo(), c2.GetInfo())
	fmt.Println()

	s1.SetTakes(*c1, *c2)
	fmt.Println(s1.GetInfo())
}
