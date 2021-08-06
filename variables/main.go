package main

import "fmt"

func main() {
	fmt.Println("Variables in Golang")
	var name string
	fmt.Printf("Me llamo %s\n", name)
	name = "John Doe"
	fmt.Printf("Me llamo %s\n", name)
	var age int = 26
	fmt.Printf("My age is %d\n", age)
	var city = "Madrid"
	fmt.Printf("Yo vivo en %s\n", city)
	country := "Espana"
	fmt.Printf("Yo soy de %s\n", country)
	var weight, height int = 80, 180 // int can be dropped
	fmt.Println("Height and Weight are ", weight, height)
	var (
		profession = "medico"
		college    = "No lo se"
	)
	fmt.Println("I am a ", profession, "from ", college)
}
