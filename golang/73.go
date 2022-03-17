func setZeroes(matrix [][]int)  {
    R := len(matrix)
    C := len(matrix[0])
    
    firstCol := false
    for i:=0; i<R; i++ {
        if matrix[i][0] == 0 {
            firstCol = true
        }
        for j:=1; j<C; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    
    for i:=1; i<R; i++ {
        for j:=1; j<C; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    
    if matrix[0][0] == 0 {
        for i:=0;i<C;i++{
            matrix[0][i] = 0
        }
    }
    if firstCol {
        for i:=0;i<R;i++{
            matrix[i][0] = 0
        }
    }
}
