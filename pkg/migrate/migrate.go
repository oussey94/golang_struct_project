package main

import (
	"github.com/oussey94/golang_struct_project/pkg/db"
	"github.com/oussey94/golang_struct_project/pkg/models"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(&models.Article{})
}
