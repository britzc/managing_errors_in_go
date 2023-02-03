package main

import (
	"fmt"
	"strings"
)

func ChannelWithoutSwitch(c chan int, count int) {
	for i := 0; i < count; i++ {
		c <- i
		fmt.Printf("Sent %d\n", i)
	}

	fmt.Println("Done")
}

func ChannelWithSwitch(c chan int, count int) {
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
	heading("Channel With Switch Output")
	c2 := make(chan int, 2)
	ChannelWithSwitch(c2, 5)

	heading("Channel Without Switch Output")
	c1 := make(chan int, 2)
	ChannelWithoutSwitch(c1, 5)
}

func heading(val string) {
	output := fmt.Sprintf("***** %s *****", val)
	line := strings.Repeat("-", len(output))

	fmt.Println(line)
	fmt.Println(output)
}
