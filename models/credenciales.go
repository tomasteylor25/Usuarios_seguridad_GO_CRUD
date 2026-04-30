package models

import "time"

type Credenciales struct {
	IDUsuario          int        `json:"id_usuario"`
	ContrasenaHash     string     `json:"contrasena_hash"`
	IntentosFallidos   int        `json:"intentos_fallidos"`
	Bloqueado          bool       `json:"bloqueado"`
	UltimoLogin        *time.Time `json:"ultimo_login"`
	Activo             bool       `json:"activo"`
	FechaCreacion      time.Time  `json:"fecha_creacion"`
	FechaModificacion  time.Time  `json:"fecha_modificacion"`
}
