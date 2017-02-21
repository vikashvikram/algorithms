package algorithms

type ListNode struct {
	Val int
	Next *ListNode
}

func prepareLinkedList(arr []int) *ListNode  {
	if len(arr) == 0 {
		return nil
	}

	result := []*ListNode{}

	for _, v := range arr {
		result = append(result, &ListNode{v, nil})
	}

	for i:= 0; i < len(arr) - 1; i++ {
		result[i].Next = result[i+1]
	}

	return result[0]
}

func ShowAsArray(node *ListNode) []int  {
	if node == nil {
		return []int{}
	}

	result := []int{}
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	result := []int{}

	mod := 0

	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + mod
		result = append(result, val % 10)
		mod = val / 10
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val := l1.Val + mod
		result = append(result, val % 10)
		mod = val / 10
		l1 = l1.Next
	}

	for l2 != nil {
		val := l2.Val + mod
		result = append(result, val % 10)
		mod = val / 10
		l2 = l2.Next
	}

	if mod != 0 {
		result = append(result, mod)
	}

	return prepareLinkedList(result)

}
