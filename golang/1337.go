type Soldier struct {
    cnt, ind int
}
func kWeakestRows(mat [][]int, k int) []int {
    ar := make([]Soldier, 0, 0)
    for ind, row := range mat {
        cnt := 0
        for _, val := range row {
            cnt += val
        }
        ar = append(ar, Soldier{cnt, ind})
    }
    sort.Slice(ar, func(i,j int)bool{
        if ar[i].cnt == ar[j].cnt {
            return ar[i].ind < ar[j].ind
        }
        return ar[i].cnt < ar[j].cnt
    })
    ans := make([]int, 0, 0)
    for i:=0;i<k;i++{
        ans = append(ans, ar[i].ind)
    }
    return ans
}
