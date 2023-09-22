package models

import "fmt"

type Course struct {
	id           int
	name         string
	prerequisite *Course
}

func NewCourse(id int, name string, prerequisite *Course) *Course {
	return &Course{
		id,
		name,
		prerequisite,
	}
}

func (c *Course) SetName(name string) {
	c.name = name
}
func (c *Course) GetName() string {
	return c.name
}

func (c *Course) SetPrerequisite(anotherCourse *Course) {
	c.prerequisite = anotherCourse
}
func (c *Course) GetPrerequisite() *Course {
	return c.prerequisite
}

func (c *Course) GetInfo() string {
	var nullCourse *Course
	if c.prerequisite == nullCourse {
		return fmt.Sprintf("Course info: (ID: %d, Title: %s)", c.id, c.name)
	}
	return fmt.Sprintf("Course info: (ID: %d, Title: %s, Prerequisite: %s)", c.id, c.name, c.prerequisite.GetName())
}
