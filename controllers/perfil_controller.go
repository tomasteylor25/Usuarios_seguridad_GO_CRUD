package controllers

import (
	"Usuarios_seguridad_GO_CRUD/config"
	"Usuarios_seguridad_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL perfiles
func GetAllPerfiles(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		`SELECT id_perfil, id_usuario, nombre, apellido, telefono, whatsapp, municipio, barrio, genero, fecha_nacimiento, biografia, foto_perfil, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Perfil"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Perfil
	for rows.Next() {
		var p models.Perfil
		rows.Scan(&p.IDPerfil, &p.IDUsuario, &p.Nombre, &p.Apellido, &p.Telefono, &p.Whatsapp, &p.Municipio, &p.Barrio, &p.Genero, &p.FechaNacimiento, &p.Biografia, &p.FotoPerfil, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)
		list = append(list, p)
	}
	respondJSON(w, 200, list)
}

// GET BY ID perfil
func GetPerfilByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Perfil

	err := config.DB.QueryRow(
		`SELECT id_perfil, id_usuario, nombre, apellido, telefono, whatsapp, municipio, barrio, genero, fecha_nacimiento, biografia, foto_perfil, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Perfil" WHERE id_perfil=$1`, id,
	).Scan(&p.IDPerfil, &p.IDUsuario, &p.Nombre, &p.Apellido, &p.Telefono, &p.Whatsapp, &p.Municipio, &p.Barrio, &p.Genero, &p.FechaNacimiento, &p.Biografia, &p.FotoPerfil, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Perfil no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, p)
}

// POST crear perfil
func CreatePerfil(w http.ResponseWriter, r *http.Request) {
	var p models.Perfil
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO "Usuario_seguridad"."Perfil" (id_usuario, nombre, apellido, telefono, whatsapp, municipio, barrio, genero, fecha_nacimiento, biografia, foto_perfil, activo) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id_perfil, fecha_creacion, fecha_modificacion`,
		p.IDUsuario, p.Nombre, p.Apellido, p.Telefono, p.Whatsapp, p.Municipio, p.Barrio, p.Genero, p.FechaNacimiento, p.Biografia, p.FotoPerfil, p.Activo,
	).Scan(&p.IDPerfil, &p.FechaCreacion, &p.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, p)
}

// PUT actualizar perfil
func UpdatePerfil(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Perfil
	json.NewDecoder(r.Body).Decode(&p)

	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Perfil" 
		SET nombre=$1, apellido=$2, telefono=$3, whatsapp=$4, municipio=$5, barrio=$6, genero=$7, fecha_nacimiento=$8, biografia=$9, foto_perfil=$10, activo=$11, fecha_modificacion=now() 
		WHERE id_perfil=$12`,
		p.Nombre, p.Apellido, p.Telefono, p.Whatsapp, p.Municipio, p.Barrio, p.Genero, p.FechaNacimiento, p.Biografia, p.FotoPerfil, p.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Perfil actualizado correctamente"})
}

// DELETE perfil (baja lógica)
func DeletePerfil(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Perfil" SET activo=FALSE, fecha_modificacion=now() WHERE id_perfil=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Perfil eliminado (baja lógica)"})
}
