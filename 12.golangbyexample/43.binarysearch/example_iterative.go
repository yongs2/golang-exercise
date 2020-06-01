package main

import (
	"fmt"
)

type bstnode struct {
	value int
	left  *bstnode
	right *bstnode
}

type bst struct {
	root *bstnode
}

func (b *bst) reset() {
	b.root = nil
}

func (b *bst) insert(value int) {
	b.insertRec(b.root, value)
}

func (b *bst) insertRec(node *bstnode, value int) *bstnode {
	if b.root == nil {
		b.root = &bstnode{value: value}
		return b.root
	}
	if node == nil {
		return &bstnode{value: value}
	}
	if value <= node.value {
		node.left = b.insertRec(node.left, value)
	}
	if value > node.value {
		node.right = b.insertRec(node.right, value)
	}
	return node
}

func (b *bst) find(value int) error {
	node := b.findRec(b.root, value)
	if node == nil {
		return fmt.Errorf("Value: %d not found in tree", value)
	}
	return nil
}

func (b *bst) findRec(node *bstnode, value int) *bstnode {
	if node == nil {
		return nil
	}
	if node.value == value {
		return b.root
	}
	if value < node.value {
		return b.findRec(node.left, value)
	}
	return b.findRec(node.right, value)
}

func (b *bst) inorder() {
	b.inorderRec(b.root)
}

func (b *bst) inorderRec(node *bstnode) {
	if node != nil {
		b.inorderRec(node.left)
		fmt.Println(node.value)
		b.inorderRec(node.right)
	}
}

func main() {
	var value int
	var err error

	bst := &bst{}
	eg := []int{2, 5, 7, -1, -1, 5, 5}
	for _, val := range eg {
		bst.insert(val)
	}
	fmt.Printf("Printing Inorder:\n")
	bst.inorder()

	bst.reset()
	eg = []int{4, 5, 7, 6, -1, 99, 5}
	for _, val := range eg {
		bst.insert(val)
	}
	fmt.Printf("\nPrinting Inorder:\n")
	bst.inorder()

	value = 2
	fmt.Printf("\nFinding Values:\n")
	err = bst.find(value)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", value)
	} else {
		fmt.Printf("Value %d Found\n", value)
	}

	value = 6
	fmt.Printf("\nFinding Values:\n")
	err = bst.find(value)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", value)
	} else {
		fmt.Printf("Value %d Found\n", value)
	}

	value = 5
	fmt.Printf("\nFinding Values:\n")
	err = bst.find(value)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", value)
	} else {
		fmt.Printf("Value %d Found\n", value)
	}

	value = 1
	fmt.Printf("\nFinding Values:\n")
	err = bst.find(value)
	if err != nil {
		fmt.Printf("Value %d Not Found\n", value)
	} else {
		fmt.Printf("Value %d Found\n", value)
	}
}