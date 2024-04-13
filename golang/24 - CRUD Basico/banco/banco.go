package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de Conexão com o MySQL
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/testeBanco?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	// Testa se a conexão com o Banco de Dados está ativa, se Ping nao retornar erro
	// A conexão esta ativa.
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
