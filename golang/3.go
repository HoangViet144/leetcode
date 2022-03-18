func max(a,b int) int {
    if a<b {
        return b
    }
    return a
}
func lengthOfLongestSubstring(s string) int {
    letter := make([]int, 256, 256)
    left := 0
    right:=0
    ans := 0
    
    arChar := []rune(s)
    for ;right<len(arChar); {
        letter[arChar[right]]++
        for ;letter[arChar[right]]>1;{
            letter[arChar[left]]--
            left++
        }
        ans = max(ans, right-left+1)
        right++
    }
    return ans
}
