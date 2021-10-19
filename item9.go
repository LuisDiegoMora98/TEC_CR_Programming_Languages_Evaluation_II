package main

// Item 9
// Inspired on https://www.golangprograms.com/golang-program-to-implement-binary-tree.html

// Insertion from root value
// insert comparisons starts in 1, and ++ every recurrent insert
// receive the newValue and ,by reference the tree, returns the node itself and the comparisons number
func (tree *BinaryTree) insert(newValue int) (*BinaryTree, int) {
	comparisons := 0
	if tree.root == nil {
		tree.root = &BinaryNode{value: newValue, parent: nil, left: nil, right: nil}
	} else {
		comparisons = tree.root.insert(newValue, comparisons)
	}
	return tree, comparisons
}

// Recursive insertion for the binary tree
// insert pass the current node, compares to newValue and recurse if needed, returns the comparisons number
func (currentNode *BinaryNode) insert(newValue int, comparisons int) int {
	comparisons++
	if currentNode == nil || currentNode.value == newValue { // Avoid repeated numbers insertion
		return comparisons
	} else if newValue < currentNode.value { //change to <= if you want equal values on right, < on left
		if currentNode.left == nil {
			currentNode.left = &BinaryNode{value: newValue, parent: currentNode, left: nil, right: nil}
		} else {
			comparisons = currentNode.left.insert(newValue, comparisons)
		}
	} else {
		if currentNode.right == nil {
			currentNode.right = &BinaryNode{value: newValue, parent: currentNode, left: nil, right: nil}
		} else {
			comparisons = currentNode.right.insert(newValue, comparisons)
		}
	}
	return comparisons
}
