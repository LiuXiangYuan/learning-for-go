package main

import (
	"fmt"
	"unicode/utf8"
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

const (
	a = iota
	b
	c = 100
	d
	e = iota
)

var (
	v_1     int
	v_2     string = "hello world!"
	v_3     float64
	chinese string = "中文"
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	var a map[string]int
	a = make(map[string]int) // 让map可编辑
	a["a"] = 1
	a["b"] = 2
	a["c"] = 3
	value, ok := a["a"]
	if ok {
		fmt.Println("存在", value)
	} else {
		fmt.Println("不存在")
	}

	for key := range a {
		fmt.Println(key)
	}

	fmt.Println(len(v_2))

	fmt.Println("chinese in ASCII:", len(chinese))
	fmt.Println("chinese in utf-8", utf8.RuneCountInString(chinese))

	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
}
