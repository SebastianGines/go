package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Realiza la conexión a la BD
var dsn = "root:@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en conexión con BD.", err)
		panic(err)
	} else {
		fmt.Println("Conexión con DB exitosa")
		return db
	}
}()
