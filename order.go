package p1

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

// 输入[3,1,4,2] left=0 right=3
// 输出[2,1,3,4] return 2
func partition(arr []int, left, right int) int {
	index := left + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[left] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}
	arr[left], arr[index-1] = arr[index-1], arr[left]
	return index - 1
}

// bubbleSort函数实现冒泡排序
func bubbleSort(numbers []int) []int {
	n := len(numbers)
	for index := range numbers {
		if index == n-1 {
			// 最后一个不需要处理了
			continue
		}
		// 每次子循环,都能确定最右边的一定是最大的
		for j := range numbers {
			// 数组之中0的对应的下标就是len-1比如四个元素的数组的0和3
			// -index就是跳过
			if j >= (n-1)-index {
				continue
			}

			if numbers[j] > numbers[j+1] {
				// 将大的排在后面
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}
