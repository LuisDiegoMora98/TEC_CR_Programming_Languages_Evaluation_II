package main

import (
	"fmt"
	"strconv"
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
	var insertions []int
	for k := range arrayA {
		insertions = append(insertions, item3(&TS, arrayA[k]))
	}
	experimento_h(experimento_b(insertions), "Inserciones Lineales "+strconv.Itoa(len(arrayA)))
	return TS
}

func experimento_d(arrayA []int) []int {
	//EXPERIMENTO D
	var TOS = make([]int, len(arrayA))
	copy(TOS, arrayA)
	TOS = selectionsort(TOS)
	fmt.Println("\n--- Ordenados ---\n\n", TOS)
	return TOS
}

func experimento_e(arrayA []int) []int {
	//EXPERIMENTO E
	var TOQ = make([]int, len(arrayA))
	copy(TOQ, arrayA)
	TOQ = quicksort(TOQ)
	fmt.Println("\n--- Ordenados ---\n\n", TOQ)
	return TOQ
}

func experimento_f(arrayA []int) *BinaryTree {
	//EXPERIMENTO F
	var insertions []int
	Abb := &BinaryTree{}
	for i := 0; i < len(arrayA); i++ {
		_, temp := Abb.insert(arrayA[i])
		insertions = append(insertions, temp)
	}
	//print(os.Stdout, Abb.root, 0, 'M')
	experimento_h(experimento_b(insertions), "Inserciones ABB "+strconv.Itoa(len(arrayA)))
	return Abb
}

func experimento_h(Distr [54]int, num string) {
	values := make(map[int]float64)
	for k := range Distr {
		values[k] = float64(Distr[k])
	}
	Item2(values, num+" Elementos")
}

//Using the item6, search in array 10000 values, get the number of comparisons
func experimento_i_1(arr_generado []int, arr_busqueda []int, num string) int {
	var counter int
	var temp int
	var comparisons []int
	for k := range arr_generado {
		_, temp = Item6(arr_busqueda, len(arr_busqueda), arr_generado[k])
		counter += temp
		comparisons = append(comparisons, temp)
	}
	experimento_h(experimento_b(comparisons), "Busqueda de 10000 Elementos en TS de "+num)
	return counter
}

//Using the item7, search in array 10000 values, get the number of comparisons
func experimento_i_2(arr_generado []int, arr_busqueda []int, num string) int {
	var counter int
	//variable := "hi"
	var comparisons []int
	for k := range arr_generado {
		temp := 0
		_, temp = Item7(arr_busqueda, len(arr_busqueda), arr_generado[k], temp)
		counter += temp
		comparisons = append(comparisons, temp)
	}
	experimento_h(experimento_b(comparisons), "Busqueda de 10000 Elementos en TOS de "+num)
	return counter
}

//Using the item7, search in a sorted array via quicksort for 10000 values, get the number of comparisons
func experimento_i_3(arreglo []int, busqueda []int, num string) int {
	var comparisons int
	var comparisonsList []int
	for k := range arreglo {
		temp := 0
		_, temp = Item7(busqueda, len(busqueda), arreglo[k], temp)
		comparisons += temp
		comparisonsList = append(comparisonsList, temp)
	}
	experimento_h(experimento_b(comparisonsList), "Busqueda de 10000 Elementos en TOQ de "+num)
	return comparisons
}

//Using the item10, search in a binary tree 10000 values, get the number of comparisons
func experimento_i_4(arreglo []int, tree BinaryTree, num string) int {
	var comparisons int
	var comparisonsList []int
	for i := 0; i < len(arreglo); i++ {
		_, temp := tree.search(arreglo[i])
		comparisons += temp
		comparisonsList = append(comparisonsList, temp)
	}
	/*for k := range arreglo {
		_, temp := tree.search(arreglo[k])
		comparisons += temp
		comparisonsList = append(comparisonsList,temp)
	}*/
	experimento_h(experimento_b(comparisonsList), "Busqueda de 10000 Elementos en ABB de "+num)
	return comparisons
}
