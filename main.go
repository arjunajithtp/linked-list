package main

import (
	"fmt"
	"github.com/arjunajithtp/linked-list/doubly"
	"github.com/arjunajithtp/linked-list/singly"
)

func main() {

	// example of usage for singly linked list
	fmt.Println("SINGLY LINKED LIST")
	sl := singly.List{}

	sl.Push(5, 6, 7, 8)
	fmt.Println("Pushed data: ", sl.Read())

	fmt.Println("Popped head item: ", sl.Pop())
	fmt.Println("Data after head pop: ", sl.Read())

	sl.Reverse()
	fmt.Println("Data after reversal: ", sl.Read())

	// example usage for doubly linked list
	fmt.Println("\nDOUBLY LINKED LIST")
	dl := doubly.List{}
	dl.Push(5, 6, 7, 8)
	fmt.Println("Pushed data in LIFO: ", dl.ReadLIFO())
	fmt.Println("Pushed data in FIFO: ", dl.ReadFIFO())

	fmt.Println("Popped head item: ", dl.PopHead())
	fmt.Println("Data after head pop in LIFO: ", dl.ReadLIFO())

	fmt.Println("Popped tail item: ", dl.PopTail())
	fmt.Println("Data after tail pop in LIFO: ", dl.ReadLIFO())

}
