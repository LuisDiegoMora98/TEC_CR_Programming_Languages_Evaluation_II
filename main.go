package main

//Terminal: go get github.com/wcharczuk/go-chart@3a7bc55431138d52d74bf4a1388374c01e09445d
//go-chart doesn't show x and y labels due to go.mod, so, it's necessary a previous version

import (
	"fmt"
)

func main() {
	fmt.Println("Evaluation 2 for Programming Languages:")
	var unOrderedArray = []int {1000, 2, 3, 17, 50}
	fmt.Println(Item6(unOrderedArray, 5, 17))
	var orderedArray = []int {20, 40, 50, 75, 80, 82, 84, 86, 87, 88, 89, 90, 95, 100}
	fmt.Println(Item7(orderedArray, 14, 86, 1))
}

/*6. Diseñar y construir una función que haga la búsqueda de un valor entero (la llave) en un arreglo que no
esté ordenado, mediante el algoritmo de búsqueda secuencial. La función debe retornar un par ordenado:
un valor booleano que indique si encontró la llave buscada (true = la encontró, false = no la encontró),
junto con un entero que indique la cantidad de comparaciones realizadas hasta determinar si la llave está o
no presente en el arreglo.*/
func Item6(pArray []int, pSize int, pValue int)(bool, int){
	for index := 0; index < pSize; index++ {
		if pArray[index] == pValue {
			return true, index + 1
		}
	}
	return false, 0
}

/*Diseñar y construir una función que haga la búsqueda de un valor entero (la llave) en un arreglo que esté
ordenado, mediante el algoritmo de búsqueda binaria. La función debe retornar un par ordenado: un valor
booleano que indique si encontró la llave buscada (true = la encontró, false = no la encontró), junto con un
entero que indique la cantidad de comparaciones realizadas hasta determinar si la llave está o no presente
en el arreglo.*/
/*
		20 40 50 75 80 82 84 86 87 88 89 90 95 100
	90					                 *
*/
func Item7(pArray []int, pSize int, pValue int, pComparisons int)(bool, int){
	if pSize == 0 {
		return false, 0
	}
	if pSize <= 2 {
		if (pArray[0] == pValue) || (pArray[1] == pValue) {
			return true, pComparisons
		}
		return false, pComparisons
	}
	pivot := pSize / 2
	if pArray[pivot] == pValue {
		return true, pComparisons
	}
	if pArray[pivot] > pValue {
		return Item7(pArray[:pivot], pivot, pValue,pComparisons + 1)
	}
	return Item7(pArray[pivot:], pSize - pivot, pValue,pComparisons + 1)
}

/*11. Diseñar y construir una función que transforme un árbol binario de búsqueda ordinario en un árbol binario
de búsqueda quasi-completo, conforme lo logra el algoritmo DSW. Un árbol de altura (o profundidad) a es
quasi-completo cuando los nodos más profundos se encuentran todos en los niveles a o a – 1. Base su
función en el algoritmo DSW (Day-Stout-Warren); ver referencias al final y en el tecDigital.*/
func Item11(){

}
