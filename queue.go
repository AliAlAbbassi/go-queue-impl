package main

import (
	"sync"
	"time"
)

type QueueMessage struct {
	time    int64
	message Message
}

type InnerMessage struct {
	width         int
	girth         int
	height        int
	length        int
	depth         int
	circumference int
}

type Message struct {
	message          string
	another_property InnerMessage
}
type Node struct {
	data QueueMessage
	next *Node
}

type Queue struct {
	length int

	head  *Node
	tail  *Node
	mutex sync.Mutex
}

var m sync.Mutex
var queue *Node

func makeTimestamp() int64 {
	return time.Now().UnixMilli()
}

func NewQueue() *Queue {
	return &Queue{0, nil, nil, sync.Mutex{}}
}

func (q *Queue) enqueue(node *Node) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.length++
	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = q.tail.next
}

func (q *Queue) deque() *Node {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.head == nil {
		return nil
	}
	q.length--
	out := q.head
	q.head = q.head.next
	out.next = nil
	return out
}

func (q *Queue) emptyQueue() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	now := makeTimestamp()
	for {
		node := q.head
		if node != nil && node.data.time < now {
			q.deque()
		} else {
			break
		}
	}
}

func newNode(data QueueMessage) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}
