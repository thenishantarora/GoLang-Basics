package main

import "fmt"

type Rectangle struct {
	
	Length int
	Breadth int
}

func main() {
	

	// rect := Rectangle{5,4}
	var rect Rectangle
	rect.Length = 7
	rect.Breadth = 4
	fmt.Println(rect)	
}
