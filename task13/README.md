### 并发编程

#### 并发方式调用匿名函数
```
go func() {
    .....

}()
```

#### WaitGroup
`Add(n)`把计数器设置为n,`Done()`会将计数器每次减1，`Wait()`函数会阻塞代码运行，直到计数器减0

ps.
- 计数器不能为负值
- WaitGroup对象不是引用类型

```
// 这是我们将在每个goroutine中运行的函数。
// 注意，等待组必须通过指针传递给函数。
func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
```

#### Once
`sync.Once`可以控制函数只能被调用一次，不能多次重复调用

#### 互斥锁Mutex 

使用`sync.Mutex`读操作与写操作都会被阻塞。其实读操作的时候我们是不需要进行阻塞的，因此sync中还有另一个锁：读写锁RWMutex,这是一个单写多读模型。

`sync.RWMutex`分为：读、写锁。在读锁占用下，会阻止写，但不会阻止读，多个goroutine可以同时获取读锁，调用`RLock()`函数即可，`RUnlock()`函数释放。写锁会阻止任何goroutine进来，整个锁被当前goroutine，此时等价于Mutex,写锁调用Lock启用，通过UnLock()释放

#### 条件变量Cond
`sync.Cond`是条件变量，它可以让一系列的 Goroutine 都在满足特定条件时被唤醒

条件变量通常与互斥锁一起使用，条件变量可以在共享资源的状态变化时通知相关协程。 经常使用的函数如下：
- NewCond
创建一个Cond的条件变量
```
func NewCond(l Locker) *Cond
```
- Broadcast
广播通知，调用时可以加锁，也可以不加
```
func (c *Cond) Broadcast()
```
- Signal
单播通知，只唤醒一个等待c的goroutine
```
func (c *Cond) Signal()
```
- Wait
等待通知, Wait()会自动释放c.L，并挂起调用者的goroutine。之后恢复执行，Wait()会在返回时对c.L加锁，除非被Signal或者Broadcast唤醒，否则Wait()不会返回
```
func (c *Cond) Wait()
```

#### 原子操作
在`sync/atomic`中，提供了一些原子操作，包括加法（Add）、比较并交换（Compare And Swap，简称 CAS）、加载（Load）、存储（Store）和交换（Swap）

##### 加法操作
供了32/64位有符号与无符号加减操作

##### 比较并交换
如果addr和old相同,就用new代替addr
```
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
```

##### 交换
不管旧值与新值是否相等，都会通过新值替换旧值，返回的值是旧值
```
func SwapInt32(addr *int32, new int32) (old int32)
```

##### 加载（体现不出来）
当读取该指针指向的值时，CPU 不会执行任何其它针对此值的读写操作
```
func LoadInt32(addr *int32) (val int32)
```

##### 存储
加载逆向操作
```
var xx int32 = 1
var yy int32 = 2
atomic.StoreInt32(&yy, atomic.LoadInt32(&xx))
fmt.Println(xx, yy)
```

##### 原子类型
sync/atomic中添加了一个新的类型Value
```
v := atomic.Value{}
v.Store(1)
fmt.Println(v.Load())
```

#### 临时对象池Pool
`sync.Pool`可以作为临时对象的保存和复用的集合

对于需要重复分配、回收内存的地方，`sync.Pool`是一个很好的选择。减少GC负担,如果Pool中有对象，下次直接取，不断服用对象内存，减轻 GC 的压力，提升系统的性能

```
var pool *sync.Pool

type Foo struct {
	Name string
}

func Init() {
	pool = &sync.Pool{
        // 设置池子里能装入的对象
		New: func() interface{} {
			return new(Foo)
		},
	}
}

func main() {
	fmt.Println("Init p")
	Init()

	p := pool.Get().(*Foo)
	fmt.Println("第一次取：", p)
	p.Name = "bob"
	pool.Put(p)

	fmt.Println("池子有对象了，调用获取", pool.Get().(*Foo))
	fmt.Println("池子空了", pool.Get().(*Foo))
}
```

#### Channel

##### 1）使用
Channel的使用需要通过make创建
```
unBufferChan := make(chan int)  // 无缓冲
bufferChan := make(chan int, x)  // 有缓冲
```
读写操作
```
ch := make(chan int, 1)
// 读操作
x <- ch
// 写操作
ch <- x
```
关闭
```
close(ch)
```
当channel关闭后会引发下面相关问题：
- 重复关闭Channel会panic
- 向关闭的Channel发数据会Panic，读关闭的Channel不会Panic，但读取的是默认值

Channel本身的值是默认值又或者是读到的是关闭后的默认值，可以通过下面进行区分
```
val, ok := <-ch
if ok == false {
    // channel closed
}
```

##### 2）Channel分类
- 无缓冲的Channel
发送与接受同时进行。如果没有Goroutine读取Channel(<-Channel)，发送者(Channel<-x)会一直阻塞

- 有缓冲的Channel
发送与接受并非同时进行。当队列为空，接受者阻塞;队列满，发送者阻塞

#### Select
- 每个case都必须是一个通信
- 所有channel表达式都会被求值
- 如果没有default语句，select将阻塞，直到某个通信可以运行
- 如果多个case都可以运行，select会随机选择一个执行

##### 1）随机选择
select特性之一：随机选择，下面会随机打印不同的case结果
```
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
```
假设chan中没有值，有可能引发死锁，如下例
```
ch := make(chan int, 1)
select {
case <-ch:
	fmt.Println("ch 1")
case <-ch:
	fmt.Println("ch 2")
}
```
加上default即可解决

另外，还可以添加超时
```
timeout := make(chan bool, 1)
go func() {
	time.Sleep(2 * time.Second)
	timeout <- true
}()
ch := make(chan int, 1)

select {
case <-ch:
	fmt.Println("ch 1")
case <-timeout:
	fmt.Println("timeout 1")
case <-time.After(time.Second * 1):
	fmt.Println("timeout 2")
}
```

##### 2） 检查chan
select+defaul方式来确保channel是否满
```
ch := make(chan int, 1)
ch <- 1
select {
case ch <- 1:
	fmt.Println("channel value is ", <-ch)
	fmt.Println("channel value is ", <-ch)
default:
	fmt.Println("channel blocking")
}
```
如果要调整channel大小，可以在make的时候改变size，这样就可以在case中往channel继续写数据

##### 3） 选择循环
当多个channel需要读取数据的时候，就必须使用 for+select

例如：下面例子需要从两个channel中读取数据，当从channel1中数据读取完毕后，会像signal channel中输入stop，此时终止for+select

```
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
```
