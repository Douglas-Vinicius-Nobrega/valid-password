package main

import (
	"fmt"
	"valid-password/routes"
)

func main() {
	fmt.Println("Starting server")
	fmt.Println("Literning on port 8080")
	routes.HandleRequest()
}