package main

import (
	"GO_CRUD_API/initializers"
	"GO_CRUD_API/models"
)

func init() {
	initializers.EnvLoadVariables()
	initializers.ConnectTodatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
