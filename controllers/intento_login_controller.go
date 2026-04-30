package controllers

import (
	"Usuarios_seguridad_GO_CRUD/config"
	"Usuarios_seguridad_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL intentos de login
func GetAllIntentosLogin(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		`SELECT id_intento, id_usuario, email_ingresado, exitoso, motivo_fallo, ip_origen, fecha, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."IntentoLogin"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.IntentoLogin
	for rows.Next() {
		var i models.IntentoLogin
		rows.Scan(&i.IDIntento, &i.IDUsuario, &i.EmailIngresado, &i.Exitoso, &i.MotivoFallo, &i.IPOrigen, &i.Fecha, &i.Activo, &i.FechaCreacion, &i.FechaModificacion)
		list = append(list, i)
	}
	respondJSON(w, 200, list)
}

// GET BY ID intento login
func GetIntentoLoginByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var i models.IntentoLogin

	err := config.DB.QueryRow(
		`SELECT id_intento, id_usuario, email_ingresado, exitoso, motivo_fallo, ip_origen, fecha, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."IntentoLogin" WHERE id_intento=$1`, id,
	).Scan(&i.IDIntento, &i.IDUsuario, &i.EmailIngresado, &i.Exitoso, &i.MotivoFallo, &i.IPOrigen, &i.Fecha, &i.Activo, &i.FechaCreacion, &i.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Intento de login no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, i)
}

// POST registrar intento login
func CreateIntentoLogin(w http.ResponseWriter, r *http.Request) {
	var i models.IntentoLogin
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO "Usuario_seguridad"."IntentoLogin" (id_usuario, email_ingresado, exitoso, motivo_fallo, ip_origen, activo) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_intento, fecha, fecha_creacion, fecha_modificacion`,
		i.IDUsuario, i.EmailIngresado, i.Exitoso, i.MotivoFallo, i.IPOrigen, i.Activo,
	).Scan(&i.IDIntento, &i.Fecha, &i.FechaCreacion, &i.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, i)
}

// PUT actualizar intento login
func UpdateIntentoLogin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var i models.IntentoLogin
	json.NewDecoder(r.Body).Decode(&i)

	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."IntentoLogin" 
		SET exitoso=$1, motivo_fallo=$2, activo=$3, fecha_modificacion=now() 
		WHERE id_intento=$4`,
		i.Exitoso, i.MotivoFallo, i.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Intento de login actualizado correctamente"})
}

// DELETE intento login (baja lógica)
func DeleteIntentoLogin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."IntentoLogin" SET activo=FALSE, fecha_modificacion=now() WHERE id_intento=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Intento de login eliminado (baja lógica)"})
}
