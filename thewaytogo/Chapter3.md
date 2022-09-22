# Chapter 3

- IDE：各种各样，自己选择
- 调试器
  - `gdb`
  - `fmt.Printf`
  - `panic()`获取栈跟踪信息
  - `defer`跟踪代码执行过程
- 格式化代码
  - 在命令行输入`gofmt –w program.go`
- 生成代码文档
  - `go doc`
- 与其它语言进行交互
  - 与 C 进行交互`import "C"`
  - 与 C++ 进行交互`SWIG`