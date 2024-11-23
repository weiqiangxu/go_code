package p1

// twoSum 两数之和 - 通过mapper[target - value]
func twoSum(nums []int, target int) []int {
	var mapper = map[int]int{}
	var targetIndexList []int
	for index, value := range nums {
		if rightIdx, exist := mapper[target-value]; exist {
			targetIndexList = append(targetIndexList, index, rightIdx)
		} else {
			mapper[value] = index
		}
	}
	return targetIndexList
}

// isValid 栈-验证括号是否有效
func isValid(s string) bool {
	arr := []rune(s)
	// 用栈计数字
	var stack []rune
	mapper := map[rune]rune{}
	// runeLeft 的参数格式是 int32
	// runeLeft := '{'
	mapper['{'] = '}'
	mapper['['] = ']'
	mapper['('] = ')'
	for _, item := range arr {
		if len(stack) == 0 {
			// 长度为0也入栈
			stack = append(stack, item)
		} else {
			left := stack[len(stack)-1]
			if right, ok := mapper[left]; ok {
				if item == right {
					// 如果相等就出栈
					stack = stack[:len(stack)-1]
				} else {
					// 不相等就入栈
					stack = append(stack, item)
				}
			} else {
				// 字符异常也入栈
				stack = append(stack, item)
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

// moveZeroes 不改变原来数组的情况下，将0移动到后面，非0元素保持原来的位置
func moveZeroes(nums []int) {
	for i, value := range nums {
		if value == 0 {
			for rightIndex := (i + 1); rightIndex < len(nums); rightIndex++ {
				if nums[rightIndex] != 0 {
					nums[i], nums[rightIndex] = nums[rightIndex], nums[i]
					break
				}
			}
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 获取两个链表相交的点
// 注意这个链表是保证不会有环也就是说不会有两个链表节点相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 这里是最精彩的
	// 用的是链表节点的地址,注意是地址
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

// reverseList 反转链表-执行时间超时
func reverseList(head *ListNode) *ListNode {
	var preNode *ListNode
	var curr *ListNode
	for curr = head; curr != nil; {
		// 记录下一个节点地址
		next := curr.Next
		// 当前节点的下一个节点地址替换成上一个节点
		curr.Next = preNode
		// 更新上一个节点为当前节点地址
		preNode = curr
		// 更新当前节点为下一个节点-继续遍历
		curr = next
	}
	return preNode
}

// reverseList 反转链表-执行时间超时
func reverseListSuccess(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		// 下一个节点
		next := curr.Next
		// 当前节点的Next指向上一个节点
		curr.Next = prev
		// 上一个节点指向当前节点
		prev = curr
		// 当前节点换成下一个节点
		curr = next
	}
	return prev
}
