package main

import (
	"backend-go/mensajes"
	"backend-go/rutas"
	"fmt"
	"net/http"
)

func main() {
	rutas.ConfigurarRutas()
	go mensajes.RepartirMensajes()

	fmt.Println("Servidor ejecutándose en http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
