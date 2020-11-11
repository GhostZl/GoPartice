package main

import (
	"GoPartice/tree"
	"fmt"
)

type myNode struct {
	*tree.Node //内嵌
}

type myInt int

func (num myInt) add(value int) int {
	return int(num) + value
}

func (node *myNode) postOrder() {
	if node == nil || node.Node == nil {
		return
	}
	right := myNode{node.Right}
	left := myNode{node.Left}
	left.postOrder()
	right.postOrder()
	node.Print()

}
func main() {
	root := myNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Left.Right = &tree.Node{Value: 2}
	root.Right = tree.CreateNode(5)
	root.Right.Left = tree.CreateNode(4)
	fmt.Println("POST_ORDER")
	root.postOrder()
	fmt.Println("POST_traverse")
	root.Traverse()

	num := myInt(3)
	fmt.Println(num.add(4))
}
