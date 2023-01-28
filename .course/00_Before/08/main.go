package main

import "fmt"

func transmit(c chan int, count int) {
	for i := 0; i < count; i++ {
		c <- i
		fmt.Printf("Sent %d\n", i)
	}

	fmt.Println("Done")
}

func transmit(c chan int, count int) {
	for i := 0; i < count; i++ {
		select {
		case c <- i:
			fmt.Printf("Sent %d\n", i)
		default:
			fmt.Printf("Fail %d\n", i)
		}
	}

	fmt.Println("Done")
}

func main() {
	c := make(chan int, 2)

	transmit(c, 5)
}
