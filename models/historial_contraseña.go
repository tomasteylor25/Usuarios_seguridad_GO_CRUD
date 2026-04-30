package models

import "time"

type HistorialContraseña struct {
	IDHistorial       int       `json:"id_historial"`
	IDUsuario         int       `json:"id_usuario"`
	ContrasenaHash    string    `json:"contrasena_hash"`
	FechaCambio       time.Time `json:"fecha_cambio"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
