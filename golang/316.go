func removeDuplicateLetters(s string) string {
    letter := make([]int, 26,26)
    for _, val := range s {
        letter[val-'a']++
    }
    visited := make([]bool, 26, 26)
    var ans []rune
    for _, val := range s{
        letter[val-'a']--
        if !visited[val-'a'] {
            for ;len(ans)>0 && ans[len(ans)-1] > val && letter[ans[len(ans)-1]-'a']>0; {
                visited[ans[len(ans)-1]-'a']= false
                ans[len(ans)-1] = 0
                ans = ans[:len(ans)-1]
            }
            visited[val-'a']=true
            ans = append(ans, val)
        }
    }
    return string(ans)
}
