//CÓDIGO FUNCIONAL HASTA treeToVine

package main

import (
	"fmt"
	"math"
	"os"
)

func creatBackbone() {}

/*11. Diseñar y construir una función que transforme un árbol binario de búsqueda ordinario en un árbol binario
de búsqueda quasi-completo, conforme lo logra el algoritmo DSW. Un árbol de altura (o profundidad) a es
quasi-completo cuando los nodos más profundos se encuentran todos en los niveles a o a – 1. Base su
función en el algoritmo DSW (Day-Stout-Warren); ver referencias al final y en el tecDigital.*/
func Item11() {
	tree := &BinaryTree{}
	tree.insert(8)
	tree.insert(3)
	tree.insert(6)
	tree.insert(1)
	tree.insert(4)
	tree.insert(7)
	tree.insert(10)
	tree.insert(14)
	tree.insert(13)
	print(os.Stdout, tree.root, 0, 'M')
	fmt.Println("Tree to Vine")
	tree.root = treeToVine(tree.root)
	print(os.Stdout, tree.root, 0, 'M')
	/**tree.root = balanceVine(*tree, 9)
	print(os.Stdout, tree.root, 0, 'M')*/
	fmt.Println("Vine to Tree")
	*tree = vineToTree(tree, 9)

	//vineToTree(tree, 9)
	print(os.Stdout, tree.root, 0, 'M')
	fmt.Println("Check point 3")

}

//Takes a tree root and turn the tree into an ordered right branch
func treeToVine(pRoot *BinaryNode) *BinaryNode {
	temp := pRoot
	for temp != nil {
		if temp.value < pRoot.value { //making sure we get the min as resultant root
			pRoot = temp
		}
		//The rest is DSW, smashing the tree into an ascendant list called Vine or backbone
		if temp.left != nil {
			temp = rightRotation(temp, temp.left)
		} else {
			temp = temp.right
		}
	}
	return pRoot
}

/*
Function CreateBalancedTree(int numOfNodes){
	int m = 2^(lg2(numOfNodes + 1)) – 1;
	make n – m right rotations
	while (m > 1)
	//tree has not attained a balanced state
		m = m/2
		make m left rotations starting from the top of the backbone
}*/

//NO FUNCTIONAL
//Takes a right branched tree and makes the necessary rotations to make the almost perfect tree
func vineToTree(pRoot *BinaryTree, pTreeSize int) BinaryTree {
	lowerQuasiHeight := math.Floor(math.Log2(float64(pTreeSize + 1))) //this gets the Tree height rounding down
	upperQuasiHeight := math.Ceil(math.Log2(float64(pTreeSize + 1)))  //same but rounding up
	//This is the candidate for closest balanced tree height respect the actual tree
	quasiHeight := int(lowerQuasiHeight)
	//this calculates how far are the sizes and chooses the closest one
	if (math.Abs((math.Pow(2, lowerQuasiHeight) - 1) - float64(pTreeSize))) <
		(math.Abs((math.Pow(2, upperQuasiHeight) - 1) - float64(pTreeSize))) {
		quasiHeight = int(lowerQuasiHeight)
	}
	//Number of nodes from a perfect balanced tree
	quasiSize := math.Pow(2, float64(quasiHeight)) - 1
	//Difference between quasi and actual tree sizes
	excessedNodes := math.Abs(float64(pTreeSize) - quasiSize)
	rightRotations := excessedNodes

	auxNode := pRoot.root
	fmt.Println("Checkpoint 1")
	for rightRotations > 1 {
		//make rotations over second nodes
		fmt.Println("Checkpoint 2")
		auxNode = rightRotation(auxNode, auxNode.right)

		rightRotations--
	}
	return *pRoot
}

func rightRotation(pRoot *BinaryNode, pChild *BinaryNode) *BinaryNode {
	if pRoot.parent == nil {
		pRoot.left = pChild.right
		pChild.right = pRoot
		pChild.parent = nil
		pRoot.parent = pChild
		return pChild
	}
	pRoot.left = pChild.right
	pChild.right = pRoot
	pRoot.parent.right = pChild
	pRoot.parent = pChild
	return pChild

}

func leftRotation(pRoot *BinaryNode, pChild *BinaryNode) *BinaryNode {
	if pRoot.parent == nil {
		pRoot.parent = pChild
		pChild.left = pRoot
		pRoot.right = nil
		return pChild.right
	}
	pRoot.parent.right = pChild
	pChild.parent = pRoot.parent.right
	pChild.left = pRoot
	pChild.right = nil
	return pChild.right
}

//Ches
func leftRotate(node BinaryNode) BinaryNode {
	if node.right != nil {
		rightChild := node.right
		node.right = rightChild.right
		rightChild.right = rightChild.left
		rightChild.left = node.left
		node.left = rightChild

		temp := node.value
		node.value = rightChild.value
		rightChild.value = temp
	}
	return node
}

//Ches
func balanceVine(tree BinaryTree, nodeCount int) BinaryNode {
	times := int(math.Floor(math.Log2(float64(nodeCount))))
	newRoot := *tree.root
	for i := 0; i < times; i++ {
		newRoot = leftRotate(newRoot)
		root := newRoot.right
		for j := 0; j < (nodeCount/2)-1; j++ {
			root = leftRotation(&newRoot, root)
			root = root.right
			print(os.Stdout, tree.root, 0, 'M')
		}
		nodeCount >>= 1
	}
	return newRoot
}
