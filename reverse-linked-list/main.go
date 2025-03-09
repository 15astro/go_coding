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

func (linkedList *LinkedList) DeleteAtFront(){
	nextNode := linkedList.Head.Next
	linkedList.Head = nextNode
 }
 
 func (linkedList *LinkedList) Reverse(){
	var next *Node 
	var prev *Node
	c := linkedList.Head
	for c != nil{
		// save next node
		next = c.Next
		// switch link to pre node
		c.Next = prev
		// save prev node
		prev = c
		// move forwaed
		c = next
	}
	linkedList.Head = prev
 }

func main(){
	myList := NewLinkedList("A")
	//fmt.Println(myList.Head.Val)
	myList.AddToFront("B")
	//fmt.Println(myList.Head.Val)
	myList.AddToFront("C")
	//fmt.Println(myList.Head.Val)
	//myList.DeleteAtFront()
	fmt.Println(myList.Show())
	myList.Reverse()
	fmt.Println(myList.Show())
}