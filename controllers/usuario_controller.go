package controllers

import (
	"Usuarios_seguridad_GO_CRUD/config"
	"Usuarios_seguridad_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL usuarios
func GetAllUsuarios(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		`SELECT id_usuario, correo, correo_verificado, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Usuario"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Usuario
	for rows.Next() {
		var u models.Usuario
		rows.Scan(&u.IDUsuario, &u.Correo, &u.CorreoVerificado, &u.Activo, &u.FechaCreacion, &u.FechaModificacion)
		list = append(list, u)
	}
	respondJSON(w, 200, list)
}

// GET BY ID usuario
func GetUsuarioByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var u models.Usuario

	err := config.DB.QueryRow(
		`SELECT id_usuario, correo, correo_verificado, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Usuario" WHERE id_usuario=$1`, id,
	).Scan(&u.IDUsuario, &u.Correo, &u.CorreoVerificado, &u.Activo, &u.FechaCreacion, &u.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Usuario no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, u)
}

// POST crear usuario
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var u models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO "Usuario_seguridad"."Usuario" (correo, correo_verificado, activo) 
		VALUES ($1, $2, $3) RETURNING id_usuario, fecha_creacion, fecha_modificacion`,
		u.Correo, u.CorreoVerificado, u.Activo,
	).Scan(&u.IDUsuario, &u.FechaCreacion, &u.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, u)
}

// PUT actualizar usuario
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var u models.Usuario
	json.NewDecoder(r.Body).Decode(&u)

	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Usuario" 
		SET correo=$1, correo_verificado=$2, activo=$3, fecha_modificacion=now() 
		WHERE id_usuario=$4`,
		u.Correo, u.CorreoVerificado, u.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Usuario actualizado correctamente"})
}

// DELETE usuario
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Usuario" SET activo=FALSE, fecha_modificacion=now() WHERE id_usuario=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Usuario eliminado (baja lógica)"})
}
