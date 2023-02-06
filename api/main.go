package main

import (
	"fmt"

	"github.com/lucasalustiano/TCC/api/routes"
)

func main() {
  fmt.Println("Servidor Iniciado!")
  
	routes.HandleResquests()
}