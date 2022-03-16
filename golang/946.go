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

func validateStackSequences(pushed []int, popped []int) bool {
    st := NewStack()
    ind := 0
    sz := len(popped)
    for _, value := range pushed {
        st.push(value)
        for ;st.size()>0 && ind < sz && st.top().(int) == popped[ind];{
            st.pop()
            ind++
        }
    }
    return ind == sz
}
