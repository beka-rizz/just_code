package models

import "fmt"

type Person struct {
	name, surname string
	age           int
	isMarried     bool
}

func NewPerson(name, surname string, age int, isMarried bool) *Person {
	return &Person{
		name,
		surname,
		age,
		isMarried,
	}
}

func (p *Person) SetName(name string) {
	p.name = name
}
func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetSurname(surname string) {
	p.surname = surname
}
func (p *Person) GetSurname() string {
	return p.surname
}

func (p *Person) SetAge(age int) {
	p.age = age
}
func (p *Person) GetAge() int {
	return p.age
}
func (p *Person) SetIsMarried(isMarried bool) {
	p.isMarried = isMarried
}
func (p *Person) GetIsMarried() bool {
	return p.isMarried
}

func (p *Person) GetInfo() string {
	return fmt.Sprintf("Name: %s, Surname: %s, Age: %d, Married: %t", p.GetName(), p.GetSurname(), p.GetAge(), p.GetIsMarried())
}
