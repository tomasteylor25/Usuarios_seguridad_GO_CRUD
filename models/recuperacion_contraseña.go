package models

import "time"

type RecuperacionContraseña struct {
	IDRecuperacion    int        `json:"id_recuperacion"`
	IDUsuario         int        `json:"id_usuario"`
	Token             string     `json:"token"`
	Codigo            *string    `json:"codigo"`
	Usado             bool       `json:"usado"`
	FechaExpiracion   time.Time  `json:"fecha_expiracion"`
	Activo            bool       `json:"activo"`
	FechaCreacion     time.Time  `json:"fecha_creacion"`
	FechaModificacion time.Time  `json:"fecha_modificacion"`
}
