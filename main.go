package main

import (
  	"github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/models"
    "github.com/frendysanusi05/Seleksi-Asisten-Laboratorium-Programming-SingleService/routes"
)

func main() {
	models.ConnectDatabase()
    routes.Route()
}