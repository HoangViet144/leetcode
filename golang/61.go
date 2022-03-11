/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    pCur := head
    n := 0
    for {
        if pCur == nil {
            break
        }
        
        n += 1
        pCur = pCur.Next
    }
    
    if n == 0 {
        return head
    }
    
    k %= n
    k = n - k
    
    if k == n {
        return head
    }
    
    pCur = head
    var pNext *ListNode = nil

    fmt.Println(k,n)
    for {
        if k == 1 {
            pNext = pCur.Next
            pCur.Next = nil
            break
        }
        pCur = pCur.Next
        k--
    }
    
    pCur = pNext
    
    for {
        if pCur.Next == nil {
            break
        }
        pCur = pCur.Next
    }
    
    pCur.Next = head
    
    return pNext
}
