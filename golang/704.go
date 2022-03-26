func search(nums []int, target int) int {
    L:= 0
    R:= len(nums)-1
    for ;L<=R;{
        mid:=L+(R-L)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            L = mid + 1
        } else {
            R = mid - 1
        }
    }
    return -1
}
