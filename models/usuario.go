package models

import "time"

type Usuario struct {
	IDUsuario          int       `json:"id_usuario"`
	Correo             string    `json:"correo"`
	CorreoVerificado   bool      `json:"correo_verificado"`
	Activo             bool      `json:"activo"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaModificacion  time.Time `json:"fecha_modificacion"`
}
