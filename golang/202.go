func sum(n int) int {
    ans := 0
    for ;n>0;{
        ans+= (n%10)*(n%10)
        n/=10
    }
    return ans
}
func isHappy(n int) bool {
    value := make(map[int]bool)
    for {
        n = sum(n)
        if n == 1 {
            return true
        }
        _, isExist := value[n]
        if isExist {
            return false
        }
        value[n]=true
    }
    return false
}
