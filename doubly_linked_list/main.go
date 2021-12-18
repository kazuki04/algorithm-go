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

func (d *DoublyLinkedList) insert(data int) {
	new_node := &Node{data: data}
	if d.head == nil {
		d.head = new_node
		return
	}

	d.head.previous_node = new_node
	new_node.next_node = d.head
	d.head = new_node
}

func (d *DoublyLinkedList) print() {
	current_node := d.head
	for current_node != nil {
		fmt.Println(current_node.data)
		current_node = current_node.next_node
	}
}

func (d *DoublyLinkedList) remove(data int) {
	current_node := d.head
	if current_node != nil && current_node.data == data {
		if current_node.next_node == nil {
			current_node = nil
			return
		} else {
			next_node := current_node.next_node
			d.head = next_node
			next_node.previous_node = nil
			current_node = nil
			return
		}
	}

	for current_node != nil && current_node.data != data {
		current_node = current_node.next_node
	}

	if current_node == nil {
		return
	}

	if current_node.next_node == nil {
		previous_node := current_node.previous_node
		previous_node.next_node = nil
		current_node = nil
		return
	} else {
		previous_node := current_node.previous_node
		next_node := current_node.next_node
		previous_node.next_node = next_node
		next_node.previous_node = previous_node
		current_node = nil
		return
	}
}

func main() {
	d := &DoublyLinkedList{nil}
	d.append(1)
	d.append(2)
	d.insert(0)
	d.print()
	d.remove(2)
	fmt.Println("##########")
	d.print()
}
