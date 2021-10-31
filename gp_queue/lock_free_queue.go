package gp_queue

import (
	"sync/atomic"
	"unsafe"
)

type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type node struct {
	value interface{}
	next  unsafe.Pointer
}

func NewNode(v interface{}) *node {
	return &node{
		value: v,
		next:  nil,
	}
}

func NewLKqueue() *LKQueue {
	uPointer := unsafe.Pointer(NewNode(nil))
	return &LKQueue{
		head: uPointer,
		tail: uPointer,
	}
}

func (l *LKQueue) Enqueue(v interface{}) {
	n := NewNode(v)
	tail := load(&l.tail)
	next := load(&tail.next)
	for {
		if tail == load(&l.tail) { //判断 tail 还是原来的tai
			if next == nil { //判断 没有新数据插入队尾
				if cas(&tail.next, next, n) { //向队尾添加数据
					cas(&l.tail, tail, n) //将tail指向队尾
					return
				}
			} else {
				cas(&l.tail, tail, n) //已有数据加入队尾，移动tail至队尾
			}
		}
	}
}

func (l *LKQueue) Dequeue() interface{} {
	head := load(&l.head)
	tail := load(&l.tail)
	next := load(&head.next)
	for {
		if head == load(&l.head) { //判断 head 还是原来的head
			if head == tail { //判断是否为空
				if next == nil { //有数据插入
					return nil
				}
				cas(&l.tail, tail, next) //将tail指向队尾
			} else {
				v := next.value //读取出队数据
				//头指针移动到下一个，如果next的v已经读出去,那么head肯定以及变更，cas失败重新循环
				if cas(&l.head, head, next) {
					return v
				}
			}
		}
	}
}

func load(p *unsafe.Pointer) *node {
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *node) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
