package main

import (
	"fmt"
	"linkedlist"
)

func main() {
	// Creo un conjunto de letras
	// letras := ejercicios.Letras("Holaaaa Mundo")
	// fmt.Println(letras)
	// Creo dos conjuntos de números
	// n1 := set.NewSet(1, 10, 5)
	// n2 := set.NewSet(5, 15, 1)
	// fmt.Println("Diferencia Simetrica: ", ejercicios.DiferenciaSimetrica(n1, n2))
	// fmt.Println("Diferencia: ", set.Difference(n1, n2))

	prueba := [5]int{}
	for _, i := range prueba {
		fmt.Println(prueba[i])
	}
	fmt.Println(prueba.esDeEnteros())
}
