package linkedlist

import (
	"errors"
	"fmt"
)

// node es el nodo de la lista enlazada
// contiene un valor y un puntero al siguiente nodo
// el valor es de tipo genérico, comparable
type node[T any] struct {
	value T
	next  *node[T]
}

// newNode crea un nuevo nodo, con el valor recibido
// y el puntero al siguiente nodo en nil
func newNode[T comparable](value T) *node[T] {
	return &node[T]{value: value, next: nil}
}

/************************************************************/

// LinkedList es la lista enlazada simple
// contiene punteros al primer nodo y al último
type LinkedList[T comparable] struct {
	head *node[T] // puntero al primer nodo
	tail *node[T] // puntero al último nodo
	size int
}

// NewLinkedList crea una nueva lista enlazada, vacía
// En nuestra implementación representamos la lista vacía
// con un puntero al primer nodo en nil
// y un puntero al último nodo en nil
// O(1)
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil}
}

// Append agrega un nuevo nodo, con el valor recibido, al final de la lista
// O(1)
func (l *LinkedList[T]) Append(value T) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return
	}
	l.tail.next = newNode
	l.tail = newNode
	l.size++
}

// Prepend agrega un nuevo nodo, con el valor recibido,
// al inicio de la lista
// O(1)
func (l *LinkedList[T]) Prepend(value T) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return
	}
	newNode.next = l.head
	l.head = newNode
	l.size++
}

// InsertAt agrega un nuevo nodo, con el valor recibido,
// en la posición recibida.
// Si la posición es inválida, no hace nada
// O(n)
func (l *LinkedList[T]) InsertAt(value T, position int) {
	if position < 0 {
		return
	}
	newNode := newNode(value)
	if position == 0 {
		l.Prepend(value)
		l.size++
		return
	}
	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
		l.size++
	}
	if current == nil {
		return
	}
	newNode.next = current.next
	current.next = newNode
}

// Remove elimina el primer nodo que contenga el valor recibido
// O(n)
func (l *LinkedList[T]) Remove(value T) {
	if l.head == nil {
		return // no hay nada que eliminar
	}
	if l.head.value == value {
		l.head = l.head.next
		l.size--
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.size--
			return
		}
		current = current.next
	}
}

// String devuelve una representación en cadena de la lista
// en el formato [1 2 3].
// Se puede usar para imprimir la lista con fmt.Println
// O(n)
func (l *LinkedList[T]) String() string {
	if l.head == nil {
		return "[]"
	}
	current := l.head
	result := "["
	for current != nil {
		result += fmt.Sprintf("%v", current.value)
		if current.next != nil {
			result += " "
		}
		current = current.next
	}
	result += "]"
	return result
}

// Search busca el primer nodo que contenga el valor recibido
// y devuelve su posición en la lista o -1 si no lo encuentra
// O(n)
func (l *LinkedList[T]) Search(value T) int {
	if l.head == nil {
		return -1
	}
	current := l.head
	position := 0
	for current != nil {
		if current.value == value {
			return position
		}
		current = current.next
		position++
	}
	return -1
}

// Get devuelve el valor del nodo en la posición recibida
// o un valor nulo si la posición es inválida
// O(n)
func (l *LinkedList[T]) Get(position int) (T, error) {
	if l.head == nil {
		var t T
		return t, errors.New("lista vacía")
	}
	current := l.head
	for current != nil && position > 0 {
		current = current.next
		position--
	}
	if current == nil {
		var t T
		return t, errors.New("posición inválida")
	}
	return current.value, nil
}

// Size devuelve la cantidad de nodos en la lista
// O(n)
func (l *LinkedList[T]) Size() int {
	if l.head == nil {
		return 0
	}
	current := l.head
	position := 0
	for current != nil {
		current = current.next
		position++
	}
	return position
}

func (l *LinkedList[T]) ConcatenarListas(l1, l2 *LinkedList[T]) *LinkedList[T] {

	// si lista 1 esta vacia, develve la lista 2
	if l1.head == nil {
		return l2
	}
	// si l2 esta vacia no hace nada y devuelve la lista 1
	if l2.head == nil {
		return l1
	}
	l1.tail.next = l2.head // aca apunto el nodo next de l1 al head de l2
	l1.tail = l2.tail      // aca apunto el ultimo nodo de l1 al ultimo de l2
	l1.size += l2.size     // sumo los size de ambas
	return l1
}

func (l *LinkedList[T]) IntercalarListas(l1, l2 *LinkedList[T]) *LinkedList[T] {

	if l1 == nil || l2 == nil {
		return nil
	}

	result := &LinkedList[T]{} // creo una linkedList vacia
	current1 := l1.head
	current2 := l2.head // los curren apuntan a los nodos head de las listas que recibo

	for current1 != nil && current2 != nil {
		result.Append(current1.value)
		result.Append(current2.value) // itero ambas y voy agregando el valor de una y otra
		current1 = current1.next
		current2 = current2.next // muevo los current al siguiente nodo
	}

	// Agregar los elementos restantes de la primera lista, uso el mismo current así empiezo donde termino el otro for de arriba, lo mismo para las 2 listas
	for current1 != nil {
		result.Append(current1.value)
		current1 = current1.next
	}
	for current2 != nil {
		result.Append(current2.value)
		current2 = current2.next
	}

	return result

}
