package main

import "fmt"

type Node struct{
	Val string
	Next *Node
	Prev *Node
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
   oldHeadNode.Prev = newNode
   newNode.Next = oldHeadNode
   linkedList.Head = newNode
   //linkedList.Head.Prev = &Node{}
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
	var prev *Node
	c := linkedList.Head
	for c != nil{
		
		// switch link to pre node
		c.Next, c.Prev = c.Prev, c.Next
		// save prev node
		prev = c
		// move forwaed
		c = c.Prev
	}
	linkedList.Head = prev
 }

func main(){
	myList := NewLinkedList("A")
	//fmt.Println(myList.Head.Val)
	myList.AddToFront("B")
	//fmt.Println(myList.Head.Val)
	myList.AddToFront("C")
	fmt.Println(myList.Head.Val)
	fmt.Println(myList.Head.Next.Val)
	fmt.Println(myList.Head.Prev)
	fmt.Println(myList.Head.Next.Prev.Val)
	//myList.DeleteAtFront()
	myList.Reverse()
	fmt.Println(myList.Show())
}