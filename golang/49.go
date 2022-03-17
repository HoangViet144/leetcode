type anagram struct {
    normalized, origin string
}

func sortString(str string) string {
    runeAr := []rune(str)
    sort.Slice(runeAr, func(i, j int)bool {
        return runeAr[i] < runeAr[j]
    })
    return string(runeAr)
}
func groupAnagrams(strs []string) [][]string {
    anaLst := make([]anagram, len(strs), len(strs))
    for ind, str := range strs {
        anaLst[ind] = anagram{sortString(str), str}
    }
    sort.Slice(anaLst, func(i, j int)bool {
        return anaLst[i].normalized < anaLst[j].normalized
    })
    fmt.Println(anaLst)
    
    curAnagram := "#"
    ans := make([][]string,0,0)
    for _, anagram := range anaLst {
        if  curAnagram != anagram.normalized {
            ans = append(ans,make([]string,0,0))
            ans[len(ans)-1] = append(ans[len(ans)-1], anagram.origin)
            curAnagram = anagram.normalized
        } else {
            ans[len(ans)-1] = append(ans[len(ans)-1], anagram.origin)
        }
    }

    
    return ans
}
