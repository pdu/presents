package main

import "fmt"

// START1 OMIT
type node struct {
	val  int
	next *node
}

// END1 OMIT

// START6 OMIT
func insert(val int, head, tail *node) (*node, *node) {
	if val <= head.val {
		head = &node{
			val:  val,
			next: head,
		}
		return head, tail
	}
	for iter := head; iter.next != nil; iter = iter.next {
		if val <= iter.next.val {
			iter.next = &node{
				val:  val,
				next: iter.next,
			}
			return head, tail
		}
	}
	tail.next = &node{
		val:  val,
		next: nil,
	}
	tail = tail.next
	return head, tail
}

// END6 OMIT

// START8 OMIT
func remove(val int, head, tail *node) (*node, *node) {
	if val == head.val {
		if head == tail {
			tail = head.next
		}
		head = head.next
		return head, tail
	}
	for iter := head; iter.next != nil; iter = iter.next {
		if iter.next.val == val {
			iter.next = iter.next.next
			if iter.next == nil {
				tail = iter
			}
			return head, tail
		}
	}
	return head, tail
}

// END8 OMIT

func main() {
	// START2 OMIT
	head := &node{
		val:  1,
		next: nil,
	}
	tail := head
	// END2 OMIT
	// START3 OMIT
	for i := 3; i <= 20; i += 2 {
		tail.next = &node{
			val:  i,
			next: nil,
		}
		tail = tail.next
	}
	// END3 OMIT
	// START4 OMIT
	for iter := head; iter != nil; iter = iter.next {
		fmt.Println("val: ", iter.val)
	}
	// END4 OMIT
	// START5 OMIT
	for i := 0; i <= 20; i += 2 {
		head, tail = insert(i, head, tail)
	}
	for iter := head; iter != nil; iter = iter.next {
		fmt.Println("val: ", iter.val)
	}
	fmt.Println("head val: ", head.val)
	fmt.Println("tail val: ", tail.val)
	// END5 OMIT
	// START7 OMIT
	head, tail = remove(0, head, tail)
	fmt.Println("head val: ", head.val)
	fmt.Println("tail val: ", tail.val)
	head, tail = remove(20, head, tail)
	fmt.Println("head val: ", head.val)
	fmt.Println("tail val: ", tail.val)
	for i := 0; i < 20; i++ {
		head, tail = remove(i, head, tail)
	}
	for iter := head; iter != nil; iter = iter.next {
		fmt.Println("val: ", iter.val)
	}
	// END7 OMIT
}
