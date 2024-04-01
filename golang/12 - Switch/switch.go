package main

import "fmt"

func main() {

	dia := diaDaSemana(5)
	fmt.Println(dia)

	fmt.Println("-----------")
	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)
}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Dia de pinga, Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Insira um número de 1 - 7"
	}
}

func diaDaSemana2(numero int) string {

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Dia de pinga, Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		return "Insira um número de 1 - 7"
	}

	return diaDaSemana
}
