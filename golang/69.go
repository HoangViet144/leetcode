func mySqrt(x int) int {
    left :=0
    right :=x
    ans:= 0
    for ;left<=right ;{
        mid := (left+right)/2
        if mid*mid <= x{
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}
