package controllers

import (
	"Usuarios_seguridad_GO_CRUD/config"
	"Usuarios_seguridad_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL credenciales
func GetAllCredenciales(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		`SELECT id_usuario, contrasena_hash, intentos_fallidos, bloqueado, ultimo_login, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Credenciales"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Credenciales
	for rows.Next() {
		var c models.Credenciales
		rows.Scan(&c.IDUsuario, &c.ContrasenaHash, &c.IntentosFallidos, &c.Bloqueado, &c.UltimoLogin, &c.Activo, &c.FechaCreacion, &c.FechaModificacion)
		list = append(list, c)
	}
	respondJSON(w, 200, list)
}

// GET BY ID credencial
func GetCredencialesByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var c models.Credenciales

	err := config.DB.QueryRow(
		`SELECT id_usuario, contrasena_hash, intentos_fallidos, bloqueado, ultimo_login, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Credenciales" WHERE id_usuario=$1`, id,
	).Scan(&c.IDUsuario, &c.ContrasenaHash, &c.IntentosFallidos, &c.Bloqueado, &c.UltimoLogin, &c.Activo, &c.FechaCreacion, &c.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Credencial no encontrada"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, c)
}

// POST crear credencial
func CreateCredenciales(w http.ResponseWriter, r *http.Request) {
	var c models.Credenciales
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO "Usuario_seguridad"."Credenciales" (id_usuario, contrasena_hash, intentos_fallidos, bloqueado, activo) 
		VALUES ($1, $2, $3, $4, $5) RETURNING fecha_creacion, fecha_modificacion`,
		c.IDUsuario, c.ContrasenaHash, c.IntentosFallidos, c.Bloqueado, c.Activo,
	).Scan(&c.FechaCreacion, &c.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, c)
}

// PUT actualizar credencial
func UpdateCredenciales(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var c models.Credenciales
	json.NewDecoder(r.Body).Decode(&c)

	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Credenciales" 
		SET contrasena_hash=$1, intentos_fallidos=$2, bloqueado=$3, activo=$4, fecha_modificacion=now() 
		WHERE id_usuario=$5`,
		c.ContrasenaHash, c.IntentosFallidos, c.Bloqueado, c.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Credencial actualizada correctamente"})
}

// DELETE credencial (baja lógica)
func DeleteCredenciales(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Credenciales" SET activo=FALSE, fecha_modificacion=now() WHERE id_usuario=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Credencial eliminada (baja lógica)"})
}
