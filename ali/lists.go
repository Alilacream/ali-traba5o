package ali

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
		head *node
		lenght int
}
func (l *linkedlist) prepend (n *node) {
	second :=  l.head
	l.head = n
	l.head.next = second	
	l.lenght++
}
func main() {
	mylist := linkedlist{}			
	 node1 := &node{data:48}
	// node2 := &node{data:10}
	node3 := &node{data:8}
	mylist.prepend(node1)
	// mylist.prepend(node2)
	mylist.prepend(node3)
	fmt.Print(mylist)
}