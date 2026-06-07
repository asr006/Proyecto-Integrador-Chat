package mensajes

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var actualizador = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clientesConectados = make(map[*websocket.Conn]bool)
var canalDifusion = make(chan Mensaje)

func ManejarConexiones(w http.ResponseWriter, r *http.Request) {
	conexion, err := actualizador.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error en la conexion", err)
		return
	}
	defer conexion.Close()

	clientesConectados[conexion] = true
	log.Println("¡Una Persona Se Unio A La Sala!")

	for {
		var nuevoMensaje Mensaje
		err := conexion.ReadJSON(&nuevoMensaje)
		if err != nil {
			delete(clientesConectados, conexion)
			break
		}
		canalDifusion <- nuevoMensaje
	}
}

func RepartirMensajes() {
	for {
		mensaje := <-canalDifusion
		for cliente := range clientesConectados {
			err := cliente.WriteJSON(mensaje)
			if err != nil {
				cliente.Close()
				delete(clientesConectados, cliente)
			}
		}
	}
}
