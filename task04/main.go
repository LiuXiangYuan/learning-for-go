package main

import (
	"fmt"
	"time"
)

func Chann(ch chan int, stopCh chan bool) {
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println("x的类型:%T", i)
	case int:
		fmt.Println("x是int类型")
	case float64:
		fmt.Printf("x 是float64型")
	case func(int) float64:
		fmt.Printf("x 是func(int)型")
	case bool, string:
		fmt.Printf("x 是bool或string型")
	default:
		fmt.Printf("未知型")
	}

	marks := false
	switch marks {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	score := 60
	switch score {
	case 90:
		fmt.Println(90)
	case 80:
		fmt.Println(80)
	case 50, 60, 70:
		if score == 50 {
			fmt.Println("不及格")
		} else if score == 60 {
			fmt.Println("刚及格")
		} else {
			fmt.Println("还凑合")
		}
	default:
		fmt.Printf("猜多少分")
	}

	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("C Recvice", c)
		case s := <-ch:
			fmt.Println("S Receive", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}
