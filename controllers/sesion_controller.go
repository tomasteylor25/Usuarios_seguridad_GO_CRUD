package controllers

import (
	"Usuarios_seguridad_GO_CRUD/config"
	"Usuarios_seguridad_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL sesiones
func GetAllSesiones(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		`SELECT id_sesion, id_usuario, token_sesion, ip_origen, user_agent, fecha_inicio, fecha_expiracion, revocada, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Sesion"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Sesion
	for rows.Next() {
		var s models.Sesion
		rows.Scan(&s.IDSesion, &s.IDUsuario, &s.TokenSesion, &s.IPOrigen, &s.UserAgent, &s.FechaInicio, &s.FechaExpiracion, &s.Revocada, &s.Activo, &s.FechaCreacion, &s.FechaModificacion)
		list = append(list, s)
	}
	respondJSON(w, 200, list)
}

// GET BY ID sesion
func GetSesionByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var s models.Sesion

	err := config.DB.QueryRow(
		`SELECT id_sesion, id_usuario, token_sesion, ip_origen, user_agent, fecha_inicio, fecha_expiracion, revocada, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Sesion" WHERE id_sesion=$1`, id,
	).Scan(&s.IDSesion, &s.IDUsuario, &s.TokenSesion, &s.IPOrigen, &s.UserAgent, &s.FechaInicio, &s.FechaExpiracion, &s.Revocada, &s.Activo, &s.FechaCreacion, &s.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Sesión no encontrada"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, s)
}

// POST crear sesion
func CreateSesion(w http.ResponseWriter, r *http.Request) {
	var s models.Sesion
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO "Usuario_seguridad"."Sesion" (id_usuario, token_sesion, ip_origen, user_agent, fecha_expiracion, revocada, activo) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_sesion, fecha_inicio, fecha_creacion, fecha_modificacion`,
		s.IDUsuario, s.TokenSesion, s.IPOrigen, s.UserAgent, s.FechaExpiracion, s.Revocada, s.Activo,
	).Scan(&s.IDSesion, &s.FechaInicio, &s.FechaCreacion, &s.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, s)
}

// PUT actualizar sesion
func UpdateSesion(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var s models.Sesion
	json.NewDecoder(r.Body).Decode(&s)

	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Sesion" 
		SET revocada=$1, activo=$2, fecha_modificacion=now() 
		WHERE id_sesion=$3`,
		s.Revocada, s.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Sesión actualizada correctamente"})
}

// DELETE sesion (baja lógica)
func DeleteSesion(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Sesion" SET activo=FALSE, revocada=TRUE, fecha_modificacion=now() WHERE id_sesion=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Sesión cerrada (baja lógica)"})
}
