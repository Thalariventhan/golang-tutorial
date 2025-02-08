package main

import (
	"fmt"
	"golang-tutorial/examples/simple-http-server/routes"
	"net/http"
)

func main() {
	routes.RouterConfig()
	
	fmt.Println("Serving on port http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}