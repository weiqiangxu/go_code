package p1

// 不可以做寻址的操作有:
// 1. 常量
// 2. 字面量
// 3. 函数返回值（部分情况）
// 4. 基本类型的临时值

const myConstant = 10

func tool() {
	// 以下这行代码会报错，因为常量不可寻址
	// &myConstant
	// fmt.Println(&myConstant)

	// 尝试对数字字面量寻址，会报错
	// &10
	// fmt.Println(&10)

	// 尝试对字符串字面量寻址，会报错
	// &"hello"
	// fmt.Println(&"hello")

	// 尝试对布尔字面量寻址，会报错
	// &true
	// fmt.Println(&true)

	// 直接对函数返回的临时值寻址，会报错
	// &returnValue()
	// fmt.Println(&returnValue())

	// 对表达式中的临时整数取值，会报错
	// &(10 + 5)
	// fmt.Println(&(10 + 5))

	// 对表达式中的临时字符串拼接结果取值，会报错
	// &("go" + "lang")
	// fmt.Println(&("go" + "lang"))
}

func returnValue() int {
	return 20
}
