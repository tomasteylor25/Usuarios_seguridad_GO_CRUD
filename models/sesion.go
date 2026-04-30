package models

import "time"

type Sesion struct {
	IDSesion          int       `json:"id_sesion"`
	IDUsuario         int       `json:"id_usuario"`
	TokenSesion       string    `json:"token_sesion"`
	IPOrigen          *string   `json:"ip_origen"`
	UserAgent         *string   `json:"user_agent"`
	FechaInicio       time.Time `json:"fecha_inicio"`
	FechaExpiracion   time.Time `json:"fecha_expiracion"`
	Revocada          bool      `json:"revocada"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
