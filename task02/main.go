package main

import (
	"fmt"
	"math"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	var b bool = true
	fmt.Println("bool output:", b)
	var x float64 = 3.0
	var y float64 = 4.0
	var name complex128 = complex(x, y)
	fmt.Println("complex output:", name)
	z := complex(x, y)
	var _x = real(z)
	fmt.Println("real output:", _x)
	var _y = imag(z)
	fmt.Println("imag output:", _y)

	fmt.Println("max float32:", math.MaxFloat32)
	fmt.Println("max float64:", math.MaxFloat64)

	const e = .71828
	const Planck = 6.62606957e-34

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	var flist []func()
	for i := 0; i < 3; i++ {
		i := i
		flist = append(flist, func() {
			fmt.Println(i)
		})
	}

	for _, f := range flist {
		f()
	}
}
