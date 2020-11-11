package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	key         string
	Left, Right *Node
}

type arr struct {
	Id int
	Aa int
}

func InitArr() arr {
	return arr{}
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
func (node Node) Print() {
	fmt.Printf("%d ", node.Value)
}
func (node *Node) SetValue(value int) {
	node.Value = value
}
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
	// node.print()
}
