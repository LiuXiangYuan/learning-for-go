Link: https://github.com/datawhalechina/go-talent/blob/master/11.%E5%8D%95%E5%85%83%E6%B5%8B%E8%AF%95.md

### 单元测试

#### 对函数进行功能测试
鼠标放在函数上右键，选择GO:Generate Unit Tests For Function即可生成file_test.go文件

#### 单测要点

1. 单元测试的时候，如果有一些打印log信息，我们运行xxx_test.go是输出不出来的，此时需要使用
```
go test xxx_test.go -v
```
2. 单测覆盖率，覆盖率可以简单理解为进行单元测试mock的时候，能够覆盖的代码行数占总代码行数的比率，当然是高一点要好些。可以通过`-cover`指定

#### 基准测试

基准测试函数名字必须以Benchmark开头，代码在xxx_test.go中


#### mock/stub测试

该命令待测试，未能正常运行
```
mockgen -source=db.go -destination=db_mock.go -package=db
```

#### 浏览器实时测试

1. `go get github.com/smartystreets/goconvey`
2. `$GOPATH/bin/goconvey`
3. `http://localhost:8080`