结构体中字段的类型可以是任何类型，包括函数类型，接口类型，甚至结构体类型本身

在声明结构体时我们也可以不给字段指定名字，例如下面这样,这种我们称其为**匿名字段**
- ps.对于一个结构体来说，每一种**数据类型**只能有一个匿名字段
```
type Person struct {
	ID string
	int
}
```

#### 操作结构体
- 使用new函数会创建一个指向结构体类型的指针，创建过程中会自动为结构体分配内存，结构体中每个变量被赋予对应的零值
- 也可以使用第二种方式生命结构类型，需要注意的是此时给结构体赋值的顺序需要与结构体字段声明的顺序一致
- 第三种方式更为常用，我们创建结构体的同时显示的为结构体中每个字段进行赋值

```
s1 := new(Student) //第一种方式
s2 := Student{"james", 35} //第二种方式
s3 := &Student { //第三种方式
	Name: "LeBron",
	Age:  36,
}
```

#### 标签
标记的tag只有reflect包可以访问到，一般用于orm或者json的数据传递
```
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

#### 内嵌结构体
结构体作为一种数据类型也可以将其生命为匿名字段，此时我们称其为内嵌结构体，当结构体中有变量命名有冲突时，由外到内优先级递减

#### 方法
方法与函数类似，只不过在方法定义时会在func和方法名之间增加一个参数，其中r被称为方法的接收者
```
func (r Receiver)func_name(){
  // body
}
```

使用值接收者定义的方法，在调用的时使用的其实是值接收者的一个拷贝，所以对该值的任何操作，都不会影响原来的类型变量。

但是如果使用指针接收者的话，在方法体内的修改就会影响原来的变量，因为指针传递的也是地址，但是是指针本身的地址，此时拷贝得到的指针还是指向原值的，所以对指针接收者操作的同时也会影响原来类型变量的值。

而且在go语言中还有一点比较特殊，我们使用值接收者定义的方法使用指针来调用也是可以的，反过来也是如此

#### 接口
接口相当于一种规范，它需要做的是谁想要实现我这个接口要做哪些内容，而不是怎么做

空接口是一个比较特殊的类型，因为其内部没有定义任何方法所以空接口可以表示任何一个类型
```
var any interface{}

any = 1
fmt.Println(any)

any = "hello"
fmt.Println(any)

any = false
fmt.Println(any)
```