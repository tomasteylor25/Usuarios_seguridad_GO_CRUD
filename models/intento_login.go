package models

import "time"

type IntentoLogin struct {
	IDIntento         int        `json:"id_intento"`
	IDUsuario         *int       `json:"id_usuario"`
	EmailIngresado    *string    `json:"email_ingresado"`
	Exitoso           bool       `json:"exitoso"`
	MotivoFallo       *string    `json:"motivo_fallo"`
	IPOrigen          *string    `json:"ip_origen"`
	Fecha             time.Time  `json:"fecha"`
	Activo            bool       `json:"activo"`
	FechaCreacion     time.Time  `json:"fecha_creacion"`
	FechaModificacion time.Time  `json:"fecha_modificacion"`
}
