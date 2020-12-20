package main

import (
	"fmt"
	"errors"
	"strconv"
)

func GetSum(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func GetSum1(num1, num2 int) int {
	result := num1 + num2
	return result
}

func paramFunc(a int, b *int, c []int) {
	a = 100
	*b = 100
	c[1] = 100

	fmt.Println("paramFunc:")
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(c)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("The divisor cannot be zero.")
	}
	return a / b, nil
}

func div2(a, b float64) (result float64, err error) {
	if b == 0 {
		return 0, errors.New("被除数不能等于0")
	}
	result = a / b
	return
}

func add() func(int) int {
	n := 10
	str := "string"
	return func(x int) int {
		n = n + x
		str += strconv.Itoa(x)
		fmt.Print(str, " ")
		return n
	}
}

func main ()  {
	a := 1
	b := 1
	c := []int{1, 2, 3}
	paramFunc(a, &b, c)

	fmt.Println("main:")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	slice := []int{7, 9, 3, 5, 1}
	x := min(slice...)
	fmt.Printf("The minimum is: %d", x)

	result, err := div(1, 2)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Println("result: ", result)

	if r, e := div(1, 0); err != nil {
		fmt.Printf("error: %v", e)
	} else {
		fmt.Println("result: ", r)
	}

	// 匿名函数
	f := func() string {
		return "hello world"
	}
	fmt.Println(f())

	f1 := add()
	fmt.Println(f1(1))
	fmt.Println(f1(2))
	fmt.Println(f1(3))

	f1 = add()
	fmt.Println(f1(1))
	fmt.Println(f1(2))
	fmt.Println(f1(3))
}