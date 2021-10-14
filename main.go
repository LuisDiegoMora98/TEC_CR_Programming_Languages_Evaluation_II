package main

//Terminal: go get github.com/wcharczuk/go-chart@3a7bc55431138d52d74bf4a1388374c01e09445d
//go-chart doesn't show x and y labels due to go.mod, so, it's necessary a previous version

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/wcharczuk/go-chart"
)

func main() {
	fmt.Println("Hello")

	fmt.Println(item3([]int{1, 2, 43, 88, 10}, 18))
	fmt.Println(item3([]int{1, 2, 43, 88, 10}, 10))
	fmt.Println("Hello")

	valuesMap := map[int]float64{
		1: 100.0,
		2: 94.88,
		3: 4.74,
		4: 3.22,
		5: 3,
		6: 2.27,
		7: 1,
	}
	Item2(valuesMap)
	tree := &BinaryTree{}
	_, comp := tree.insert(5)
	fmt.Println(comp)
	_, comp = tree.insert(20)
	fmt.Println(comp)
	_, comp = tree.insert(50)
	fmt.Println(comp)
	_, comp = tree.insert(15)
	fmt.Println(comp)
	_, comp = tree.insert(60)
	fmt.Println(comp)
	_, comp = tree.insert(50)
	fmt.Println(comp)
	_, comp = tree.insert(60)
	fmt.Println(comp)
	_, comp = tree.insert(55)
	fmt.Println(comp)
	_, comp = tree.insert(85)
	fmt.Println(comp)
	_, comp = tree.insert(15)
	fmt.Println(comp)
	_, comp = tree.insert(5)
	fmt.Println(comp)
	_, comp = tree.insert(10)
	fmt.Println(comp)
	_, comp = tree.insert(2)
	fmt.Println(comp)
	_, comp = tree.insert(1)
	fmt.Println(comp)
	var answer bool
	answer, comp = tree.search(55)
	fmt.Println("Busqueda de 55", comp, answer)

	print(os.Stdout, tree.root, 0, 'M')
}

// Item2 receive a map to create bar chart->  (label)Key: int, (value)value: float64
func Item2(values map[int]float64) {

	//Sort of the Numbers in the map
	keys := make([]int, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	//Adding the values to the Bar Values List
	var charValues []chart.Value
	for _, k := range keys {
		valV := values[k]
		valL := strconv.Itoa(k)
		charValues = append(charValues, chart.Value{Label: valL, Value: valV})
	}

	graph := chart.BarChart{
		Title: "Number Aparition Bar Chart",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars:     charValues,
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func Item8() {

}

// BinaryNode and BinaryTree : Item 8
// nil is the null for maps, structs and similar structures, default value when not initialized

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	value int
}

type BinaryTree struct {
	root *BinaryNode
}

// Item 9
// Inspired on https://www.golangprograms.com/golang-program-to-implement-binary-tree.html

// Insertion from root value
// insert comparisons starts in 1, and ++ every recurrent insert
// receive the newValue and ,by reference the tree, returns the node itself and the comparisons number
func (tree *BinaryTree) insert(newValue int) (*BinaryTree, int) {
	comparisons := 1
	if tree.root == nil {
		tree.root = &BinaryNode{value: newValue, left: nil, right: nil}
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
			currentNode.left = &BinaryNode{value: newValue, left: nil, right: nil}
		} else {
			comparisons = currentNode.left.insert(newValue, comparisons)
		}
	} else {
		if currentNode.right == nil {
			currentNode.right = &BinaryNode{value: newValue, left: nil, right: nil}
		} else {
			comparisons = currentNode.right.insert(newValue, comparisons)
		}
	}
	return comparisons
}

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

//Item 3, receive an int Array and a nalue to insert, then it compares all elements in the array to the new value
//returns the number of comparisons until the number insertion or to determine the number is already in the array
func item3(arrayTo []int, newValue int) int {
	comparisons := 0
	for k := range arrayTo {
		if arrayTo[k] == newValue {
			return comparisons
		}
		comparisons++
	}
	arrayTo = append(arrayTo, newValue)
	fmt.Println(arrayTo)
	return comparisons
}
