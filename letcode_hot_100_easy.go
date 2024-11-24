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
	// 注意这个条件是: tmp.Next
	// 注意这个条件是: tmp.Next
	// 注意这个条件是: tmp.Next
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

// isPalindromeErr 检查是否是回文链表
// 这个是错误的,因为回文链表不代表是对称的
// 比如 [1]是回文 [1,0,1]也是回文
func isPalindromeErr(head *ListNode) bool {
	var stack []int
	for tmp := head; tmp != nil; tmp = tmp.Next {
		if len(stack) == 0 {
			stack = append(stack, tmp.Val)
		} else {
			if stack[len(stack)-1] == tmp.Val {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, tmp.Val)
			}
		}
	}
	return len(stack) == 0
}

// isPalindrome 双指针(遍历前半部分,在数组的对称位置有相同的数值表示这是一个回文)
func isPalindrome(head *ListNode) bool {
	var arr []int
	for tmp := head; tmp != nil; tmp = tmp.Next {
		arr = append(arr, tmp.Val)
	}
	for i := 0; i < (len(arr) / 2); i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}

// hasCycle 链表是否是环形
func hasCycle(head *ListNode) bool {
	// 重点是这个Node地址,一旦在后面遍历链表的时候出现一样的地址表示环形
	mapper := map[*ListNode]bool{}
	for tmp := head; tmp != nil; tmp = tmp.Next {
		if _, exist := mapper[tmp]; exist {
			return true
		} else {
			mapper[tmp] = true
		}
	}
	return false
}

// singleNumber 数组之中找出只出现了一次的数据
// 按位异或-主要特性是两个相同的会抵消为0
// 同样二进制最低位有5个数字 0 0 1 1 1
// 最终就会变成1 不管怎么运算都会变成找出 奇数的那个位 0/1
// 1^2^3^4^1^2^3 = 4
func singleNumber(nums []int) int {
	// 这个初始值的设定是很关键的
	// 按位异或运算与 0 异或的特性
	// 后续任何数与 0 进行按位异或运算都能得到该数本身
	single := 0
	for _, num := range nums {
		// single 初始值为 0，根据与 0 异或特性，此时 single 变为 num
		// 当第二次遇到 num 时，根据自反性特性，num ^ num = 0，所以 single 又变回了 0
		single = single ^ num
	}
	return single
}

// majorityElement 找出数组之中多数的元素
func majorityElement(nums []int) int {
	mapper := map[int]int{}
	for _, v := range nums {
		if tmp, exist := mapper[v]; exist {
			mapper[v] = tmp + 1
		} else {
			mapper[v] = 1
		}
	}
	l := len(nums) / 2
	for num, count := range mapper {
		if count > l {
			return num
		}
	}
	return 0
}

// climbStairs 爬楼梯
// f(0) = 1
// f(1) = 1
// f(2) = 2
// f(3) = 3
// f(4) = 5
// ...
// f(x) = f(x-1) + f(x-2) // 5 = 3 + 2
func climbStairs(n int) int {
	lastOfLast, last, result := 0, 0, 1
	for i := 1; i <= n; i++ {
		// 上一步f(x-1) 已经不再是上一步了, 而是上两步 f(x-2)
		// 因为index向前推进了
		lastOfLast = last
		// 之前的结果 f(x)也就是 result
		// 变成上一步 f(x-1)
		last = result
		// 将 f(x-2) + f(x-1)
		result = lastOfLast + last
	}
	return result
}

// maxProfit 获取买卖股票的最佳时机
// 这里主要是会出现运算时间太长
func maxProfit(prices []int) int {
	maxNum := 0
	for index, value := range prices {
		if index == len(prices)-1 {
			break
		}
		for i := index + 1; i < len(prices); i++ {
			tmp := value - prices[i]
			if tmp > maxNum {
				maxNum = tmp
			}
		}
	}
	return maxNum
}

// maxProfitSuccess 计算最大利润
// 想要少循环就把最小值和最大利润都记录,在遍历的时候记录最小和最大
func maxProfitSuccess(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		// 更新最低价格
		if price < minPrice {
			minPrice = price
		} else {
			// 计算当前价格与最低价格的差值，并更新最大利润
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

// searchInsert 搜索插入位置
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	targetIndex := 0
	for index, value := range nums {
		if value == target {
			return index
		} else {
			if value > target {
				targetIndex = index
				break
			} else {
				// 如果小于这个targetIndex也需要往后挪动
				targetIndex = index + 1
			}
		}
	}

	return targetIndex
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal 二叉树的中序遍历
// 中序遍历的顺序是先访问左子树，然后访问根节点，最后访问右子树
// 这种遍历方式能够按照节点值的大小顺序（假设二叉树是二叉搜索树）
func inorderTraversal(root *TreeNode) []int {
	// 这里非常巧妙的用到了闭包的函数全局变量的概念
	// result是共用的函数全局变量
	var result []int
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 先递归遍历左子树
		inorder(node.Left)

		// 将当前节点的值添加到结果列表中
		result = append(result, node.Val)

		// 再递归遍历右子树
		inorder(node.Right)
	}

	inorder(root)

	return result
}

// maxDepth 获取树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		// 每次遍历节点不释放,到下一个层级
		// 直到层级节点是nil
		return 0
	}
	// 每次递归都+1深度
	// 并且每次都从左右子树选一个最深的
	return getMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// invertTree 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		// 直到节点为nil才释放
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	// 所谓翻转
	root.Left = right
	root.Right = left
	// 一直子树开始往root走翻转
	return root
}

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

// check 检查二叉树是否对称
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// diameterOfTree 函数用于计算二叉树的直径
func diameterOfTree(root *TreeNode) int {
	var maxDiameter int

	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左子树深度
		leftDepth := maxDepth(node.Left)
		// 递归计算右子树深度
		rightDepth := maxDepth(node.Right)

		// 更新最大直径
		// 这里是最重要的,左子树的深度和右子树的深度,加起来,就是最大直径
		maxDiameter = getMax(maxDiameter, leftDepth+rightDepth)

		// 返回当前节点的最大深度（左右子树深度较大值 + 1）
		return getMax(leftDepth, rightDepth) + 1
	}

	maxDepth(root)

	return maxDiameter
}

// diameterOfBinaryTree 最长直径
func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		lLen := dfs(node.Left) + 1  // 左子树最大链长+1
		rLen := dfs(node.Right) + 1 // 右子树最大链长+1
		ans = max(ans, lLen+rLen)   // 两条链拼成路径
		return max(lLen, rLen)      // 当前子树最大链长
	}
	dfs(root)
	return
}

// 有序数组转二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	// 所有的树操作都是假定最小树怎么写的func
	// 然后递归,直到一个条件的出现停止递归
	if left > right {
		return nil
	}
	// 取下标中间的(int/int会自动向下取整)
	// (0+1)/2 = 0
	// (0+2)/2 = 1
	// (0+3)/2 = 1
	mid := (left + right) / 2

	// 组装root节点
	root := &TreeNode{Val: nums[mid]}
	// 相当于二分法
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

// addTwoNumbers 两数相加 l1 = [2,4,3], l2 = [5,6,4] = [7,0,8]
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n = &ListNode{}
	var Link = n
	carry := 0
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		a := 0
		b := 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + carry
		n.Next = &ListNode{
			Val: sum % 10,
		}
		n = n.Next
		carry = sum / 10
	}
	return Link.Next
}

// longestPalindrome 最长回文子串
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for {
		if right >= len(s) {
			break
		}
		if left < 0 {
			break
		}
		if s[left] != s[right] {
			break
		}
		left, right = left-1, right+1
	}
	return left + 1, right - 1
}

// buildLink 构建链表
func buildLink(list []int) *ListNode {
	// 逆序不断构建Next指向的链表
	var nextNode *ListNode
	// 定义FirstNode
	var firstNode *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		item := &ListNode{
			Val:  list[i],
			Next: nextNode,
		}
		nextNode = item
		firstNode = item
	}
	return firstNode
}
