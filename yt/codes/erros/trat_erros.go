package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func regraSoma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}

func soma() {
	res, err := regraSoma(10, 300)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)
}

func webRequest() {
	res, err := http.Get("https://google.com.br")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(res.Header)
}

// Mas e se eu não quiser tratar o erro?

func exemplo() {
	res, _ := regraSoma(6, 2)

	fmt.Println(res)
}

func main() {

}
