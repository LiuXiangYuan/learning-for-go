go语言在传递数组时会对其进行拷贝，所以如果传递的是大数组的话会非常占内存，所以一般情况下很少直接传递一个数组，避免这种情况我们可以使用以下两种方式：
- 传递数组的指针
- 传递切片