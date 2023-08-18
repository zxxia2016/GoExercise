# Chapter 7 

## About

- 数组与切片(array & slice)
- 剖析 **集合**，它是可以包含大量条目 (item) 的数据结构，例如数组、切片和 `map`

## 数组

- 特点：**唯一类型**的序列
  - 类型：基本类型、自定义类型、任意类型（用空接口做为类型）
- 语法：`[5]int` `var identifier [len]type`
- `for`遍历
  - `for i:=0; i < len(arr1); i++ {`
  - `for i,_:= range arr1 {`