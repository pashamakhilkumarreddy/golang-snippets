package employee

import "fmt"

type employee struct {
	firstName, lastName string
	salary              int
	position            string
}

func NewEmployee(firstName, lastName string, salary int, position string) employee {
	return employee{firstName, lastName, salary, position}
}

func (e employee) FullName() string {
	return fmt.Sprintf("%s %v", e.firstName, e.lastName)
}
