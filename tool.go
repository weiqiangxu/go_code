package p1

import (
	"fmt"
	"sync"
	"time"
)

// for循环字符串byte转string
func testFor() string {

	// 零切片
	zeroSlice := make([]int, 0)
	fmt.Printf("zeroSlice: %v, len: %d, cap: %d, is nil: %t\n", zeroSlice, len(zeroSlice), cap(zeroSlice), zeroSlice == nil)

	// 空切片
	var emptySlice []int = []int{}
	fmt.Printf("emptySlice: %v, len: %d, cap: %d, is nil: %t\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)

	// nil 切片
	var nilSlice []int
	fmt.Printf("nilSlice: %v, len: %d, cap: %d, is nil: %t\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)

	s := "hello"

	// string转byte数组
	p := []byte(s)
	z := string(p)
	fmt.Println(z)

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

	// go的slice可以直接用== 比较是否相等吗
	arr11 := []int{1, 2}
	arr22 := []int{1, 2}
	//  不能直接用 == 比较
	// if arr11 == arr22
	fmt.Println(arr11)
	fmt.Println(arr22)

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
			// 没有执行
			fmt.Println(1)
		}
	}()
	// 如果没有这个go的话，那么会打印一个1
	// 如果有的话这个子协程里面的panic没有办法被捕获
	go doSomething()
	defer func() {
		if r := recover(); r != nil {
			// 没有执行
			fmt.Println(2)
		}
	}()
	time.Sleep(time.Second * 2)
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

type Book struct {
	title string
}

func testGetPrivate() {
	b := Book{title: "Go Programming"}
	p := &b
	// 显式解引用
	fmt.Println((*p).title)
}

func deferNumber() {
	var u User
	fmt.Println("u.Name = ", u.Name)
	x := 0
	defer func() {
		// 这里输出得结果是什么
		// 这个x是引用 所以输出得是1
		fmt.Println("deferred x:", x)
	}()
	x = 1
	fmt.Println("Normal x:", x)

	for i := 0; i < 3; i++ {
		defer func(item int) {
			// 函数调用 defer如果是显式值传入就是正确打印 2/1/0
			fmt.Println(item)
		}(i)

		defer func() { // 函数调用
			// 这里会统统打印为 3 也就是最后的值
			fmt.Println("i = ", i)
		}()
	}
}

type User struct {
	Name string
}

func testSwitch() {

	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	// case自带break
	// 显式使用 fallthrough 可以多执行下面的case
	switch marks {
	case 90:
		grade = "A"
		// fallthrough
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

// 字符串和rune数组的转换
func testRuneTrans() {
	s := "今天happy"
	arr := []rune(s)
	newStr := string(arr)
	fmt.Println(newStr)

	// 用单引号（'）包围的字符（如'a'）是一个rune类型的字面量
	// rune类型实际上是int32的别名，用于表示一个 Unicode 码点。
	// 小写字母a对应的十进制值是 97 所以z是97
	var z int32
	z = 'a'
	fmt.Println(z)
	fmt.Println(string(z))
}

func testMapper() {
	m := map[string]int{}
	// 下面的写法是错误的 m 是不能被寻址
	// map的地址是没法获取的
	// map是一种引用类型 m实际上是一个指向底层map数据结构的引用
	// 不像数组等类型可以通过取地址操作符
	// l := *m ???  写错了把兄弟，&才是取地址
	fmt.Println(m)

	arr := [1]int{10}
	arr1 := &arr
	// 这里直接通过数组指针可以改数组的值，不需要额外写 (*arr1)[0] = 3
	arr1[0] = 2
	(*arr1)[0] = 3
	fmt.Println(arr1)

	var newM map[string]int
	// 这里会panic 因为newM还是nil
	newM["name"] = 99
}

func testMapperDelete() {
	myMap := &map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	xx := (*myMap)["a"]
	fmt.Println(xx)

	// 可以对map取地址
	newMap := map[string]interface{}{}
	h := &newMap
	v := (*h)["c"]
	fmt.Println(v)
	// youMap := map[string]int{}
	// 这样会报语法错误
	// if youMap == myMap
	//for range myMap {
	//	myMap["d"] = 2
	//}
	fmt.Println(myMap)
}

func testCopy() {
	//创建切片
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr3 := make([]int, 2)

	// arr3 = [1,2]
	copy(arr3, arr1)
	fmt.Println(arr3)
}

func testConst() {
	const N = 100
	var x int = N

	const M int32 = 100
	// 下面这个会异常
	// var y int = M
	fmt.Println(x)
}

type T struct{}

func (t T) f(n int) T {
	fmt.Print(n)
	return t
}

func testPrint() {
	var t T
	// 不明白为什么打印132
	defer t.f(1).f(2)
	fmt.Print(3)
}

func testPrintRange() {
	strs := []string{"one", "two", "three"}

	for _, s := range strs {
		go func() {
			time.Sleep(1 * time.Second)
			// range先跑完 获取到的都是最后一个item的地址
			fmt.Printf("%s ", s)
		}()
	}
	time.Sleep(3 * time.Second)
}

func testPrintInCre() {
	v := 1
	incr(&v)
	fmt.Println(v)
}

func incr(p *int) int {
	*p++
	return *p
}
