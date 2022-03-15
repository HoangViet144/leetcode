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

func minRemoveToMakeValid(s string) string {
    runeAr := []rune(s)
    var openBrSt IntStack
    for index, ele := range runeAr{
        if ele == '(' {
            openBrSt.push(index)
        } else if ele==')'{
            if openBrSt.size() == 0 {
                runeAr[index] = '#'
            } else {
                openBrSt.pop()
            }
        }
    }
    
    for ;openBrSt.size()>0; {
        ind := openBrSt.pop()
        runeAr[ind] = '#'
    }
    
    ans := ""
    for _, ele := range runeAr {
        if ele != '#' {
            ans += string(ele)
        }
    }
    return ans
}
