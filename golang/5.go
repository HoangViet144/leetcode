func longestPalindrome(s string) string {
    n := len(s)
    isPalindrome := make([][]bool, n+1, n+1)
    fmt.Println(len(isPalindrome))
    for i:=0;i<n+1;i++{
        isPalindrome[i] = make([]bool, n+1, n+1)
    }
    for i:=0;i<n;i++{
        isPalindrome[i][i] = true
        isPalindrome[i+1][i] = true
    }
    
    curLen:=1
    left:=0
    right:=1
    for sz:=2;sz<=n;sz++ {
        for i:=0;i<n-sz+1;i++ {
            j:= i + sz -1
            isPalindrome[i][j] = isPalindrome[i+1][j-1] && (s[i]==s[j])
            // fmt.Println(i+1,j-1,i, j, isPalindrome[i][j])

            if isPalindrome[i][j] && curLen < j-i+1{
                left = i
                right = j+1
                curLen = j-i+1
            }
        }
    }
    fmt.Println(left,right)
    return s[left:right]
}
