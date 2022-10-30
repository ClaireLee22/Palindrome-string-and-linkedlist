package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  rune
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	reversedSecondHalfNode := reverseLinkedList(slow)
	firstHalfNode := head

	for reversedSecondHalfNode != nil {
		if firstHalfNode.Val != reversedSecondHalfNode.Val {
			return false
		}
		reversedSecondHalfNode = reversedSecondHalfNode.Next
		firstHalfNode = firstHalfNode.Next
	}
	return true
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	node1 := &ListNode{Val: 'a'}
	node2 := &ListNode{Val: 'b'}
	node3 := &ListNode{Val: 'c'}
	node4 := &ListNode{Val: 'c'}
	node5 := &ListNode{Val: 'b'}
	node6 := &ListNode{Val: 'a'}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	head := node1
	fmt.Println("Is palindrome?", isPalindrome(head))
}

// output: Is palindrome? true
