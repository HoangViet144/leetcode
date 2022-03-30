func searchMatrix(matrix [][]int, target int) bool {
    row := len(matrix)
    col := len(matrix[0])
    start := 0
    end := row*col - 1
    for ;start <= end; {
        mid := start + (end-start)/2
        i := mid/col
        j := mid%col
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] > target {
            end = mid-1
        } else {
            start = mid + 1
        }
    }
    return false
}
