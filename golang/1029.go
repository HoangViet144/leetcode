func twoCitySchedCost(costs [][]int) int {
    ans := 0
    for _, cost := range costs{
        ans+= cost[0]
    }
    n:= len(costs)
    diff := make([]int, n, n)
    for i:=0;i<n;i++ {
        diff[i]= costs[i][1]-costs[i][0]
    }
    sort.Slice(diff,  func(i,j int) bool{
        return diff[i] < diff[j]
    })
    for i:=0;i<n/2;i++{
        ans+= diff[i]
    }
    return ans
}
