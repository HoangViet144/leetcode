func naturalPow(x float64, n int) float64{
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x
    }
    
    tmp := naturalPow(x, n/2)
    tmp = tmp*tmp
    
    if n&1 == 1 {
        return tmp*x
    } else {
        return tmp
    }
}

func AbsInt(n int) int{
    if n < 0 {
        return -n
    }
    return n
}

func myPow(x float64, n int) float64 {
    if x == 0 {
        return 0
    }
    if n == 0 {
        return 1
    }
    isNegativeX := x < 0
    isNegativeN := n < 0
    
    ans := naturalPow(math.Abs(x), AbsInt(n))
    if isNegativeX && (AbsInt(n) & 1 == 1){
        ans = - ans
    }
    
    if isNegativeN {
        ans = 1/ans
    }
    
    return ans
}
