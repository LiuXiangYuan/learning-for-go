package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var m1 map[string]int
	m1 = make(map[string]int)
	m1["test"] = 1
	m1["test"] = 2
	fmt.Println(m1["test"])

	m2 := make(map[int]interface{}, 100)
	fmt.Println(len(m2)) // map不可用cap

	m3 := map[string]string{
		"name": "james",
		"age":  "35",
	}
	m3["key1"] = "v1"
	m3["key2"] = "v2"
	m3["key3"] = "v3"
	for key, value := range m3 {
		fmt.Println("key:", key, " value:", value)
	}
	fmt.Println("*****************")
	// 用sclice有序化map 负责存k v，slice负责维护k的有序索引位置
	// 如果要对value有序，则可以sclice存储key后根据value的值进行重排
	var keys []string
	for key := range m3 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("key:", key, " value:", m3[key])
	}
	fmt.Println("*****************")
	fmt.Println(len(m3))
	delete(m3, "key1")
	fmt.Println(len(m3))
	fmt.Println("*****************")

	m4 := map[string]bool{
		"check1": true,
		"check2": false,
	}
	if value, ok := m4["check1"]; ok {
		fmt.Println(value)
		if value, ok := m4["check3"]; !ok {
			fmt.Println(value)
		}
	}

	m := make(map[string]func(a, b int) int)
	m["add"] = func(a, b int) int {
		return a + b
	}
	m["multi"] = func(a, b int) int {
		return a * b
	}
	fmt.Println(m["add"](3, 2))
	fmt.Println(m["multi"](3, 2))

	s := "hello"
	b := []byte(s)
	b[0] = 'g'
	s = string(b)
	fmt.Println(s)

	fmt.Println("*****************")
	s = "hello你好中国"
	fmt.Println(len(s)) //17
	fmt.Println(utf8.RuneCountInString(s))
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}
	fmt.Println()
	fmt.Println("*****************")
	var str string = "This is an example of a string"
	//判断字符串是否以Th开头
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	//判断字符串是否以aa结尾
	fmt.Printf("%t\n", strings.HasSuffix(str, "aa"))
	//判断字符串是否包含an子串
	fmt.Printf("%t\n", strings.Contains(str, "an"))

	i, err := strconv.Atoi("-42") //将字符串转为int类型
	ss := strconv.Itoa(-42)       //将int类型转为字符串
	fmt.Println(i, err)
	fmt.Println(ss)

}
