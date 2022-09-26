# Chapter 5 控制结构

## 备注

- `select` 结构，用于 channel 的选择
- `for`
  - `for-range` 结构:`for pos, char := range str {`
  - `for index, rune := range str2 {`
- `if-else` 结构
- `f, err := os.Open(name) if err != nil {`
- `switch` 结构
- `label` & `goto`: `for`、`switch` 或 `select` 语句都可以配合标签使用