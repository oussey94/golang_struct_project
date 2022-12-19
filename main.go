package main

import (
	"github.com/oussey94/golang_struct_project/pkg/db"
	"github.com/oussey94/golang_struct_project/pkg/router"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	r := router.SetupRouter()

	r.Run()
}
