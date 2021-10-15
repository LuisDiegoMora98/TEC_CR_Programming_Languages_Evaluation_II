package main

// Item 10
func (tree *BinaryTree) search(searchValue int) (bool, int) {
	comparisons := 1
	answer := false
	if tree.root != nil {
		answer, comparisons = tree.root.search(searchValue, comparisons, answer)
	}
	return answer, comparisons
}

func (currentNode *BinaryNode) search(searchValue int, comparisons int, answer bool) (bool, int) {
	comparisons++
	if currentNode.value == searchValue {
		answer = true
	} else if searchValue < currentNode.value {
		if currentNode.left != nil {
			answer, comparisons = currentNode.left.search(searchValue, comparisons, answer)
		}
	} else if currentNode.right != nil {
		answer, comparisons = currentNode.right.search(searchValue, comparisons, answer)
	}
	return answer, comparisons
}
