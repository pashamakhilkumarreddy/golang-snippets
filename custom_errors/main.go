package main

import (
	"fmt"
	"log"
	"math"
)

type areaError struct {
	err    string  // err description
	radius float64 // radius which caused the error
}

type rectError struct {
	err           string
	length, width float64
}

func (e *rectError) Error() string {
	return e.err
}

func (e *rectError) lengthNegative() bool {
	return e.length < 0
}

func (e *rectError) widthNegative() bool {
	return e.width < 0
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}

	}
	if err != "" {
		return 0, &rectError{err, length, width}
	}
	return length * width, nil
}

func calculateArea(radius float64) (float64, error) {
	if radius < 0 {
		// return 0, errors.New("area calculation failed, radius is less than zero")
		// return 0, fmt.Errorf("area calculation failed, radius %0.2f is less than zero", radius)
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	log.Println("Custom errors in Golang")

	func() {
		radius := -12.0

		area, err := calculateArea(radius)
		if err != nil {
			if err, ok := err.(*areaError); ok {
				fmt.Printf("radius %0.2f is less then zero\n", err.radius)
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Println("Area of circle is", area)
	}()

	func() {
		length, width := -12.0, -3.0
		area, err := rectArea(length, width)
		if err != nil {
			if err, ok := err.(*rectError); ok {
				if err.lengthNegative() {
					fmt.Printf("error: length %0.2f is less than 0\n", length)
				}
				if err.widthNegative() {
					fmt.Printf("error: width  %0.2f is less than 0\n", width)
				}
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Println("Area of the rectangle is ", area)
	}()
}
