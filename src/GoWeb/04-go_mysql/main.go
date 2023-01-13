package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	//db.Ping()
	//fmt.Println(db.ExistTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	//user := models.CreateUser("Roel", "roel1234", "roel@gmail.com")
	//fmt.Println(user)
	//db.TruncateTable("users")
	//users := models.ListUsers()
	//user := models.GetUser(2)
	//user.UserName = "Juan"
	//user.Password = "Juan7890"
	//user.Save()
	db.TruncateTable("users")
	user := models.CreateUser("Roel", "roel1234", "roel@gmail.com")
	fmt.Println(user)
	user = models.CreateUser("Seba", "seba8791", "seba@gmail.com")
	fmt.Println(user)
	fmt.Println(models.ListUsers())
	db.Close()
}
