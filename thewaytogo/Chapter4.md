# Chapter 4

## Go 程序的基本结构和要素

- 包：`package main`

  导入: `import "fmt"; import "os"`

  ```go
  import (
    "fmt"
    "os"
  )
  ```
  
  如果包名不是以 `.` 或 `/` 开头，会在全局文件进行查找；如果包名以 `./` 开头，则 Go 会在相对目录中查找；如果包名以 `/` 开头（在 Windows 下也可以这样使用），则会在系统的绝对路径中查找

  包也可以作为命名空间使用，帮助避免命名冲突（名称冲突）

  别名：`import fm "fmt"`

- 左大括号 { 必须与方法的声明放在同一行

  ```go
  func functionName(parameter_list) (return_value_list) {
    …
  }
  ```

- 类型可以是基本类型，如：`int`、`float`、`bool`、`string`；

  结构化的（复合的），如：`struct`、`array`、切片 (slice)、`map`、通道 (channel)；

  只描述类型的行为的，如：`interface`。

  结构化的类型没有真正的值，它使用 `nil` 作为默认值

  函数也可以是一个确定的类型，就是以函数作为返回类型

  使用 `type` 关键字可以定义你自己的类型

  ```go
  type IZ int
  ```

  ```go
  type (
    IZ int
    FZ float64
    STR string
  )
  ```

- Go 程序的执行（程序启动）顺序：`main` >> 递归包

- 类型转换

  示例：

  ```go
  a := 5.0
  b := int(a)
  ```

- 命名规范：返回某个对象的函数或方法的名称一般都是使用名词，没有 `Get...` 之类的字符，如果是用于修改某个对象，则使用 `SetName()`

## 常量(const)

- 用作枚举

- 位运算

## 变量

- 声明变量的一般形式是使用 `var` 关键字：`var identifier type`。

- 值类型和引用类型(值赋值是拷贝，引用就是指针)

- 打印:`fmt.Print("Hello:", 23)`

- 简短形式，使用 := 赋值操作符;不可以再次对于相同名称的变量使用初始化声明，例如：`a := 20`又`a := 20`

## 基本类型和运算符

- `int、uint 和 uintptr`
  - `int` 和 `uint` 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。
  - `uintptr` 的长度被设定为足够存放一个指针即可。
- 整数：`int8` `int16` `int32``int64`；无符号整数`+u`
- 浮点型:`float32` `float64`
- 整型的零值为 `0`，浮点型的零值为 `0.0`。
- 复数

    ```go
    complex64 (32 位实数和虚数)
    complex128 (64 位实数和虚数)
    ```

- 位运算:只能用于整数类型的变量:`&` `|` `^` 位清除 `&^`：将指定位置上的值设置为 `0`。 `<<` `>>`
- 随机数:`rand.Int()` `rand.Intn(8) [0, n)之间` `rand.Seed(timens)` `rand.Float32()`

## 字符串

- `Join 拼接` `s += "world!"` `bytes.Buffer` `HasPrefix 以开头` `HasSuffix 以结尾` `Contains` `Index` `LastIndex` `Replace` `Count` `Repeat` `ToLower` `ToUpper` `Trim TrimSpace TrimLeft TrimRight修剪字符串` `Fields Split 分割字符串` `NewReader`
- `strconv` 包
- 判断一个字符串是否为空：
  - `if str == "" { ... }`
  - `if len(str) == 0 {...}`

## 时间和日期

- `time` 包

## 指针

- 取地址符是 `&`
- `%p` 输出地址
- `var intP *int` `var p *type`
- 取值`var == *(&var)`