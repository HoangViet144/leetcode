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
func evalRPN(tokens []string) int {
    st := NewStack()
    for _, val := range tokens {
        if val == "+" {
            num1,_ := strconv.Atoi(st.pop().(string))
            num2,_ := strconv.Atoi(st.pop().(string))
            st.push(strconv.Itoa(num1+num2))
        } else if val == "-" {
            num1,_ := strconv.Atoi(st.pop().(string))
            num2,_ := strconv.Atoi(st.pop().(string))
            st.push(strconv.Itoa(num2-num1))
        } else if val == "*" {
            num1,_ := strconv.Atoi(st.pop().(string))
            num2,_ := strconv.Atoi(st.pop().(string))
            st.push(strconv.Itoa(num1*num2))
        } else if val == "/" {
            num1,_ := strconv.Atoi(st.pop().(string))
            num2,_ := strconv.Atoi(st.pop().(string))
            st.push(strconv.Itoa(num2/num1))
        } else {
            st.push(val)
        }
    }
    ans,_ := strconv.Atoi(st.pop().(string))
    return ans
}
