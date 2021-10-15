package main

import "fmt"
import "os"

func experimento_a() [][]int {
	A := gen_array(100)
	B := gen_array(400)
	C := gen_array(600)
	D := gen_array(800)
	E := gen_array(1000)
	lista := [][]int{A, B, C, D, E}
	return lista
}

func experimento_b(arrayA []int, title string) {
	var aparitions [54]int
	for k := range arrayA {
		aparitions[arrayA[k]]++
	}

	values := make(map[int]float64)
	for k := range aparitions {
		values[k] = float64(aparitions[k])
	}
	fmt.Println(aparitions)
	Item2(values, title)
}

func experimento_c(arrayA []int) {
	var TS []int
	for k := range arrayA {
		fmt.Println(item3(&TS, arrayA[k]))
		//fmt.Println(TS[:])
	}

}

func experimento_d(arrayA []int) {
	//EXPERIMENTO D
	TOS := arrayA
	selectionsort(TOS)
	fmt.Println("\n--- Ordenados ---\n\n", TOS)
}

func experimento_e(arrayA []int) {
	//EXPERIMENTO E
	TOQ := arrayA
	quicksort(TOQ)
	fmt.Println("\n--- Ordenados ---\n\n", TOQ)
}

func experimento_f(arrayA []int) {
	//EXPERIMENTO F
	Abb := &BinaryTree{}
	for _, value := range arrayA {
		Abb.insert(value)
	}
	print(os.Stdout, Abb.root, 0, 'M')
}
