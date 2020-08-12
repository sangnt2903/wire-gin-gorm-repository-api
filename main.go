package main

import (
	"golang-repository-gin-gorm-wire/config"
	"golang-repository-gin-gorm-wire/routes"
)

func main(){
	db := config.ConnectDatabase()

	r := routes.SetRoutesConfig(db)

	r.Run()
}

