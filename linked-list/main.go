package main

import "fmt"

type Node struct{
	Val string
	Next *Node
}

type LinkedList struct{
  Head  *Node
}

func NewLinkedList(Val string) *LinkedList{
	newNode := &Node{
		Val: Val,
	}

	return &LinkedList{
		Head : newNode,
	}
}

func (linkedList *LinkedList) AddToFront(Val string){
   oldHeadNode := linkedList.Head
   newNode := &Node{Val: Val}
   newNode.Next = oldHeadNode
   linkedList.Head = newNode
} 

func (linkedList *LinkedList) Show() []string{

	var listItems []string
	h := linkedList.Head
	listItems = append(listItems, h.Val)
	for h.Next != nil{
		h = h.Next
		listItems = append(listItems, h.Val)
	}
     return listItems
}

func main(){
	myList := NewLinkedList("A")
	fmt.Println(myList.Head.Val)
	myList.AddToFront("B")
	fmt.Println(myList.Head.Val)
	myList.AddToFront("C")
	fmt.Println(myList.Head.Val)
	fmt.Println(myList.Show())
}