package main

import "fmt"

func main() {
	
	var x int
	x = 10
	sum := 0
	for i := 0; i < x ; i++ {
		sum += i
	}



	fmt.Println(sum)
}