package main

import "fmt"

// Isso é gambiarra, pq todo tipo atenderia essa interface
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(1223))

	// Exemplo de uma gambiarra horrivel
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(321): true,
		"String":     "string",
		true:         float64(12),
	}

	fmt.Println(mapa)
}
