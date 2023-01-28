package main

import "fmt"

var ErrDivideByZer = erros.New("The Divisor cannot be zero")

func main() {
	result, err := Divide(9, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("The divisor cannot be zero")
	}
	return a / b, nil
}

type error interface {
	Error() string
}

func Create(name string) (*File, error) {
	return nil, nil
}

func ReadAll(r Reader) ([]byte, error) {
	return nil, nil
}

func Write(_p []byte) (n int, error) {
	return nil, nil
}
