// ans[i] = ans[i//2] + (i&1==1)
func bool2int(n bool) int{
    if n {
        return 1
    }
    return 0
}
func countBits(n int) []int {
    ans := make([]int, n+1, n+1)
    ans[0] = 0
    for i:=1;i<=n;i++{
        ans[i] = ans[i/2] + bool2int((i&1)==1)
    }
    return ans
}