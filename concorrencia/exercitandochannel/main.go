package main

import "fmt"

func printaAlgo(c chan int, n int) {

	for i := 0; i < n; i++ {

		fmt.Println(i)
		//c <- i
	}

	//close(c)

}

func main() {
	buffer := 30000
	c := make(chan int, buffer)

	printaAlgo(c, buffer)

	// for i := range c {
	// 	fmt.Println(i)
	// }
}
