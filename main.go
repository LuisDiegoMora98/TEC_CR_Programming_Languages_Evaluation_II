package main

//Terminal: go get github.com/wcharczuk/go-chart@3a7bc55431138d52d74bf4a1388374c01e09445d
//go-chart doesn't show x and y labels due to go.mod, so, it's necessary a previous version

import (
	"fmt"
	"os"
)

func main() {
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
