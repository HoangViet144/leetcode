func isSubsequence(s string, t string) bool {
    i := 0
    runeS := []rune(s)
    n := len(runeS)
    for _, c := range t{
        if i >= n{
            break
        }
        if runeS[i] == c{
            i++
        }
    }
    return i==n
}
