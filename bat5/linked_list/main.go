package main

import "fmt"

// START1 OMIT
type Node struct {
	val         int
	left, right *Node
}

// END1 OMIT

// START2 OMIT
func insert(root *Node) *Node {
	for i := 0; i < 10; i++ {
		node := &Node{val: i}
		if root == nil {
			root = node
		} else if root.left == nil {
			node.left, root = root, node
		} else {
			node.left, node.right = root.left, root
			root.left = nil
			root = node
		}
	}
	return root
}

// END2 OMIT

// START4 OMIT
func find(node *Node, val int) *Node {
	if node == nil {
		return nil
	}
	if node.val == val {
		return node
	}
	if left := find(node.left, val); left != nil {
		return left
	}
	return find(node.right, val)
}

// END4 OMIT

// START6 OMIT
func del(node *Node, val int) *Node {
	if node == nil {
		return nil
	}
	if node.val == val {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		var leftMost *Node
		for leftMost = node.left; leftMost.left != nil; leftMost = leftMost.left {
		}
		leftMost.left = node.right
		return node.left
	}
	node.left, node.right = del(node.left, val), del(node.right, val)
	return node
}

// END6 OMIT

func travel(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	travel(root.left)
	travel(root.right)
}

// START8 OMIT
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	preOrder(root.left)
	preOrder(root.right)
}

// END8 OMIT

// START9 OMIT
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Println(root.val)
	inOrder(root.right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Println(root.val)
}

// END9 OMIT

func main() {
	// START3 OMIT
	var root *Node
	root = insert(root)
	// END3 OMIT

	// START5 OMIT
	for i := 0; i < 20; i++ {
		fmt.Println(find(root, i))
	}
	// END5 OMIT

	fmt.Println("preOrder")
	preOrder(root)
	fmt.Println("inOrder")
	inOrder(root)
	fmt.Println("postOrder")
	postOrder(root)

	// START7 OMIT
	for i := 20; i >= 0; i-- {
		root = del(root, i)
	}
	// END7 OMIT

	travel(root)
}
