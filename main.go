package main

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type NodeList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *NodeList[T]) CreateNode(value T) *Node[T] {
	return &Node[T]{
		value: value,
		next:  nil,
	}
}

func (l *NodeList[T]) AppendNode(node *Node[T]) {
	// si una lista esta vacia, entonces no tiene nigun nodo.
	if l.head == nil {
		// creamos el inicio de la lista
		l.head = node
		l.tail = node
		return
	}

	// en caso de que si exista un nodo, hacemos este apunte a el nuevo nodo
	l.tail.next = node // primero seteamos la referencia a el siguiente nodo
	l.tail = node      // y luego actualizamos la posicion del nodo
}

func (l *NodeList[T]) CreateNodes(values ...T) {
	for _, value := range values {
		node := l.CreateNode(value)
		l.AppendNode(node)
	}
}

func (l *NodeList[T]) InsertAfter(node *Node[T], newNode *Node[T]) {
	newNode.next = node.next
	node.next = newNode
}

func (l *NodeList[T]) Delete(node *Node[T]) {
	if l.head == node {
		l.head = l.head.next
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		if currentNode.next == node {
			currentNode.next = currentNode.next.next
		}
		currentNode = currentNode.next
	}
}

func (l *NodeList[T]) Find(callback func(*Node[T]) bool) *Node[T] {
	currentNode := l.head
	for currentNode != nil {
		if callback(currentNode) {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

func (l *NodeList[T]) Print() {
	node := l.head
	for node != nil {
		fmt.Printf("%v-> ", node.value)
		node = node.next
	}
}

func main() {
	linkedList := &NodeList[int]{}

	linkedList.CreateNodes(1, 2, 3, 54, 6, 76, 231, 21)

	findNode := linkedList.Find(func(n *Node[int]) bool { return n.value == 1 })

	newNode := linkedList.CreateNode(69)

	linkedList.InsertAfter(findNode, newNode)

	linkedList.Delete(findNode)

	linkedList.Print()

}
