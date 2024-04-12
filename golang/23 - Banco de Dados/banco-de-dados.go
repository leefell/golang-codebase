package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Construir a string de conexão.
	stringConexao := "golang:golang@/testeBanco?charset=utf8&parseTime=True&loc=Local"

	// Abrir a conexão com o banco de dados MySQL.
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println("Dentro do SQL.OPEN")
		log.Fatal(erro)
	}
	defer db.Close()

	// Verificar se a conexão com o banco de dados é válida.
	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
