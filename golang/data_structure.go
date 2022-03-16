type node struct {
    value interface{}
    prev *node
}	

type Stack struct {
    head *node
    length int
}


func NewStack() *Stack {
	return &Stack{nil,0}
}

func (st *Stack) size() int {
	return st.length
}

func (st *Stack) top() interface{} {
	if st.length == 0 {
		return nil
	}
	return st.head.value
}

func (st *Stack) pop() interface{} {
	if st.length == 0 {
		return nil
	}
	
	headNode := st.head
	st.head = headNode.prev
	st.length--
	return headNode.value
}

func (st *Stack) push(value interface{}) {
	st.head = &node{value,st.head}
	st.length++
}

\\-------------------------------------------------\\

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

