#### error
1. 如果函数需要处理异常，通常将error作为多值返回的最后一个值，返回的error值为nil则表示无异常，非nil则是有异常。
2. 一般先用if语句处理error!=nil，正常逻辑放if后面。

#### panic
1. 如果在defer函数之外调用recover，则不会停止panicking的序列
2. defer和recover必须在panic之前定义，否则无效