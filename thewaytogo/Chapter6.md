# Chapter 6 Function

## About

- DRY (Don't Repeat Yourself)原则:

## init 函数

- 全局声明中初始化，也可以在 `init()` 函数中初始化；它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 `main()` 函数高。

  `init()` 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine，如下面这个例子当中的 `backend()`：

## 类型

- 普通的带有名字的函数
- 匿名函数或者lambda函数
- 方法（Methods）
- 内置函数

## Notices

- 申明一个在外部定义的函数，你只需要给出函数名与函数签名，不需要给出函数体`func flushICache(begin, end uintptr) // implemented externally`
- 函数也可以以申明的方式被使用，作为一个函数类型`type binOp func(int, int) int`
- 按值传递 (call by value) 按引用传递 (call by reference)
  - 在函数调用时，像切片 (slice)、字典 (map)、接口 (interface)、通道 (channel) 这样的引用类型都是默认使用引用传递
- 返回值
  - 返回多个参数：
    - `func getX2AndX3(input int) (int, int) { return 2 * input, 3 * input}`
    - `func getX2AndX3_2(input int) (x2 int, x3 int) { x2 = 2 * input x3 = 3 * input}`
  - 空白符 (blank identifier) `_`:返回值赋值空白符自动丢弃
- 传递变长参数
  - 称为变参函数:`func myFunc(a, b, arg ...int) {}`
  - 变参使用：

  ```go
  func F1(s ...string) {
    F2(s...)
    F3(s)
  }

  func F2(s ...string) { }
  func F3(s []string) { }
  ```

  - 变长参数可以作为对应类型的 slice 进行二次传递
  - 变长参数的类型不相同：
    - 使用结构传参：`type Options struct { par1 type1, par2 type2, ...}`
    - 使用空接口传参:
    
    ```go
    func typecheck(..,..,values … interface{}) {
      for _, value := range values {
        switch v := value.(type) {
          case int: …
          case float: …
          case string: …
          case bool: …
          default: …
        }
      }
    }
    ```

- defer: 推迟到函数返回之前一刻才执行某个语句或函数
  - 用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源
  - 使用场景：
    - 关闭文件流`defer file.Close()`
    - 解锁一个加锁的资源`mu.Lock()  defer mu.Unlock()`
    - 打印最终报告`printHeader() defer printFooter()`
    - 关闭数据库链接`defer disconnectFromDB()`
    - **使用 `defer` 语句实现代码追踪**
    - **使用 `defer` 语句来记录函数的参数与返回值**
- 内置函数
  |名称|说明|
  |---|---|
  |`close()`|用于管道通信|
  |`len()`、`cap()`|`len()` 用于返回某个类型的长度或数量（字符串、数组、切片、`map` 和管道）；`cap()` 是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道，不能用于 `map`）|
  |`new()`、`make()`|`new()` 和 `make()` 均是用于分配内存：`new()` 用于值类型和用户定义的类型，如自定义结构，`make` 用于内置引用类型（切片、`map` 和管道）。它们的用法就像是函数，但是将类型作为参数：`new(type)`、`make(type)`。`new(T)` 分配类型 `T` 的零值并返回其地址，也就是指向类型 `T` 的指针（详见[第 10.1 节](10.1.md)）。它也可以被用于基本类型：`v := new(int)`。`make(T)` 返回类型 `T` 的初始化之后的值，因此它比 `new()` 进行更多的工作（详见[第 7.2.3/4 节](07.2.md)、[第 8.1.1 节](08.1.md)和[第 14.2.1 节](14.2.md)）。**`new()` 是一个函数，不要忘记它的括号**。|
  |`copy()`、`append()`|用于复制和连接切片|
  |`panic()`、`recover()`|两者均用于错误处理机制|
  |`print()`、`println()`|底层打印函数（详见[第 4.2 节](04.2.md)），在部署环境中建议使用 `fmt` 包|
  |`complex()`、`real ()`、`imag()`|用于创建和操作复数（详见[第 4.5.2.2 节](04.5.md)）|
- 递归函数(只是算法，没有go的特性)
- 将函数作为参数`func callback(y int, f func(int, int)) {}`
  - 将函数作为参数的最好的例子是函数 `strings.IndexFunc()`和`strings.Map()`
- 闭包:`func(x, y int) int { return x + y }`
  - 不能独立存在,可以赋值给变量或直接调用`func(x, y int) int { return x + y } (3, 4)`
  - 应用：将函数作为返回值、goroutine 和管道操作
  - 应用：使用闭包输出调试信息`where := func() { _, file, line, _ := runtime.Caller(1) log.Printf("%s:%d", file, line) }`
- 高阶函数:返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，是函数式语言的特点
- 计算函数执行时间`start := time.Now() end := time.Now() delta := end.Sub(start)`
- 通过内存缓存来提升性能
  - 进行大量的计算时，提升性能最直接有效的一种方式就是避免重复计算。通过在内存中缓存和重复利用相同计算的结果，称之为内存缓存
  - 譬如大量进行相同参数的运算(例如：纯函数中，即相同输入必定获得相同输出的函数)