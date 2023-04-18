package ejercicios

import (
	"guia4/set"

	"github.com/agrison/go-commons-lang/stringUtils"
)

//1. Escribir una función que reciba una cadena y devuelva el conjunto de todas las letras de la cadena:

func Letras(s string) *set.Set[string] {
	set := set.NewSet[string]()

	for _, letra := range s {
		if stringUtils.IsAlpha(string(letra)) { //The isAlpha() method is used to check given character is an alphabet or not.
			set.Add(string(letra))
		}
	}
	return set
}
