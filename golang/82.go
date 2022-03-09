/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var pHead *ListNode = &ListNode{-1000, nil}
    
    pCur := pHead
    
    for {
        if head == nil {
            break
        }
        
        duplicate := false
        for {
            if head.Next == nil || head.Val != head.Next.Val{
                break
            }
            head = head.Next
            duplicate = true
        }
        
        if !duplicate {
            pCur.Next = head
            pCur = pCur.Next
        }
        
        head = head.Next
    }
    pCur.Next = nil
    return pHead.Next
}
