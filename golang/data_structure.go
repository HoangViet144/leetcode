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

