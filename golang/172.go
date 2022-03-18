func trailingZeroes(n int) int {
    ans := 0
    factor5 := 5
    for {
        inc := n/factor5
        if inc == 0 {
            break
        }
        factor5*=5
        ans += inc
    }
    return ans
}
