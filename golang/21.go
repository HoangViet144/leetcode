/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var nodeHead *ListNode = new(ListNode)
    var pCur *ListNode = nodeHead
    
    for {
        if list1 == nil || list2 == nil{
            break
        }
        if list1.Val < list2.Val {
            pCur.Next = new(ListNode)
            pCur = pCur.Next
            pCur.Val = list1.Val
            list1 = list1.Next
        } else{
            pCur.Next = new(ListNode)
            pCur = pCur.Next
            pCur.Val = list2.Val
            list2 = list2.Next
        }
    }
    
    for ;list1 != nil; {
        pCur.Next = new(ListNode)
        pCur = pCur.Next
        pCur.Val = list1.Val
        list1 = list1.Next
    }
    
    for ;list2 != nil; {
        pCur.Next = new(ListNode)
        pCur = pCur.Next
        pCur.Val = list2.Val
        list2 = list2.Next
    }
    return nodeHead.Next
}
