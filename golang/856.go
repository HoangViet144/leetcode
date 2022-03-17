type node struct {
    value interface{}
    prev *node
}	

type Stack struct {
    head *node
    length int
}


func NewStack() *Stack {
	return &Stack{nil,0}
}

func (st *Stack) size() int {
	return st.length
}

func (st *Stack) top() interface{} {
	if st.length == 0 {
		return nil
	}
	return st.head.value
}

func (st *Stack) pop() interface{} {
	if st.length == 0 {
		return nil
	}
	
	headNode := st.head
	st.head = headNode.prev
	st.length--
	return headNode.value
}

func (st *Stack) push(value interface{}) {
	st.head = &node{value,st.head}
	st.length++
}

func intMax(a,b int) int {
    if a<b {
        return b
    }
    return a
}
func scoreOfParentheses(s string) int {
    st := NewStack()
    st.push(0)
    for _,value := range(s) {
        if value=='(' {
            st.push(0)
        } else {
            topEle := st.pop().(int)
            st.push(st.pop().(int)+intMax(topEle*2, 1))
        }
    }
    return st.pop().(int)
}
