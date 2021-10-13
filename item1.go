package main

import (
	"fmt"
	"math/rand"
	"time"
)

// prime returns true if n is a prime number.
func is_prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func gen_random_prime() int {
	min := 11
	max := 101
	rand.Seed(time.Now().UnixNano())
	randint := rand.Intn(max-min) + min
	val := is_prime(4)
	for val == false { // emulates while
		// do something
		randint = rand.Intn(max-min) + min
		val = is_prime(randint)
	}
	return randint
}

func lcg(a, c, m, semilla uint32) func() uint32 {
	r := semilla
	return func() uint32 {
		r = (a*r + c) % m
		return r
	}
}

func msg(semilla uint32) func() uint32 {
	g := lcg(214013, 2531011, 2048, semilla)
	return func() uint32 {
		return g()
	}
}

func gen_array(n int) {
	fmt.Println("\n--- Generando: --- ")
	slice := make([]int, n, n)
	semilla := gen_random_prime()
	fmt.Println("\n--- Valor semilla: --- ", semilla)
	n_semilla := uint32(semilla)
	msf := msg(n_semilla)
	for i := 0; i < n; i++ {
		temp := msf()
		slice[i] = int(temp) % 54
		fmt.Printf("%d\n", slice[i])
	}
}

func main() {
	gen_array(100)
}
