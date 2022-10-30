# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution(object):
    def isPalindrome(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                break
        
        reversedSecondHalfNode = self.reverseLinkedList(slow)
        firstHalfNode = head
        while reversedSecondHalfNode is not None:
            if firstHalfNode.val != reversedSecondHalfNode.val:
                return False
        
            reversedSecondHalfNode = reversedSecondHalfNode.next
            firstHalfNode = firstHalfNode.next
            
        return True
    
    def reverseLinkedList(self, head):
        prev = None
        current = head
        while current is not None:
            next = current.next
            current.next = prev
            prev = current
            current = next
            
        return prev

if __name__ == "__main__":
    node1 = ListNode("a")
    node2 = ListNode("b")
    node3 = ListNode("c")
    node4 = ListNode("c")
    node5 = ListNode("b")
    node6 = ListNode("a")
    node1.next = node2
    node2.next = node3
    node3.next = node4
    node4.next = node5
    node5.next = node6

    head = node1
    solution = Solution()
    print("Is palindrome?", solution.isPalindrome(head))

# output: ('Is palindrome?', True)