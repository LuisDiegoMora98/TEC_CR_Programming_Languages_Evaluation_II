package main

//Terminal: go get github.com/wcharczuk/go-chart@3a7bc55431138d52d74bf4a1388374c01e09445d
//go-chart doesn't show x and y labels due to go.mod, so, it's necessary a previous version

import "fmt"

func main() {
	fmt.Println("Hello")

	HolaMundo := map[string]float64{
		"Hola":    float64(1),
		"Mundo":   float64(2),
		"Pero":    float64(3),
		"En":      float64(4),
		"Gráfico": float64(5),
		"De":      float64(6),
		"Barras":  float64(7),
	}

	Item2Palabras(HolaMundo, "HolaMundo")

	// Se generan las listas con el respectivo contenido, para 200, 400, 600, 800 y 1000 elementos
	item1 := experimento_a()
	array10000 := gen_array(10000)

	//EXPERIMENTO CON 200 elementos
	Distr := experimento_b(item1[0])
	TS := experimento_c(item1[0])
	TOS := experimento_d(item1[0])
	TOQ := experimento_e(item1[0])
	ABB := experimento_f(item1[0])

	fmt.Println("TAMA;O DE TS ", len(item1[0]))
	//experimento_g(item1[0])		//necesario item 11
	experimento_h(Distr, "Distr 200")
	//Experimentos i
	comparacionesTS := experimento_i_1(array10000, TS, "200")
	comparacionesTOS := experimento_i_2(array10000, TOS, "200")
	comparacionesTOQ := experimento_i_3(array10000, TOQ, "200")
	comparacionesABB := experimento_i_4(array10000, *ABB, "200")
	experimentoi := map[string]float64{
		"TS":  float64(comparacionesTS),
		"TOS": float64(comparacionesTOS),
		"TOQ": float64(comparacionesTOQ),
		"ABB": float64(comparacionesABB),
	}
	Item2Palabras(experimentoi, "Experimento i, 200 Elementos")

	//EXPERIMENTO CON 400 elementos
	Distr = experimento_b(item1[1])
	TS = experimento_c(item1[1])
	TOS = experimento_d(item1[1])
	TOQ = experimento_e(item1[1])
	ABB = experimento_f(item1[1])
	//experimento_g(item1[1])		//necesario item 11
	experimento_h(Distr, "Distr 400")
	//Experimentos i
	comparacionesTS = experimento_i_1(array10000, TS, "400")
	comparacionesTOS = experimento_i_2(array10000, TOS, "400")
	comparacionesTOQ = experimento_i_3(array10000, TOQ, "400")
	comparacionesABB = experimento_i_4(array10000, *ABB, "400")
	experimentoi = map[string]float64{
		"TS":  float64(comparacionesTS),
		"TOS": float64(comparacionesTOS),
		"TOQ": float64(comparacionesTOQ),
		"ABB": float64(comparacionesABB),
	}
	Item2Palabras(experimentoi, "Experimento i, 400 Elementos")
	fmt.Println(comparacionesTS)
	fmt.Println(comparacionesTOS)
	fmt.Println(comparacionesTOQ)
	fmt.Println(comparacionesABB)

	//EXPERIMENTO CON 600 elementos
	Distr = experimento_b(item1[2])
	TS = experimento_c(item1[2])
	TOS = experimento_d(item1[2])
	TOQ = experimento_e(item1[2])
	ABB = experimento_f(item1[2])
	//experimento_g(item1[2])		//necesario item 11
	experimento_h(Distr, "Distr 600")
	//Experimentos i
	comparacionesTS = experimento_i_1(array10000, TS, "600")
	comparacionesTOS = experimento_i_2(array10000, TOS, "600")
	comparacionesTOQ = experimento_i_3(array10000, TOQ, "600")
	comparacionesABB = experimento_i_4(array10000, *ABB, "600")
	experimentoi = map[string]float64{
		"TS":  float64(comparacionesTS),
		"TOS": float64(comparacionesTOS),
		"TOQ": float64(comparacionesTOQ),
		"ABB": float64(comparacionesABB),
	}
	Item2Palabras(experimentoi, "Experimento i, 600 Elementos")
	fmt.Println(comparacionesTS)
	fmt.Println(comparacionesTOS)
	fmt.Println(comparacionesTOQ)
	fmt.Println(comparacionesABB)

	//EXPERIMENTO CON 800 elementos
	Distr = experimento_b(item1[3])
	TS = experimento_c(item1[3])
	TOS = experimento_d(item1[3])
	TOQ = experimento_e(item1[3])
	ABB = experimento_f(item1[3])
	//experimento_g(item1[3])		//necesario item 11
	experimento_h(Distr, "Distr 800")
	//Experimentos i
	comparacionesTS = experimento_i_1(array10000, TS, "800")
	comparacionesTOS = experimento_i_2(array10000, TOS, "800")
	comparacionesTOQ = experimento_i_3(array10000, TOQ, "800")
	comparacionesABB = experimento_i_4(array10000, *ABB, "800")
	experimentoi = map[string]float64{
		"TS":  float64(comparacionesTS),
		"TOS": float64(comparacionesTOS),
		"TOQ": float64(comparacionesTOQ),
		"ABB": float64(comparacionesABB),
	}
	Item2Palabras(experimentoi, "Experimento i, 800 Elementos")
	fmt.Println(comparacionesTS)
	fmt.Println(comparacionesTOS)
	fmt.Println(comparacionesTOQ)
	fmt.Println(comparacionesABB)

	//EXPERIMENTO CON 1000 elementos
	Distr = experimento_b(item1[4])
	TS = experimento_c(item1[4])
	TOS = experimento_d(item1[4])
	TOQ = experimento_e(item1[4])
	ABB = experimento_f(item1[4])
	//experimento_g(item1[4])		//necesario item 11
	experimento_h(Distr, "Distr 1000")
	//Experimentos i
	comparacionesTS = experimento_i_1(array10000, TS, "1000")
	comparacionesTOS = experimento_i_2(array10000, TOS, "1000")
	comparacionesTOQ = experimento_i_3(array10000, TOQ, "1000")
	comparacionesABB = experimento_i_4(array10000, *ABB, "1000")
	experimentoi = map[string]float64{
		"TS":  float64(comparacionesTS),
		"TOS": float64(comparacionesTOS),
		"TOQ": float64(comparacionesTOQ),
		"ABB": float64(comparacionesABB),
	}
	Item2Palabras(experimentoi, "Experimento i, 1000 Elementos")
	fmt.Println(comparacionesTS)
	fmt.Println(comparacionesTOS)
	fmt.Println(comparacionesTOQ)
	fmt.Println(comparacionesABB)

	//findHeight(ABB.root)
	//ABB.treeDensity(1000)
	//Item11()
}
