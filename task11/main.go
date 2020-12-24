package main

import (
	"fmt"
	"reflect"
)

func hello() {
	fmt.Println("Hello world!")
}

func main() {
	var Num float64 = 3.14

	v := reflect.ValueOf(Num)
	t := reflect.TypeOf(Num)

	fmt.Println("v的可写性:", v.CanSet())

	fmt.Println("Reflect : Num.Value = ", v)
	fmt.Println("Reflect : Num.Type  = ", t)

	origin := v.Interface().(float64)
	fmt.Println(origin)

	var f float64 = 3.41
	fmt.Println(f)
	p := reflect.ValueOf(&f)
	fmt.Println("Reflect : Num.Value = ", p)
	va := p.Elem()
	va.SetFloat(6.18)
	fmt.Println(f)

	hl := hello
	fv := reflect.ValueOf(hl)
	fv.Call(nil)
}
