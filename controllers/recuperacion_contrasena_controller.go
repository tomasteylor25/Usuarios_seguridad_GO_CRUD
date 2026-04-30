package controllers

import (
	"Usuarios_seguridad_GO_CRUD/config"
	"Usuarios_seguridad_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL recuperaciones
func GetAllRecuperacionContrasena(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		`SELECT id_recuperacion, id_usuario, token, codigo, usado, fecha_expiracion, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."RecuperacionContrasena"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.RecuperacionContrasena
	for rows.Next() {
		var rc models.RecuperacionContrasena
		rows.Scan(&rc.IDRecuperacion, &rc.IDUsuario, &rc.Token, &rc.Codigo, &rc.Usado, &rc.FechaExpiracion, &rc.Activo, &rc.FechaCreacion, &rc.FechaModificacion)
		list = append(list, rc)
	}
	respondJSON(w, 200, list)
}

// GET BY ID recuperacion
func GetRecuperacionContrasenaByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var rc models.RecuperacionContrasena

	err := config.DB.QueryRow(
		`SELECT id_recuperacion, id_usuario, token, codigo, usado, fecha_expiracion, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."RecuperacionContrasena" WHERE id_recuperacion=$1`, id,
	).Scan(&rc.IDRecuperacion, &rc.IDUsuario, &rc.Token, &rc.Codigo, &rc.Usado, &rc.FechaExpiracion, &rc.Activo, &rc.FechaCreacion, &rc.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Recuperación no encontrada"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, rc)
}

// POST crear recuperacion
func CreateRecuperacionContrasena(w http.ResponseWriter, r *http.Request) {
	var rc models.RecuperacionContrasena
	if err := json.NewDecoder(r.Body).Decode(&rc); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO "Usuario_seguridad"."RecuperacionContrasena" (id_usuario, token, codigo, usado, fecha_expiracion, activo) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_recuperacion, fecha_creacion, fecha_modificacion`,
		rc.IDUsuario, rc.Token, rc.Codigo, rc.Usado, rc.FechaExpiracion, rc.Activo,
	).Scan(&rc.IDRecuperacion, &rc.FechaCreacion, &rc.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, rc)
}

// PUT actualizar recuperacion
func UpdateRecuperacionContrasena(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var rc models.RecuperacionContrasena
	json.NewDecoder(r.Body).Decode(&rc)

	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."RecuperacionContrasena" 
		SET usado=$1, activo=$2, fecha_modificacion=now() 
		WHERE id_recuperacion=$3`,
		rc.Usado, rc.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Recuperación actualizada correctamente"})
}

// DELETE recuperacion (baja lógica)
func DeleteRecuperacionContrasena(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."RecuperacionContrasena" SET activo=FALSE, fecha_modificacion=now() WHERE id_recuperacion=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Recuperación eliminada (baja lógica)"})
}
