type RuneStack struct {
    container []rune
}

func (st *RuneStack) push(ele rune) {
    st.container = append(st.container, ele)
}

func (st* RuneStack) pop() rune {
    if st.size() == 0 {
        return '.'
    }
    ele := st.container[st.size()-1]
    st.container = st.container[:st.size()-1]
    return ele
}

func (st* RuneStack) size() int {
    return len(st.container)
}

func isValid(s string) bool {
    var st RuneStack
    for _, char := range s {
        if char == '(' || char == '[' || char == '{' {
            st.push(char)
        } else {
            curTop := st.pop()
            fmt.Println(curTop)
            if curTop != '(' && char == ')' {
                return false
            }
            if curTop != '[' && char == ']' {
                return false
            }
            if curTop != '{' && char == '}' {
                return false
            }
        }
    }
    
    return st.size() == 0
}
