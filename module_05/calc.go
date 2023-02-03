package calc

import "errors"

func Calc(x int, y int) (total int, err error) {
	if y < 0 {
		return 0, errors.New("y > 0")
	}
	if x <= y {
		return 0, errors.New("x > y")
	}

	return x + y, nil
}
