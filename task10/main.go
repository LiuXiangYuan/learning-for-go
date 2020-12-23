package main

import (
	"errors"
	"fmt"
)

type NotNature float64

func (err NotNature) Error() string {
	return fmt.Sprintf("自然数为大于或等于0的数: %v", float64(err))
}

func Nature(x float64) (float64, error) {
	if x < 0 {
		return 0, NotNature(x)
	} else {
		return x, nil
	}
}

func main() {
	err := errors.New("This is an error")
	if err != nil {
		fmt.Println(err)
	}

	err1 := fmt.Errorf("This is an error2")
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(Nature(1))
	fmt.Println(Nature(-1))

	defer func() {
		fmt.Println("defer 里面的第一个println")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("defer 里面的第二个println")
	}()
	f()
}

func f() {
	fmt.Println("1")
	panic("panic")
	fmt.Println("2")
}
