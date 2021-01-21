package main

import (
	"container/list"
	"fmt"
)

func printALlElement(myList *list.List) {
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}

func main() {
	fmt.Println("Go Linked Lists Tutorial")
	myList := list.New()
	myList.PushBack(1)
	myList.PushFront("a")
	myList.PushBack(2)

	printALlElement(myList)
	fmt.Println("_______")

	// element := myList.Front()
	element := myList.Back()
	myList.Remove(element)
	printALlElement(myList)
}
