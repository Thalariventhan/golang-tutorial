package routes

import (
	"golang-tutorial/examples/simple-http-server/handlers"
	"net/http"
)

func RouterConfig() {
	http.HandleFunc("/", handlers.HelloHandler)
}