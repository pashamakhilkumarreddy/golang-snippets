package employee

type Employee struct {
	name     string
	Salary   int
	age      int
	position string
	country  string
}

func NewEmployee(name string, salary, age int, position, country string) Employee {
	e := Employee{name, salary, age, position, country}
	return e
}

func Filter(f func(Employee) bool, employees ...Employee) []Employee {
	var emps []Employee
	for _, v := range employees {
		if f(v) {
			emps = append(emps, v)
		}
	}
	return emps
}
