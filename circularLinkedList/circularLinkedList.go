package circularLinkedList

import (
	"fmt"
	"org/tik/hyper-queue-service/utils"
)

/*
Operations on the circular linked list:
Insertion
Deletion

1. Insertion in the circular linked list:
A node can be added in three ways:
Insertion at the beginning of the list
Insertion at the end of the list
Insertion in between the nodes
2. Deletion in a circular linked list:
1) Delete the node only if it is the only node in the circular linked list:

Free the node’s memory
The last value should be NULL A node always points to another node, so NULL assignment is not necessary.
Any node can be set as the starting point.
Nodes are traversed quickly from the first to the last.
2) Deletion of the last node:

Locate the node before the last node (let it be temp)
Keep the address of the node next to the last node in temp
Delete the last memory
Put temp at the end
3) Delete any node from the circular linked list:
We will be given a node and our task is to delete that node from the circular linked list.
*/

type Node[T any] struct {
	value T
	next  *Node[T]
}

type CircularLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (this *CircularLinkedList[T]) Insert(value T) {
	newNode := &Node[T]{value: value, next: nil}
	if this.head == nil {
		this.head = newNode
	} else {
		this.tail.next = newNode
	}
	this.tail = newNode
	this.tail.next = this.head
}

func (this *CircularLinkedList[T]) Delete(value T) {
	current := this.head
	if this.head == nil {
		return
	}
	for ok := true; ok; ok = !(current == this.head) {
		next := current.next
		if utils.IsEqual(next.value, value) {
			if this.tail == this.head {
				this.tail = nil
				this.head = nil
			} else {
				current.next = next.next
				if this.head == next { // deleteing the head
					this.head = this.head.next
				}
				if this.tail == next { // deleting the tail
					this.tail = current
				}
			}
			break
		}
		current = next
	}
}

func (this *CircularLinkedList[T]) Poll() T {
	// var currentValue interface{}
	var value T
	if this.head != nil {
		value = this.head.value
	}
	if this.tail == this.head {
		this.tail = nil
		this.head = nil
	} else {
		this.head = this.head.next
		this.tail.next = this.head
	}
	return value
}

func (this *CircularLinkedList[T]) Show() {
	p := this.head
	if p == nil {
		fmt.Println("Empty list")
		return
	}
	for ok := true; ok; ok = !(p == this.head) {
		fmt.Printf("-> %v", p.value)
		p = p.next
	}
	fmt.Println()
}

func (this *CircularLinkedList[T]) ToArray() []T {
	var rs []T
	p := this.head
	if p == nil {
		fmt.Println("Empty list")
		return rs
	}
	for ok := true; ok; ok = !(p == this.head) {
		fmt.Printf("-> %v", &p.value)
		rs = append(rs, p.value)
		p = p.next
	}
	fmt.Println()
	return rs
}

func (this *CircularLinkedList[T]) ContainsNode(value T) bool {
	n := this.head
	if this.head == nil {
		return false
	} else {
		for ok := true; ok; ok = !(n == this.head) {
			if utils.IsEqual(n.value, value) {
				return true
			}
			n = n.next
		}
		return false
	}
}

// func main() {
// 	sl := CircularLinkedList{}
// 	for i := 0; i < 3; i++ {
// 		sl.insert(rand.Intn(100))
// 	}
// 	sl.show()
// 	fmt.Println(sl.poll())
// 	fmt.Println(sl.poll())
// 	fmt.Println(sl.poll())
// 	fmt.Println(sl.poll())
// 	sl.insert(15)
// 	// fmt.Println(sl.poll())
// 	sl.show()
// }
