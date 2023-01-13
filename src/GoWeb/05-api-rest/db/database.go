package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database
const url = "root:@tcp(localhost:3306)/goweb_db"

var db *sql.DB

// Abrir conexión con DB
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión a DB exitosa!")
	db = conection
}

// Cerrar la conexión con DB
func Close() {
	db.Close()
	fmt.Println("Conexión a DB cerrada!")
}

// Verificar conexión con DB
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Pong...")
}

// Verifica si una tabla existe
func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return rows.Next()
}

// Crear una tabla
func CreateTable(schema string, name string) {
	if !ExistTable(name) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

// Reiniciar el registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return result, err
}

// Polimorfiscmo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return rows, err
}
