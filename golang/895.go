type FreqStack struct {
    freq map[int]int
    maxFreq int
    elesByFreq map[int][]int
}


func Constructor() FreqStack {
    return FreqStack{make(map[int]int), 0, make(map[int][]int)}
}


func (this *FreqStack) Push(val int)  {
    _, exist := this.freq[val]
    curFreq := 1
    if exist {
        curFreq = this.freq[val] + 1
    }
    this.freq[val] = curFreq
    if curFreq > this.maxFreq {
        this.maxFreq = curFreq
    }
    
    _, exist = this.elesByFreq[curFreq]
    if !exist {
        this.elesByFreq[curFreq] = make([]int, 0, 0)
    }
    this.elesByFreq[curFreq] = append(this.elesByFreq[curFreq], val)
}


func (this *FreqStack) Pop() int {
    n := len(this.elesByFreq[this.maxFreq])
    maxFreqEle := this.elesByFreq[this.maxFreq][n-1]
    this.elesByFreq[this.maxFreq] = this.elesByFreq[this.maxFreq][:n-1]
    
    this.freq[maxFreqEle]--
    
    if len(this.elesByFreq[this.maxFreq]) == 0 {
        this.maxFreq -- 
    }
    return maxFreqEle
    return 0
}


/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
