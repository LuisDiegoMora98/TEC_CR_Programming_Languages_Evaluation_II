package main

import (
	"fmt"
	"os"
)

func experimento_a() [][]int {
	A := gen_array(200)
	B := gen_array(400)
	C := gen_array(600)
	D := gen_array(800)
	E := gen_array(1000)
	lista := [][]int{A, B, C, D, E}
	return lista
}

func experimento_b(arrayA []int) [54]int {
	var aparitions [54]int
	for k := range arrayA {
		aparitions[arrayA[k]]++
	}
	return aparitions
}

func experimento_c(arrayA []int) []int {
	var TS []int
	for k := range arrayA {
		item3(&TS, arrayA[k])
	}
	return TS
}

func experimento_d(arrayA []int) []int {
	//EXPERIMENTO D
	TOS := arrayA
	TOS = selectionsort(TOS)
	fmt.Println("\n--- Ordenados ---\n\n", TOS)
	return TOS
}

func experimento_e(arrayA []int) []int {
	//EXPERIMENTO E
	TOQ := arrayA
	TOQ = quicksort(TOQ)
	fmt.Println("\n--- Ordenados ---\n\n", TOQ)
	return TOQ
}

func experimento_f(arrayA []int) *BinaryTree {
	//EXPERIMENTO F
	Abb := &BinaryTree{}
	for _, value := range arrayA {
		Abb.insert(value)
	}
	print(os.Stdout, Abb.root, 0, 'M')
	return Abb
}

func experimento_h(Distr [54]int, num string) {
	values := make(map[int]float64)
	for k := range Distr {
		values[k] = float64(Distr[k])
	}
	Item2(values, "Distr, "+num+" Elementos")
}

//Using the item6, search in array 10000 values, get the number of comparisons
func experimento_i_1(arr_generado []int, arr_busqueda []int) int {
	var counter int
	var temp int
	for k := range arr_generado {
		_, temp = Item6(arr_busqueda, len(arr_generado), arr_generado[k])
		counter += temp
	}
	return counter
}

//Using the item6, search in array 10000 values, get the number of comparisons
func experimento_i_2(arr_generado []int, arr_busqueda []int) int {
	var counter int
	//variable := "hi"
	for k := range arr_generado {
		temp := 0
		_, temp = Item7(arr_busqueda, len(arr_busqueda)-1, arr_generado[k], temp)
		counter += temp
	}
	return counter
}

//Using the item7, search in a sorted array via quicksort for 10000 values, get the number of comparisons
func experimento_i_3(arreglo []int, busqueda []int) int {
	var comparisons int
	for k := range arreglo {
		aux := 0
		_, aux = Item7(busqueda, len(busqueda)-1, arreglo[k], aux)
		comparisons += aux
	}
	return comparisons
}

//Using the item10, search in a binary tree 10000 values, get the number of comparisons
func experimento_i_4(arreglo []int, tree BinaryTree) int {
	var comparisons int
	for k := range arreglo {
		_, aux := tree.search(arreglo[k])
		comparisons += aux
	}
	return comparisons
}
