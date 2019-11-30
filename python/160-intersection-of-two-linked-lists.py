# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        memo = set()
        while headA:
            memo.add(headA)
            headA = headA.next
        while headB:
            if headB in memo:
                return headB
            else:
                memo.add(headB)
                headB = headB.next
        return
