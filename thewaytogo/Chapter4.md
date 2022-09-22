# Chapter 4

## Go 程序的基本结构和要素

  包：`package main`

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

  左大括号 { 必须与方法的声明放在同一行

  ```go
  func functionName(parameter_list) (return_value_list) {
    …
  }
  ```

  类型可以是基本类型，如：`int`、`float`、`bool`、`string`；

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

  Go 程序的执行（程序启动）顺序：`main` >> 递归包

  类型转换

  示例：

  ```go
  a := 5.0
  b := int(a)
  ```

   命名规范：返回某个对象的函数或方法的名称一般都是使用名词，没有 `Get...` 之类的字符，如果是用于修改某个对象，则使用 `SetName()`