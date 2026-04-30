package controllers

import (
	"Usuarios_seguridad_GO_CRUD/config"
	"Usuarios_seguridad_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL historial contraseñas
func GetAllHistorialContraseña(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		`SELECT id_historial, id_usuario, contrasena_hash, fecha_cambio, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."HistorialContrasena"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.HistorialContraseña
	for rows.Next() {
		var h models.HistorialContraseña
		rows.Scan(&h.IDHistorial, &h.IDUsuario, &h.ContrasenaHash, &h.FechaCambio, &h.Activo, &h.FechaCreacion, &h.FechaModificacion)
		list = append(list, h)
	}
	respondJSON(w, 200, list)
}

// GET BY ID historial
func GetHistorialContraseñaByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var h models.HistorialContraseña

	err := config.DB.QueryRow(
		`SELECT id_historial, id_usuario, contrasena_hash, fecha_cambio, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."HistorialContrasena" WHERE id_historial=$1`, id,
	).Scan(&h.IDHistorial, &h.IDUsuario, &h.ContrasenaHash, &h.FechaCambio, &h.Activo, &h.FechaCreacion, &h.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Historial no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, h)
}

// POST crear historial contraseña
func CreateHistorialContraseña(w http.ResponseWriter, r *http.Request) {
	var h models.HistorialContraseña
	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO "Usuario_seguridad"."HistorialContrasena" (id_usuario, contrasena_hash, activo) 
		VALUES ($1, $2, $3) RETURNING id_historial, fecha_cambio, fecha_creacion, fecha_modificacion`,
		h.IDUsuario, h.ContrasenaHash, h.Activo,
	).Scan(&h.IDHistorial, &h.FechaCambio, &h.FechaCreacion, &h.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, h)
}

// PUT actualizar historial
func UpdateHistorialContraseña(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var h models.HistorialContraseña
	json.NewDecoder(r.Body).Decode(&h)

	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."HistorialContrasena" 
		SET contrasena_hash=$1, activo=$2, fecha_modificacion=now() 
		WHERE id_historial=$3`,
		h.ContrasenaHash, h.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Historial actualizado correctamente"})
}

// DELETE historial (baja lógica)
func DeleteHistorialContraseña(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."HistorialContrasena" SET activo=FALSE, fecha_modificacion=now() WHERE id_historial=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Historial eliminado (baja lógica)"})
}
