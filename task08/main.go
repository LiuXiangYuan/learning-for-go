package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Person struct {
	ID   string
	name string
	int
}

func (p Person) GetID() string {
	return p.ID
}

func (p Person) Getname() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func ToJson(s *Student) (string, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

type A struct {
	X, Y int
}

type B struct {
	A
	Name string
}

type C struct {
	A
	B
	X int
}

type D struct {
	a    A
	Name string
}

type Animal interface {
	Eat()
}

type Bird struct {
	Name string
}

func (b Bird) Eat() {
	fmt.Println(b.Name + "吃虫")
}

type Dog struct {
	Name string
}

func (d Dog) Eat() {
	fmt.Println(d.Name + "吃肉")
}

func EatWhat(a Animal) {
	a.Eat()
}

// 类型断言
func IsDog(a Animal) bool {
	if v, ok := a.(Dog); ok {
		fmt.Println(v)
		return true
	}
	return false
}

func WhatType(a Animal) {
	switch a.(type) {
	case Dog:
		fmt.Println("Dog")
	case Bird:
		fmt.Println("Bird")
	default:
		fmt.Println("error")
	}
}

func main() {
	fmt.Println("start!")
	s1 := new(Student)         //第一种方式
	s2 := Student{"james", 35} //第二种方式
	s3 := &Student{            //第三种方式
		Name: "LeBron",
		Age:  36,
	}

	s1.Name = "james"
	s1.Age = 35

	s2.Age++
	s3.Age++

	p := new(Person)
	p.ID = "123"
	p.int = 10
	fmt.Println(p.GetID())
	p.SetName("kobe")
	fmt.Println(p.Getname())

	str, _ := ToJson(s3)
	fmt.Println(str)

	b := new(B)
	b.X = 10
	b.Y = 20
	b.Name = "james"

	c := new(C)
	c.X = 10
	c.Y = 11
	fmt.Println(c.A.X)
	c.A.X = 11
	fmt.Println(c.X)

	d := new(D)
	d.a.X = 10

	bird := Bird{"Bird"}
	dog := Dog{"Dog"}
	EatWhat(bird)
	EatWhat(dog)

	var any interface{}

	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = false
	fmt.Println(any)
}
