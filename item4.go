package main

import (
	//"fmt"
	"math/rand"
	"time"
)

// Genera un array con numeros aleatorios
func generateArray(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

//slice := generateArray(20)
	//fmt.Println("\n--- Desordenados --- \n\n", slice)
	//quicksort(slice)
	//fmt.Println("\n--- Ordenados ---\n\n", slice)