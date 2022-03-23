func longestCommonPrefix(strs []string) string {
    prefix := ""
    for _, val := range strs[0]{
        prefix += string(val)
        check := true
        for _, s := range strs {
            if !strings.HasPrefix(s, prefix){
                check = false
                break
            }
        }
        if !check{
            return prefix[:len(prefix)-1]
        }
    }
    return prefix
}
