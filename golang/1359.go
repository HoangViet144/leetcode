func countOrders(n int) int {
    if n == 1{
        return 1
    }
    
    var incAmount,i,m int
    ans := 1
    mod := int(1e9 + 7)
    ans += mod
    
    for i= 2;i<=n;i++{
        m = (i-1)*2+1
        incAmount = m*(m+1)/2
        ans = (ans * (incAmount%mod))%mod
    }
    return ans
}