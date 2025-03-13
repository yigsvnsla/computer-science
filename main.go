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
	// Si una lista esta vacia, entonces no tiene nigun nodo.
	if l.head == nil {
		// Creamos el inicio de la lista
		l.head = node
		l.tail = node
		return
	}

	// En caso de que si exista un nodo, hacemos este apunte a el nuevo nodo
	l.tail.next = node // Primero seteamos la referencia a el siguiente nodo
	l.tail = node      // Luego actualizamos la posicion del nodo
}

func (l *NodeList[T]) CreateNodes(values ...T) {
	for _, value := range values {
		node := l.CreateNode(value)
		l.AppendNode(node)
	}
}

func (l *NodeList[T]) InsertAfter(node *Node[T], newNode *Node[T]) {
	newNode.next = node.next // La referencia siguiente a el nodo nuevo sera la referencial del nodo actual
	node.next = newNode      // La referencia siguiente del nodo actual sera el nuevo nodo
}

func (l *NodeList[T]) Remove(node *Node[T]) {
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
	currentNode := l.head // Obtenemos el nodo actual
	for currentNode != nil {
		if callback(currentNode) { // si cumple la condicion del callback, retornamos el nodo
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil // Como no se encontro nada se devuelve un valor nulo
}

func (l *NodeList[T]) Print() {
	fmt.Println("\n -----------------------")
	node := l.head
	for node != nil {
		fmt.Printf("%v-> ", node.value)
		node = node.next
	}
	fmt.Println("\n -----------------------")
}

func (l *NodeList[T]) RemoveAfter(node *Node[T]) {
	node.next = node.next.next // El siguiente nodo actual, se asigna a su siguiente referencia
}

func (l *NodeList[T]) Reverse() {

	var prev *Node[T] // Nodo previo inicializado en NULL
	current := l.head // Nodo actual es la cabeza de la lista

	for current != nil {
		next := current.next // Guardamos la referencia al siguiente nodo
		current.next = prev  // Invertimos el enlace del nodo actual
		prev = current       // Movemos prev a la posición actual
		current = next       // Avanzamos al siguiente nodo en la lista
	}

	l.head = prev // Restablecemos la referencia de la cabeza con el nuevo nodo
}

func main() {
	linkedList := &NodeList[int]{}

	linkedList.CreateNodes(1, 2, 3, 4, 5, 6, 7, 8)

	newNode := linkedList.CreateNode(9)

	linkedList.AppendNode(newNode)

	// linkedList.Remove(newNode)

	linkedList.Reverse()

	linkedList.Print()

}
