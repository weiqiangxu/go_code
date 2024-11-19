package p1

import (
	"fmt"
	"sync"
	"time"
)

// for循环字符串byte转string
func testFor() string {
	s := "hello"
	var arr []byte
	// 字符串循环
	for i := 0; i < len(s); i++ {
		arr = append(arr, s[i])
	}
	byteStr := ""
	for _, v := range arr {
		// byte 转字符串
		byteStr += string(v)
	}
	return byteStr
}

// 切片使用
func testSlice() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// 因为arr[startIndex:endIndex]左闭右开
	// 如果start和end相等的话那么
	arr2 := numbers[1:1]
	fmt.Println(arr2)

	// arr[startIndex,endIndex,capIndex]
	// 第三个是容量索引. 容量长度的计算方式是 capIndex-startIndex 容量索引减去起始索引
	// numbers[1:4:3] 这样的会执行失败.容量索引减去起始索引，容量大小为3 - 1 = 2不够存储1-(4-1)的元素
	arr6 := numbers[1:4:4]
	fmt.Println(arr6)

	// 默认下限为0
	// 默认上限为 len(numbers)
	// arr[startIndex:endIndex]  左闭右开
	arr := numbers[:]
	fmt.Println(arr)

	arr1 := numbers[1:2]
	fmt.Println(arr1)

}

// 数组
func testArr() {
	// 数组
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 一个让编译器自动推断长度的数组
	var balance = [...]float32{.0, 1.0, 2.0, 3.4, 7.0}
	fmt.Println(balance)

	var number = 0
	for i := 0; i <= 2; i++ {
		number++
	}
	fmt.Println(number)
}

// channel使用
func testChannel() {
	// 测试一下channel
	ch := make(chan int, 2)
	for i := 0; i < 3; i++ {
		go func(n int) {
			ch <- n
		}(i)
	}
	for i := 0; i < 3; i++ {
		z := <-ch
		fmt.Println(z)
	}
	time.Sleep(time.Second * 5)
}

// 闭包-有函数全局变量
func testNoNameFunc() {
	// 测试闭包得全局变量
	f := inCre()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
}

func inCre() func() int {
	count := 0 // 外部作用域得变量 会累加
	return func() int {
		count++
		return count
	}
}

// 两个defer之间panic
func testDefer() {
	defer func() {
		if r := recover(); r != nil {
			// 执行了
			fmt.Println(1)
		}
	}()
	doSomething()
	defer func() {
		if r := recover(); r != nil {
			// 没有执行
			fmt.Println(2)
		}
	}()
}

func doSomething() {
	numbers := 0
	fmt.Println(1 / numbers)
}

func testRange() {
	var wg = sync.WaitGroup{}
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wg.Add(len(s))
	for _, val := range s {
		go func() {
			defer wg.Done()
			// 输出结果是不确定的
			// 是不是9 取决于执行顺序，这两个goroutine谁更快
			// 如果time.Sleep(1 second)就输出全部是9
			time.Sleep(time.Second)
			fmt.Println(val)
		}()
	}
	wg.Wait()
}
