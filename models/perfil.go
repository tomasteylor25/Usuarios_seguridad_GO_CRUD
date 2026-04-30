package models

import "time"

type Perfil struct {
	IDPerfil          int        `json:"id_perfil"`
	IDUsuario         int        `json:"id_usuario"`
	Nombre            string     `json:"nombre"`
	Apellido          string     `json:"apellido"`
	Telefono          *string    `json:"telefono"`
	Whatsapp          string     `json:"whatsapp"`
	Municipio         string     `json:"municipio"`
	Barrio            *string    `json:"barrio"`
	Genero            *string    `json:"genero"`
	FechaNacimiento   *time.Time `json:"fecha_nacimiento"`
	Biografia         *string    `json:"biografia"`
	FotoPerfil        *string    `json:"foto_perfil"`
	Activo            bool       `json:"activo"`
	FechaCreacion     time.Time  `json:"fecha_creacion"`
	FechaModificacion time.Time  `json:"fecha_modificacion"`
}
