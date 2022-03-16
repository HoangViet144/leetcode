type node struct {
    value interface{}
    next *node
}	

type Queue struct {
    head *node
    tail *node
    length int
}


func NewQueue() *Queue {
	return &Queue{nil,nil,0}
}

func (q *Queue) size() int {
	return q.length
}

func (q *Queue) top() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.head.value
}

func (q *Queue) deque() interface{} {
	if q.length == 0 {
		return nil
	}
	
	headNode := q.head
	q.head = headNode.next
	q.length--
    if q.length == 0 {
        q.tail = nil
    }
	return headNode.value
}

func (q *Queue) enque(value interface{}) {
    newNode := &node{value, nil}
    if q.head == nil {
        q.head = newNode
        q.tail = newNode
    } else {
        q.tail.next = newNode
        q.tail = q.tail.next
    }
	q.length++
}

func (q *Queue) empty() bool {
    return q.size() == 0
}

type RecentCounter struct {
    ReqLst *Queue
}


func Constructor() RecentCounter {
    recentCnt := RecentCounter{}
    recentCnt.ReqLst = NewQueue()
    return recentCnt
}


func (this *RecentCounter) Ping(t int) int {
    for ; !this.ReqLst.empty() && this.ReqLst.top().(int) < t - 3000 ; {
        this.ReqLst.deque()
    }
    this.ReqLst.enque(t)
    return this.ReqLst.size()
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
