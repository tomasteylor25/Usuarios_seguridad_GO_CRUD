package models

import "time"

type CrearContraseña struct {
	IDUsuario          int       `json:"id_usuario"`
	Contrasena         string    `json:"contrasena"`
	ConfirmarContrasena string   `json:"confirmar_contrasena"`
	Activo             bool      `json:"activo"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaModificacion  time.Time `json:"fecha_modificacion"`
}
