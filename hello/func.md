# 函数

- 有三种类型的函数：  
  - 普通的带有名字的函数
  - 匿名函数或者lambda函数
  - 方法（Methods）
- 函数重载是不被允许的
- 实参、形参
  - 引用传递：切片 (slice)、字典 (map)、接口 (interface)、通道 (channel)
- 空白符 (blank identifier)：`_`自动丢弃
- 传递变长参数
  - 用法：
    - 参数类型相同：`Func(params ...int)`
    - 参数类型不同：
      - 使用结构
       ```go
       type Options struct {
        par1 type1,
        par2 type2,
        ...
        }
       ```
      - 使用空接口:
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
  - 变长参数二次传递：`Fun(params...)`
- defer:
  - 用法：延迟到 `return` 语句之前执行，类似JS的`finally`
  - 为什么要用？
    - 用于释放某些已分配资源；
    - 疑问？为什么不直接写在函数最后呢？
    - 合理使用 `defer` 语句能够使得代码更加简洁
  - 应用场景
    - 关闭文件流 `defer file.Close()`
    - 解锁一个加锁的资源 `mu.Lock() defer mu.Unlock()`
    - 打印最终报告`printHeader() defer printFooter()`:看demo`func.go: deferTrace`
    - 关闭数据库链接`defer disconnectFromDB()`
- 内置函数:TODO: implement
- 内存缓存：
  - 目的提升性能
  - 空间换时间，计算过的存起来；
  - 条件：相同输入必定相同输出