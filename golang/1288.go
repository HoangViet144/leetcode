type Pair struct {
    first, second int
}

func removeCoveredIntervals(intervals [][]int) int {
    n := len(intervals)
    bound := make([]Pair,0, n*2)
    
    for _, value := range intervals{
        bound = append(bound, Pair{value[0], value[1]})
    }
    
    sort.Slice(bound, func(i, j int) bool {
        if bound[i].first == bound[j].first {
            return bound[i].second > bound[j].second
        }
        return bound[i].first < bound[j].first
    })
    
    curMax := bound[0].second
    
    for i:= 1; i < len(bound); i++{
        if curMax >= bound[i].second {
            n -= 1
        } else{
            curMax = bound[i].second
        }
    }
    
    return n
}
