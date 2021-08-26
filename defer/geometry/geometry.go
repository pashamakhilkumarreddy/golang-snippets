package geometry

import (
	"fmt"
	"sync"
)

type rect struct {
	length, width int
}

func NewRect(length, width int) rect {
	return rect{length, width}
}

func Rects(rects ...rect) []rect {
	return rects
}

func (r rect) Area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("Rectangle %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("Rectangle %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("Area of the rectangle %v is %v\n", r, area)
}
