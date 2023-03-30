package main

import (
	"fmt"
	"guia3/linkedlist"
)

func main() {
	l := linkedlist.NewLinkedList[int]()
	fmt.Println("Agregamos 1, 2 y 3 al final de la lista")
	l.Append(1)
	l.Append(2)
	l.Append(3)
	fmt.Println(l)
	fmt.Println("Agregamos 0 al inicio de la lista")
	l.Prepend(0)
	fmt.Println(l)
	fmt.Println("Buscamos el numero 3")
	fmt.Println("Se encuentra en la posicion: ", l.Search(3))

	l.Remove(2)
	l.Remove(0)
	l.Remove(3)
	fmt.Println("Buscamos el numero 3 luego de removerlo de la lista")
	fmt.Println(l)
	fmt.Println("Se encuentra en la posicion: ", l.Search(3))
	l.Remove(1)
	l.Remove(1) //No deberia hacer nada
	l.Remove(1) //No deberia hacer nada
	fmt.Println(l)

	fmt.Println("Agregamos 0, 1 , 3, 4 y 5 al final de la lista")
	l.Append(0)
	l.Append(1)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	fmt.Println(l)
	fmt.Println("Agregamos 2 en la posicion 2")
	l.InsertAt(2, 2)

	fmt.Println(l)

}
