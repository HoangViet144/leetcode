func max(a,b int) int{
	if a<b{
			return b
	}
	return a
}

func deleteAndEarn(nums []int) int {
	point := make(map[int]int)
	maxValue := 0
	
	for _, val := range nums{
			point[val] += val
			maxValue = max(val, maxValue)
	}
	
	maxPoint := make([]int, maxValue+1, maxValue+1)
	
	value, exist := point[1]
	if exist{
			maxPoint[1] = value
	}
	
	for i := 2; i <=maxValue; i++ {
			curPoint, exist := point[i]
			if !exist{
					curPoint = 0
			}
			maxPoint[i] = max(maxPoint[i-1], maxPoint[i-2]+curPoint)
	} 
	fmt.Println(point)
	fmt.Println(maxPoint)
	
	return maxPoint[maxValue]
}