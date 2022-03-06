func max(a,b int)int {
    if a>b{
        return a
    }
    return b
}
func compareVersion(version1 string, version2 string) int {
    ver1Ar := strings.Split(version1, ".")
    ver2Ar := strings.Split(version2, ".")
    n := len(ver1Ar)
    m := len(ver2Ar)
    for i := 0; i < max(n,m); i++{
        num1 := 0
        num2 := 0
        if i < n{
            num1,_ = strconv.Atoi(ver1Ar[i])
        }
        if i < m{
            num2,_ = strconv.Atoi(ver2Ar[i])
        }
        
        if num1<num2{
            return -1
        }
        
        if num1>num2{
            return 1
        }
    }
    return 0
}
