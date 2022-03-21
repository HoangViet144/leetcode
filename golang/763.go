func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}
func partitionLabels(s string) []int {
    letterInd := make([]int, 26, 26)
    for i:=0;i<26;i++{
        letterInd[i]= -1
    }
    for ind, c := range s {
        letterInd[c-'a']=ind
    }
    fmt.Println(letterInd)
    startInd := 0
    endInd := 0
    ans := make([]int,0,0)
    for ind, val := range s{
        endInd = max(endInd, letterInd[val-'a'])
        if endInd <= ind {
            ans = append(ans, endInd- startInd+1)
            startInd = ind+1
            endInd = ind+1
        }
    }
    return ans
}
