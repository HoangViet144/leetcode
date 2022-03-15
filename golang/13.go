func romanToInt(s string) int {
    char2val := make(map[rune]int)
    char2val['I'] = 1
    char2val['V'] = 5
    char2val['X'] = 10
    char2val['L'] = 50
    char2val['C'] = 100
    char2val['D'] = 500
    char2val['M'] = 1000
        
    num := 0;
    runeAr := []rune(s)
    for i := 0; i < len(runeAr); i++ {
        if i+1 <len(runeAr) && char2val[runeAr[i]] < char2val[runeAr[i + 1]] {
            num += char2val[runeAr[i+1]] - char2val[runeAr[i]];
            i++;
        } else {
            num += char2val[runeAr[i]];            
        }
    }

    return num;
}
