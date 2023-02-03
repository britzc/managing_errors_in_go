package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	heading("Defer Panic Recover Output")
	DeferPanicRecover()
}

func DeferPanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	Print(9, 0)
	fmt.Println("Completed")
}

func Print(x, y int) {
	if y <= 0 {
		panic(fmt.Sprintf("%v", y))
	}

	res := x / y

	fmt.Println("Result:", res)
}

func CopyFileWithoutDefer(srcPath string, dstPath string) (count int64, err error) {
	src, err := os.Open(srcPath)
	if err != nil {
		return 0, err
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		return 0, err
	}

	count, err = io.Copy(src, dst)

	dst.Close()
	src.Close()

	return
}

func CopyFileWithDefer(srcPath string, dstPath string) (count int64, err error) {
	src, err := os.Open(srcPath)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(src, dst)
}

func heading(val string) {
	output := fmt.Sprintf("***** %s *****", val)
	line := strings.Repeat("-", len(output))

	fmt.Println(line)
	fmt.Println(output)
}
