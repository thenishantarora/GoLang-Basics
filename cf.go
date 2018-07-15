//Nishant Arora

package main 

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
	
	// fmt.Println("Hello")
 	in := bufio.NewReader(os.Stdin)
  	var n int
  	// fmt.Fscanf(in, "%d\n", &n)	
	// var n int
	var x int
    m := make(map[int]int)
	ans := 0
  	fmt.Fscanf(in, "%d\n", &n)	
    // fmt.Println(n)
	for i := 0; i < n ; i++ {
  		fmt.Fscanf(in, "%d", &x)	
	// fmt.Println(x)
		if(x!=0) {
			m[x] = m[x] + 1
			if(m[x]==1) {
				ans++
			}
		}	
	}

	fmt.Println(ans)

}