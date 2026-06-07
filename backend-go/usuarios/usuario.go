package usuarios

type Usuario struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Correo          string `json:"correo"`
	Nickname        string `json:"nickname"`
	Password        string `json:"password"`
}
