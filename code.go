//Nishant Arora

package main

import (
        "fmt"
        "bufio"
        "os"
        // "strconv"
)

func main() {
	
	// fmt.Println("Hello")
	// var text string
	// var x int
    // m := make(map[int]int)

    // inputreader := bufio.NewReader(os.Stdin)
    reader := bufio.NewScanner(os.Stdin)
    // reader.Text()

	text := reader.Text()
	// n,  _ := strconv.Atoi(text)

    // input, _ := inputreader.ReadString(n)	
        // _, err := fmt.Scanf("%d", &n)

	// ans := 0
 //    fmt.Scan(&n)
 //    var m [100005]int
 //    // fmt.Println(n)
	// for i := 0; i < n ; i++ {
	// 	fmt.Scan(&x)
	// // fmt.Println(x)
	// 	if(x!=0) {
	// 		m[x] = m[x] + 1
	// 		if(m[x]==1) {
	// 			ans++
	// 		}
	// 	}	
	// }

	fmt.Println(text)

}