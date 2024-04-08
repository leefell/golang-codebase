package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Rafael Callegari!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func creditos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Créditos: Alexandre Augusto"))
}

func main() {
	// HTTP é um protocolo de comunicação - Base da Comunicação WEB
	// Cliente (Faz request) - Servidor (processa request e envia response)
	// Rotas
	// URI - Identificador do recurso
	// Método - GET, POST, PUT, DELETE

	// HandleFunc recebe uma URI e um método
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)
	http.HandleFunc("/creditos", creditos)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
