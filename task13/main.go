package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// SafeCounter 的并发使用是安全的。
// type SafeCounter struct {
// 	v   map[string]int
// 	mux sync.Mutex
// }

// // Inc 增加给定 key 的计数器的值。
// func (c *SafeCounter) Inc(key string) {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
// 	c.v[key]++
// }

// // Value 返回给定 key 的计数器的当前值。
// func (c *SafeCounter) Value(key string) int {
// 	c.mux.Lock()
// 	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
// 	defer c.mux.Unlock()
// 	return c.v[key]
// }

type SafeCounter struct {
	v     map[string]int
	rwmux sync.RWMutex
}

func (c *SafeCounter) Inc(key string) {
	// 写操作使用写锁
	c.rwmux.Lock()
	fmt.Println("write")
	defer c.rwmux.Unlock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	// 读的时候加读锁
	c.rwmux.RLock()
	fmt.Println("read")
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.rwmux.RUnlock()
	return c.v[key]
}

var doOnce sync.Once
var pool *sync.Pool

type Foo struct {
	Name string
}

func Init() {
	pool = &sync.Pool{
		New: func() interface{} {
			return new(Foo)
		},
	}
}

func f1(c chan int, s chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		c <- i
	}
	s <- "stop"
}

func f2(c chan int, s chan string) {
	for i := 20; i >= 0; i-- {
		time.Sleep(time.Second)
		c <- i
	}
	s <- "stop"
}

func main() {
	go func() {
		fmt.Println("you forgot me !")
	}()

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Println("after loop")
	wg.Wait()

	DoSomething()
	DoSomething()

	time.Sleep(time.Second)

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(c.Value("somekey"))
	}

	time.Sleep(time.Second)

	var j int64
	atomic.AddInt64(&j, 1)
	fmt.Println("j = j + 1 =", j)
	atomic.AddInt64(&j, -1)
	fmt.Println("j = j - 1 =", j)

	var a int32 = 1
	var b int32 = 2
	var d int32 = 3
	ok := atomic.CompareAndSwapInt32(&a, a, b)
	fmt.Printf("ok = %v, a = %v, b = %v\n", ok, a, b)
	ok = atomic.CompareAndSwapInt32(&a, d, b)
	fmt.Printf("ok = %v, a = %v, b = %v, d = %v\n", ok, a, b, d)

	var x int32 = 1
	var y int32 = 2
	old := atomic.SwapInt32(&x, y)
	fmt.Println(x, old)

	var x1 int32 = 1
	y1 := atomic.LoadInt32(&x1)
	fmt.Println("x1, y1:", x1, y1)

	var xx int32 = 1
	var yy int32 = 2
	atomic.StoreInt32(&yy, atomic.LoadInt32(&xx))
	fmt.Println(xx, yy)

	v := atomic.Value{}
	v.Store(1)
	fmt.Println(v.Load())

	fmt.Println("Init p")
	Init()

	p := pool.Get().(*Foo)
	fmt.Println("第一次取：", p)
	p.Name = "bob"
	pool.Put(p)

	fmt.Println("池子有对象了，调用获取", pool.Get().(*Foo))
	fmt.Println("池子空了", pool.Get().(*Foo))

	ch := make(chan int, 1)
	ch <- 1
	select {
	case <-ch:
		fmt.Println("ch 1")
	case <-ch:
		fmt.Println("ch 2")
	default:
		fmt.Println("ch default")
	}

	/////////////////////////////
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	chh := make(chan int, 1)

	select {
	case <-chh:
		fmt.Println("ch 1")
	case <-timeout:
		fmt.Println("timeout 1")
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 2")
	}

	ch2 := make(chan int, 1)
	ch2 <- 1
	select {
	case ch2 <- 1:
		fmt.Println("channel value is ", <-ch2)
		fmt.Println("channel value is ", <-ch2)
	default:
		fmt.Println("channel blocking")
	}

	c1 := make(chan int)
	c2 := make(chan int)
	signal := make(chan string, 10)

	go f1(c1, signal)
	go f2(c2, signal)
LOOP:
	for {
		select {
		case data := <-c1:
			fmt.Println("c1 data is ", data)
		case data := <-c2:
			fmt.Println("c2 data is ", data)
		case data := <-signal:
			fmt.Println("signal is ", data)
			break LOOP
		}
	}
}

func DoSomething() {
	doOnce.Do(func() {
		fmt.Println("Run once - first time, loading...")
	})
	fmt.Println("Run this every time")
}
