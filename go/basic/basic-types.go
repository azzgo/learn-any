package main

import "fmt"

func main() {
	// 内置类型

	// 整型: int8 int32 int64 uint8 uint32 uint64
	// int uint 和系统有关
	// uintptr 宽度不定，编译器需要保证能存下内存地址
	// byte 是 unit8 的别名，rune是 int32 的别名
	var interger = 123
	var letter rune = 'a'

	// 浮点: float32 float64
	var decimal = 121_239.12 // 数字允许使用 _ 增强可读性

	// 复数：complex64 complex128
	var complexNumber = (12 + 21i)

	// 字符串：stirng
	var word = "hello"

	word = `
	我们是
	🇨🇳人
	` // 多行字符串

	// 布尔
	const isTruth = false
	fmt.Println("类型：")
	fmt.Printf("intergar type: %T\n", interger)
	fmt.Printf("letter type: %T\n", letter)
	fmt.Printf("decimal type: %T\n", decimal)
	fmt.Printf("complexNumber type: %T\n", complexNumber)
	fmt.Printf("word type: %T\n", word)
	fmt.Printf("isTruth type: %T\n", isTruth)

	// --未声明的类型有默认值-------------------
	// 布尔的默认值 false
	var isMale bool
	// 数值型类型的默认值 0
	var sumCurrency int
	// 字符串的默认值 空字符串
	var bookTitle string

	fmt.Println("默认值：")
	fmt.Println("isMale:", isMale)
	fmt.Println("sumCurrency:", sumCurrency)
	fmt.Println("bookTitle:", bookTitle)
}
