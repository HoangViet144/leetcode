func numRescueBoats(people []int, limit int) int {
    sort.Slice(people, func(i,j int) bool{
        return people[i]<people[j]
    })
    i:=0
    j:=len(people)-1
    ans:=0
    for ;i<j; {
        if people[i]+people[j]<=limit {
            i++
            j--
        } else {
            j--
        }
        ans++
    }
    if i==j{
        ans++
    }
    return ans
}
