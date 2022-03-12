func champagneTower(poured int, query_row int, query_glass int) float64 {
    MAX_LEVEL := 100
    glasses := make([][]float64, MAX_LEVEL+1, MAX_LEVEL+1)
    for i := range glasses {
        glasses[i] = make([]float64, MAX_LEVEL+1, MAX_LEVEL+1)
    }
    
    glasses[0][0] = float64(poured)
    
    for r := 0; r <= query_row; r++{
        for c := 0; c <= r; c++{
            amount := (glasses[r][c]-1)/2
            if amount > 0 {
                glasses[r+1][c] += amount
                glasses[r+1][c+1] += amount
            }
        }
    }
    
    return math.Min(1.0,glasses[query_row][query_glass])
}
