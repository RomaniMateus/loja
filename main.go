package main

import (
	"net/http"

	"github.com/mateus/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
