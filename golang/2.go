/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var ans *ListNode = new(ListNode)
    var pCur *ListNode = ans
    
    var digit1, digit2, carry, remain, sum int
    for {
        if l1 == nil || l2 == nil {
            break
        }
        digit1 = l1.Val
        digit2 = l2.Val
        sum = digit1 + digit2 + carry
        carry = sum / 10
        remain = sum % 10
        
        pCur.Next = &ListNode{remain, nil}
        pCur = pCur.Next
    
        l1 = l1.Next
        l2 = l2.Next
    }
    
    for {
        if l1 == nil {
            break
        }
        sum = l1.Val + carry
        carry = sum / 10
        remain = sum % 10
        
        pCur.Next = &ListNode{remain, nil}
        pCur = pCur.Next
    
        l1 = l1.Next
    }
    
    for {
        if l2 == nil {
            break
        }
        sum = l2.Val + carry
        carry = sum / 10
        remain = sum % 10
        
        pCur.Next = &ListNode{remain, nil}
        pCur = pCur.Next
    
        l2 = l2.Next
    }
    
    if carry > 0 {
        pCur.Next = &ListNode{carry, nil}
    }
    return ans.Next
}
