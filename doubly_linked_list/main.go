package main

import "fmt"

type Node struct {
	data          int
	previous_node *Node
	next_node     *Node
}

type DoublyLinkedList struct {
	head *Node
}

func (d *DoublyLinkedList) append(data int) {
	new_node := &Node{data: data}
	if d.head == nil {
		d.head = new_node
		return
	}
	current_node := d.head
	for current_node.next_node != nil {
		current_node = current_node.next_node
	}
	current_node.next_node = new_node
	new_node.previous_node = current_node
}

func main() {
	d := &DoublyLinkedList{nil}
	d.append(1)
	d.append(2)
	fmt.Println(d.head.data)
	fmt.Println(d.head.next_node.data)
}
