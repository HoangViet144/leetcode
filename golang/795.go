func numSubarrayBoundedMax(nums []int, left int, right int) int {
    leftBound := 0
    rightBound := 0
    ans := 0
    cnt := 0
    for {
        if rightBound >= len(nums){
            break
        }
        if left <= nums[rightBound] && nums[rightBound] <= right {
            cnt = rightBound - leftBound + 1
            ans += cnt
        } else if nums[rightBound] < left {
            ans += cnt
        } else {
            leftBound = rightBound + 1
            cnt = 0
        }
        rightBound += 1
    }
    return ans
}
