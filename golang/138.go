/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    mapAddr := make(map[*Node]*Node)
    var copyHead *Node = new(Node)
    
    var pCur *Node = head
    var pCurCopy *Node = copyHead
    for {
        if pCur == nil {
            break
        }
        pCurCopy.Next = &Node{pCur.Val, nil, nil}
        pCurCopy = pCurCopy.Next
        mapAddr[pCur] = pCurCopy
        pCur = pCur.Next
    }
    
    pCur = head
    pCurCopy = copyHead.Next
    for {
        if pCur == nil {
            break
        }
        pCurCopy.Random,_ = mapAddr[pCur.Random]
        pCur = pCur.Next
        pCurCopy = pCurCopy.Next
    }
    
    return copyHead.Next
}
