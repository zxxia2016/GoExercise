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
- 申明一个在外部定义的函数，你只需要给出函数名与函数签名，不需要给出函数体`func flushICache(begin, end uintptr) // implemented externally`
- 函数也可以以申明的方式被使用，作为一个函数类型`type binOp func(int, int) int`