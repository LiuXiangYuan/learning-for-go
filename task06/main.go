package main

import "fmt"

func test(a [5]int) {
	a[1] = 2
	fmt.Println(a)
}

func test1(a [5]*int) {
	*a[1] = 2
	for i := 0; i < 5; i++ {
		fmt.Print(" ", *a[i])
	}
	fmt.Println()
}

func test2(aPtr *[5]int) {
	aPtr[1] = 5
	fmt.Println(aPtr)
}

func main() {
	//方式一
	var arr1 = [5]int{}
	//方式二
	var arr2 = [5]int{1,2,3,4,5}
	//方式三
	var arr3 = [5]int{3:10}

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 10
	}

	for index, value := range arr2 {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	for index, value := range arr1 {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	fmt.Println(arr3)
	test(arr3)
	fmt.Println(arr3)

	var a [5]*int
	fmt.Println(a)
	for i := 0; i < 5; i++ {
		temp := i
		a[i] = &temp
	}
	for i := 0; i < 5; i++ {
		fmt.Print(" ", *a[i])
	}
	fmt.Println()
	test1(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", *a[i])
	}

	var q [5]int
	var aPtr *[5]int
	aPtr = &q //这样简短定义也可以aPtr := &a
	fmt.Println(aPtr)
	test2(aPtr)
	fmt.Println(aPtr)

	//方法一
	var s1 = []int{}
	//方法二
	var s2 = []int{1, 2, 3}
	//方法三
	var s3 = make([]int, 5)
	//方法四
	var s4 = make([]int, 5, 10)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	arr := [5]int{1, 2, 3, 4, 5}
	s := []int{6, 7, 8, 9, 10}

	s1 = arr[2:4]
	s2 = arr[:3]
	s3 = arr[2:]
	s4 = s[1:3]

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	s1 = append(s1, s2...)
	fmt.Println("s1:", s1)

	g := []int{1, 2, 3}
	b := make([]int, 3)
	copy(b, g)
	fmt.Println(g)
	fmt.Println(b)
}