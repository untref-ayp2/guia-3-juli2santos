# Guia 3
## Listas

En la carpeta `/linkedlist` se encuentra una implementación de una lista enlazada simple, que soporta cualquier tipo genérico que sea comparable (comparable significa que puedo preguntar si dos elementos son iguales o distintos). La lista soporta las siguientes operaciones
- **_NewLinkedList:_** Crea una nueva lista vacía
- **_Append:_** Agrega un elemento al final de la lista
- **_Prepend:_** Agrega un elemento al principio de la lista
- **_InsertAt:_** Agrega un elemento en la posición dada
- **_Remove:_** Busca un elemento en la lista y lo elimina
- **_Search:_** Busca un elemento en la lista y me devuelve su posición
- **_Get:_** Devuelve el elemento de la lista que se encuentra en una posición dada
- **_Size:_** Devuelve la cantidad de elemento que tiene la lista
- **_String:_** Devuelve una cadena de caracteres que representa la lista (para usar con Print)
1. Modificar la implementacion de lista enlazada simple para que Size sea O(1)
2. Implementar en la carpeta `/slicelist`una lista usando slices en lugar de nodos enlazados, con las mismas operaciones que la lista enlazada. Analizar el orden de cada operación. 
3. Escribir una función que reciba como parámetros dos listas del mismo tipo y devuelva una lista resultante de concatenar la segunda lista al final de la primera
4. Escribir una función que reciba dos listas del mismo tipo y devuelva la lista que resulta de intercalar uno a uno los elementos de ambas listas.
5. Implementar una pila y una cola usando la lista enlazada simple