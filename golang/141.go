/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    addrMap := map[*ListNode]bool{}
    for {
        if head == nil{
            return false
        }
        _, exist := addrMap[head]
        if exist {
            return true
        }
        addrMap[head] = true
        head = head.Next
    }
    return false
}
