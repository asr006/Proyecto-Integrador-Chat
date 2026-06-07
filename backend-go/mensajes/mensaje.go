package mensajes

type Mensaje struct {
	ID        int    `json:"id"`
	Emisor    string `json:"emisor"`
	Receptor  string `json:"receptor"`
	Contenido string `json:"contenido"`
	Fecha     string `json:"fecha"`
}
