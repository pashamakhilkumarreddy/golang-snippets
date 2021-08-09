package main

import "fmt"

// func functionName(parameterName type) returnType {}

func main() {
	fmt.Println("Functions in Golang")
	fmt.Println("Total Bill = ", calculateBill(12, 9))
	area, perimeter := rectProps(12, 10)
	fmt.Printf("Area, perimeter is %f %f\n", area, perimeter)
	sqArea, sqPerimeter := squareProps(12)
	fmt.Printf("Area, perimeter of the square is %f %f\n", sqArea, sqPerimeter)
}

func calculateBill(price int, num int) int { // func calculateBill(price, num int) int
	return price * num
}

func rectProps(length, width float64) (float64, float64) {
	var area = width * length
	var perimeter = (length + width) / 2
	return area, perimeter
}

func squareProps(side float64) (area, perimeter float64) { // Named return values
	area = side * side
	perimeter = 4 * side
	return
}
