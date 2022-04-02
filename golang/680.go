func validPalindrome(s string) bool {
    charAr := []rune(s)
    n := len(charAr)
    invalidInd := -1
    for i:=0;i<n/2;i++{
        if charAr[i]!= charAr[n-1-i] {
            invalidInd = i
            break
        }
    }
    if invalidInd != -1 {
        return isPalindrome(charAr[invalidInd:n-1-invalidInd]) ||
        isPalindrome(charAr[invalidInd+1: n-1-invalidInd+1])
    }
    return true
}

func isPalindrome(charAr []rune)bool{
    n := len(charAr)
    for i:=0;i<n/2;i++{
        if charAr[i] != charAr[n-1-i] {
            return false
        }
    }
    return true
}
