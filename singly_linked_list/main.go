package main

import "fmt"

type Node struct {
	data      int
	next_node *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) append(data int) {
	new_node := &Node{data: data, next_node: nil}

	if l.head == nil {
		l.head = new_node
		return
	}

	last_node := l.head
	for last_node.next_node != nil {
		last_node = last_node.next_node
	}
	last_node.next_node = new_node
}

func (l *LinkedList) insert(data int) {
	new_node := &Node{data: data, next_node: nil}
	// To Do : Why infinite loop occurs when the below two code are switched.
	new_node.next_node = l.head
	l.head = new_node
}

func (l *LinkedList) print() {
	if l.head == nil {
		return
	}
	current_node := l.head
	for current_node != nil {
		fmt.Println(current_node.data)
		current_node = current_node.next_node
	}
}

func (l *LinkedList) remove(data int) {
	current_node := l.head
	if current_node != nil && current_node.data == data {
		l.head = current_node.next_node
		current_node = nil
		return
	}

	previous_node := &Node{}
	for current_node != nil && current_node.data != data {
		// To Do : why the node isn't removed when the below two codes are switched.
		previous_node = current_node
		current_node = current_node.next_node
	}

	if current_node == nil {
		return
	}
	previous_node.next_node = current_node.next_node
	current_node = nil
}

func (l *LinkedList) reverse_iterative() {
	current_node := l.head
	var previous_node *Node = nil
	for current_node != nil {
		next_node := current_node.next_node
		current_node.next_node = previous_node
		previous_node = current_node
		// To Do : Why print only l.head element if the below code instead of current_node = next_node
		// current_node = current_node.next_node
		current_node = next_node
	}
	l.head = previous_node
}

func main() {
	l := LinkedList{}
	l.append(1)
	l.append(2)
	l.append(3)
	l.append(4)
	l.print()
	fmt.Println("#########")
	l.reverse_iterative()
	l.print()
}
