type StrStack struct {
    ar []string
}

func (st *StrStack) size() int {
    return len(st.ar)
}

func (st *StrStack) push(ele string){
    st.ar = append(st.ar, ele)
}

func (st *StrStack) pop() string {
    if st.size() == 0 {
        return "nil"
    }
    
    top := st.ar[st.size()-1]
    st.ar = st.ar[:st.size()-1]
    return top
}

func simplifyPath(path string) string {
    var dirSt StrStack
    ans := ""
    dir := ""
    
    runeAr := []rune(path)
    n := len(runeAr)
    
    for i := 0; i < n; i++ {
        if string(runeAr[i]) == "/" {
            continue
        }
        
        dir = ""
        for ; i < n && string(runeAr[i]) != "/"; i++ {
            dir += string(runeAr[i])
        }
        
        if dir=="." {
            continue 
        } else if dir == ".." {
            dirSt.pop();
        } else {
            dirSt.push(dir)
        }
    }
    
    for ; dirSt.size() > 0 ; {
        ans = "/" + dirSt.pop() + ans
    }
    
    if ans == "" {
        ans = "/"
    }
    return ans
}
