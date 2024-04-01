package main

import (
	"fmt"
	"reflect"
)

// Array -> lista com tamanho fixo
// Slice -> lista sem tamanho fixo

func main() {

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("----------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println("Tamanho do slice3: ", len(slice3))
	fmt.Println("Capacidade maxima do slice3: ", cap(slice3))

	fmt.Println("----------------------")
	slice3 = append(slice3, 1)
	fmt.Println(slice3)
	fmt.Println("Tamanho do slice3: ", len(slice3))
	fmt.Println("Capacidade maxima do slice3: ", cap(slice3))

	fmt.Println("----------------------")
	slice3 = append(slice3, 2)
	fmt.Println(slice3)
	fmt.Println("Tamanho do slice3: ", len(slice3))
	fmt.Println("Capacidade maxima do slice3: ", cap(slice3))

	fmt.Println("----------------------")
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println("Tamanho do slice4: ", len(slice4))
	fmt.Println("Capacidade maxima do slice4: ", cap(slice4))
}
