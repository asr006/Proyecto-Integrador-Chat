package rutas

import (
	"backend-go/mensajes"
	"fmt"
	"net/http"
)

func ConfigurarRutas() {

	http.HandleFunc("/", Inicio)
	http.HandleFunc("/registro", Registro)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/mensaje", EnviarMensaje)
	http.HandleFunc("/ws", mensajes.ManejarConexiones)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sistema de Chat")
}

func Registro(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint de registro")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint de login")
}

func EnviarMensaje(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint para enviar mensajes")
}
