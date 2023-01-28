package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func readFileContent01(path string) (<-chan []byte, <-chan error) {
	dataChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	go func() {
		time.Sleep(1 * time.Second)

		data, err := os.ReadFile(path)
		if err != nil {
			errChan <- err
		} else {
			dataChan <- data
		}

		close(dataChan)
		close(errChan)
	}()

	return dataChan, errChan
}

func main01() {
	dataChan, errChan := readFileContent01("sample.txt")

	if err, open := <-errChan; open {
		log.Fatal(err)
	}
	if data, open := <-dataChan; open {
		fmt.Println(string(data))
	}

	fmt.Println("vim-go")
}

type Result struct {
	Data []byte
	Err  error
}

func readFileContent2(path string) Result {
	resultChan := make(chan Result)

	go func() {
		time.Sleep(1 * time.Second)

		data, err := os.ReadFile(path)
		resultChan <- Result{Data: data, Err: err}
	}()

	return <-resultChan
}

func main() {
	result := readFileContent2("sample.txt")

	if result.Err != nil {
		log.Fatal(result.Err)
	} else {
		fmt.Println(string(result.Data))
	}
}
