func minDominoRotations(tops []int, bottoms []int) int {
    n:= len(tops)
    freq := make([]int, 7, 7)
    for i:=0 ;i<n;i++{
        if tops[i]==bottoms[i] {
            freq[tops[i]]++
        } else {
            freq[tops[i]]++
            freq[bottoms[i]]++
        }
    }
    
    cntMap := make(map[string]int)
    for i:=0;i<n;i++{
        if freq[tops[i]] >=n {
            cntMap[strconv.Itoa(tops[i])+"top"]++
        }
        if freq[bottoms[i]] >=n {
            cntMap[strconv.Itoa(bottoms[i])+"bottoms"]++
        }
    }
    
    topFreq := make([]int,0,0)
    for key:= range cntMap {
        topFreq = append(topFreq, cntMap[key])
    }
    sort.Slice(topFreq, func(i,j int) bool {
        return topFreq[i] > topFreq[j]
    })
    if len(topFreq) == 0 {
        return -1
    }
    return n - topFreq[0]
}
