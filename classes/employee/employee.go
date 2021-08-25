package employee

import "fmt"

type employee struct {
	firstName, lastName string
	leaves, leavesTaken int
}

func NewEmployee(firstName, lastName string, leaves, leavesTaken int) employee {
	e := employee{firstName, lastName, leaves, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, (e.leaves - e.leavesTaken))
}
