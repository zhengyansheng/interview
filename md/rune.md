# rune

## 1. 概述

```go
// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.
type byte = uint8

// 翻译
byte 是uint8别名，在所有方面等价于uint8。
它被使用，用来区分字节值和8个无符号值

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
type rune = int32

// 翻译
rune 是int32别名，在所有方面等价于int32。
它被使用，用来区分字符值和整数值
```



```tex
1 byte -> 8-bit(int8)

1 int32 -> 4 * int8


// 翻译
1个字节 等价于 8个bit； 8个bit等价于 uint8
1个int32 等价于 32-bit； 等价于 4 uint8


```





## 2. 示例

```go
func main() {
	printLen()
	intercept()
}

// 打印长度
func printLen() {
	s := "中国"
	fmt.Printf("byte len: %d\n", len(s))         // byte
	fmt.Printf("rune len: %d\n", len([]rune(s))) // unicode
}

// 截取
func intercept() {
	s := "中国"
	fmt.Printf("index 0: %v\n", s[0])                      // byte
	fmt.Printf("rune index 0: %v\n", string([]rune(s)[0])) // unicode
}
```



```bash
# go run main.go
byte len: 6
rune len: 2
index 0: 228
rune index 0: 中
```



## 3. 总结

> string 底层是通过byte数组实现的。
>
> string 中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。



- byte 等同于int8，常用来处理ascii字符
- rune 等同于int32, 常用来处理unicode或utf-8字符



## 4. 参考

https://learnku.com/articles/23411/the-difference-between-rune-and-byte-of-go

https://www.jianshu.com/p/4fbf529926ca