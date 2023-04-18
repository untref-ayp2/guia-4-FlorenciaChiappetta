package ejercicios

import (
	"guia4/set"
)

// Escribir una función que reciba una cantidad variables de conjuntos de elementos comparables
// y devuelva la intersección de todos ellos.
func Interseccion[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] {
	arreglo := conjuntos[0]
	for _, c := range conjuntos {
		arreglo = set.Intersection(arreglo, c)
	}

	return arreglo
}
