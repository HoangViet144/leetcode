type IntStack struct {
    ar []int
}

func (st *IntStack) push (ele int){
    st.ar = append(st.ar, ele)
}

func (st *IntStack) size () int {
    return len(st.ar)
}

func (st *IntStack) pop() int {
    topEle := st.ar[st.size()-1]
    st.ar = st.ar[:st.size()-1]
    return topEle
}
