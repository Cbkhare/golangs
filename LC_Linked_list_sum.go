/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry int = 0
    var head *ListNode = l1
    var old *ListNode
    for (l1!= nil && l2!=nil) {
        var sum int 
        sum = l1.Val+l2.Val + carry
        l1.Val = sum%10
        if (sum>=10) {
            carry = 1
        } else {
            carry = 0
        }
        old =l1
        l1 = l1.Next 
        l2 = l2.Next
    }
    for (l1!=nil){
        l1.Val += carry
        carry = 0
        if (l1.Val==10){
            carry=1
            l1.Val%=10
        }
        if (l1.Next==nil && carry==1) {
            var t *ListNode = new(ListNode)
            t.Val = carry
            t.Next = nil
            l1.Next = t
            carry = 0
        } else if (l1.Next==nil && carry==0) {
            return head
        }
        l1 = l1.Next
    }
    old.Next = l2
    for (l2 != nil){
        l2.Val += carry 
        carry = 0
        if (l2.Val==10){
            carry = 1
            l2.Val %= 10
        }
        if (l2.Next==nil && carry==1) {
            var t *ListNode = new(ListNode)
            t.Val = carry
            t.Next = nil
            l2.Next = t
            carry = 0
        } else if (l2.Next==nil && carry==0) {
            return head
        }
        l2 = l2.Next
    }
    if carry==0 {
        return head
    }
    var t *ListNode = new(ListNode)
    t.Val = carry
    t.Next = nil
    old.Next = t 
    return head 
}
