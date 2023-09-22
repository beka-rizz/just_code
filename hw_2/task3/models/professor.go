package models

import "fmt"

type Professor struct {
	Person
	hIndex     int
	experience int
}

func NewProfessor(name, surname string, age int, isMarried bool, hIndex, experience int) *Professor {
	return &Professor{
		Person{
			name,
			surname,
			age,
			isMarried,
		},
		hIndex,
		experience,
	}
}

func (p *Professor) SetHIndex(hIndex int) {
	p.hIndex = hIndex
}
func (p *Professor) GetHIndex() int {
	return p.hIndex
}

func (p *Professor) SetExperience(experience int) {
	p.experience = experience
}
func (p *Professor) GetExperience() int {
	return p.experience
}

func (p *Professor) GetInfo() string {
	return fmt.Sprintf("%s, h-index: %d, work experience: %d", p.Person.GetInfo(), p.GetHIndex(), p.GetExperience())
}
