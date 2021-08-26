package project

import "fmt"

type Income interface {
	Calculate() int
	Source() string
}

type project struct {
	projectName  string
	biddedAmount int
}

type freeLance struct {
	projectName           string
	noOfHours, hourlyRate int
}

type digitalRevenue struct {
	adName     string
	noOfClicks int
	cPC        float64
}

func NewProject(projectName string, biddedAmount int) project {
	return project{projectName, biddedAmount}
}

func NewFreeLance(projectName string, noOfHours, hourlyRate int) freeLance {
	return freeLance{projectName, noOfHours, hourlyRate}
}

func NewDigitalRevenue(adName string, noOfClicks int, cPC float64) digitalRevenue {
	return digitalRevenue{adName, noOfClicks, cPC}
}

func (fb project) Calculate() int {
	return fb.biddedAmount
}

func (fb project) Source() string {
	return fb.projectName
}

func (tm freeLance) Calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm freeLance) Source() string {
	return tm.projectName
}

func (dr digitalRevenue) Calculate() int {
	return int(dr.cPC) * dr.noOfClicks
}

func (dr digitalRevenue) Source() string {
	return dr.adName
}

func CalculateNetIncome(ic []Income) {
	var netIncome int = 0
	for _, income := range ic {
		fmt.Printf("Income from %s = $%d\n", income.Source(), income.Calculate())
		netIncome += income.Calculate()
	}
	fmt.Printf("Net income of the organization = $%d\n", netIncome)
}
