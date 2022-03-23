func min(a,b int) int {
    if a <b {
        return a
    }
    return b
}
func getSmallestString(n int, k int) string {
    letter := make([]rune,n,n)
    k-=n
    for i:=n-1;i>=0;i--{
        val := min(25,k)
        letter[i] = rune(97+val)
        k -= val
    }
    
    return string(letter)
}
