func increasingTriplet(nums []int) bool {
    minEle := int(1<<31)
    midEle := int(1<<31)
    for _, val:= range nums{
        if minEle >= val {
            minEle = val
        } else if midEle >= val {
            midEle = val
        } else {
            return true
        }
    }
    return false
}
