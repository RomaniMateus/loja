package main

import (
	"net/http"

	"github.com/loja-1/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
