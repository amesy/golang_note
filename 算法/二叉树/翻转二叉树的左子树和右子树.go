package main

import "fmt"

type treeNode struct {
	value  int
	left, right *treeNode
}

func createTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node *treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil, node Ignored.")
		return
	}

	node.value = value
}

func (node *treeNode) reverse(v1, v2 *treeNode) {
	if node == nil {
		fmt.Println("node is nil, please check it.")
		return
	}

	node.left, node.right = v2, v1
}

func main() {
	var root treeNode

	root = treeNode{value: 3}

	root.right = &treeNode{5, nil, nil}
	//root.right.left = &treeNode{}
	root.right.left = new(treeNode)

	root.left = &treeNode{}
	root.left.right = createTreeNode(2)

	// setvalue
	root.left.setValue(4)

	root.print()
	root.left.print()
	root.left.right.print()
	root.right.print()
	root.right.left.print()
	fmt.Println()

	//reverse
	root.reverse(root.left, root.right)

	root.left.print()
	root.right.print()
}