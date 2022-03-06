func majorityElement(nums []int) int {
    sort.Slice(nums, func (i, j int) bool{
        return nums[i] < nums[j]
    })
    n:=len(nums)
    return nums[n/2]
}
