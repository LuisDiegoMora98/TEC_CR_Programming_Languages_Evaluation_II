package main

import (
	"fmt"
	"io"
)

//Item 8, representation of a binary tree
// BinaryNode and BinaryTree : Item 8
// nil is the null for maps, structs and similar structures, default value when not initialized

type BinaryNode struct {
	parent *BinaryNode
	left   *BinaryNode
	right  *BinaryNode
	value  int
}

type BinaryTree struct {
	root *BinaryNode
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}
