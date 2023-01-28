package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("test")

	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	i++
}

func Sample() {
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

func CopyFileOld(srcPath string, dstPath string) (count int64, err error) {
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

func CopyFile(srcPath string, dstPath string) (count int64, err error) {
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
