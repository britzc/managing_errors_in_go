package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Result struct {
	Data []byte
	Err  error
}

func ReadFileWithSameChannel(path string) Result {
	resultChan := make(chan Result)

	go func() {
		time.Sleep(1 * time.Second)

		data, err := os.ReadFile(path)
		resultChan <- Result{Data: data, Err: err}
	}()

	return <-resultChan
}

func SameChannel() {
	result := ReadFileWithSameChannel("sample.txt")

	if result.Err != nil {
		log.Fatal(result.Err)
	} else {
		fmt.Println(string(result.Data))
	}
}
