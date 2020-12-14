参照: https://blog.csdn.net/AdolphKevin/article/details/105480530

代理配置参照: https://goproxy.io/zh/

若采用vscode，则在配置完代理后需要重启下vscode

1. 安装地址: https://studygolang.com/dl

终端输入go version有回应即安装成功

2. VScode中安装Go插件

3. 编写.go文件

#### ps
- package main定义了包名。必须在源文件中非注释的第一行指明这个文件属于哪个包。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
- import "fmt"告诉编译器程序运行需要用fmt包。
- func main() 是程序开始执行的函数，main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
- {}中"{"不可以单独放一行。
- /.../ 是注释，在程序执行时将被忽略。//单行注释， /* ... */ 多行注释也叫块注释,不可以嵌套使用，一般用于包的文档描述或注释成块的代码片段。
- fmt.Println(...) 将字符串输出到控制台，并在最后自动增加换行字符 \n。用 fmt.Print("hello, world\n") 可以得到相同的结果。