package config

import (
	"fmt"
	_ "github.com/bmizerany/pq"
	"github.com/jinzhu/gorm"
	"golang-repository-gin-gorm-wire/models/roles"
)

func ConnectDatabase() *gorm.DB{
	username := "postgres"
	password := "S@ng29031998"
	dbName := "test_pipeline"
	dbHost := "localhost"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(roles.Role{})
	return db
}
