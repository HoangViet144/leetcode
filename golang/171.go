func titleToNumber(columnTitle string) int {
    ans:=0
    base:=1
    n:=len(columnTitle)
    for i:=n-1;i>=0;i--{
        ans+=base*(int(columnTitle[i]-'A')+1)
        base*=26
    }
    return ans
}
