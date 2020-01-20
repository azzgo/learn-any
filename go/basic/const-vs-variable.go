package main

import "fmt"

// 类型转换
func constVsVariable() {
	// 结果为complex128类型的1.0+0.0i。虚部被舍入了。
	fmt.Println(complex128(1 + -1e-1000i))
	// 结果为float32类型的0.5。这里也舍入了。
	fmt.Println(float32(0.49999999))
	// 只要目标类型不是整数类型，舍入都是允许的。
	fmt.Println(float32(17000000000000000))
	fmt.Println(float32(123))
	fmt.Println(uint(1.0))
	fmt.Println(int8(-123))
	fmt.Println(int16(6 + 0i))
	fmt.Println(complex128(789))

	fmt.Println(string(65))          // "A"
	fmt.Println(string('A'))         // "A"
	fmt.Println(string('\u68ee'))    // "森"
	fmt.Println(string(-1))          // "\uFFFD"
	fmt.Println(string(0xFFFD))      // "\uFFFD"
	fmt.Println(string(0x2FFFFFFFF)) // "\uFFFD"
}

func main() {
	constVsVariable()
}
