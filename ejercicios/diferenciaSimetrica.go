package ejercicios

import (
	"guia4/set"
)

//Escribir una función que reciba dos conjuntos A y B y devuelva la diferencia simétrica entre ambos.
//La diferencia simétrica es el conjunto de elementos que solo pertenecen a A o a B pero no a ambos a la vez.

func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {
	set1 := set.Difference(s2, s1)
	set2 := set.Difference(s1, s2)

	return set.Union(set1, set2)
}
