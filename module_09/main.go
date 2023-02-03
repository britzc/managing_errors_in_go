package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	heading("Seperate Channel Output")
	SeperateChannel()

	heading("Same Channel Output")
	SameChannel()

	files := []string{
		"error.txt",
		"sample01.txt",
		"sample02.txt",
		"sample03.txt",
	}

	heading("WaitGroup Output")
	WaitGroup(files)

	heading("ErrorGroup Output")
	ErrorGroup(files)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(5*time.Second))
	defer cancel()

	heading("ErrorGroup With Context Output")
	ErrorGroupContext(ctx, files)
}

func heading(val string) {
	output := fmt.Sprintf("***** %s *****", val)
	line := strings.Repeat("-", len(output))

	fmt.Println(line)
	fmt.Println(output)
}
