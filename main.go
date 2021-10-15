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

	A := gen_array(100)
	B := gen_array(400)
	C := gen_array(600)
	D := gen_array(800)
	E := gen_array(1000)

	for _, value := range A {
		fmt.Printf("%d ", value)
	}
	fmt.Printf("\n")
	for _, value := range B {
		fmt.Printf("%d ", value)
	}
	fmt.Printf("\n")
	for _, value := range C {
		fmt.Printf("%d ", value)
	}
	fmt.Printf("\n")
	for _, value := range D {
		fmt.Printf("%d ", value)
	}
	fmt.Printf("\n")
	for _, value := range E {
		fmt.Printf("%d ", value)
	}

	TOS := A
	selectionsort(TOS)
	fmt.Println("\n--- Ordenados ---\n\n", TOS)

	TOQ := A
	quicksort(TOQ)
	fmt.Println("\n--- Ordenados ---\n\n", TOQ)

}
