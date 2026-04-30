package controllers

import (
	"Usuarios_seguridad_GO_CRUD/config"
	"Usuarios_seguridad_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL crear contraseñas
func GetAllCrearContraseña(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		`SELECT id_usuario, contrasena, confirmar_contrasena, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Crear_contrasena"`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.CrearContraseña
	for rows.Next() {
		var cc models.CrearContraseña
		rows.Scan(&cc.IDUsuario, &cc.Contrasena, &cc.ConfirmarContrasena, &cc.Activo, &cc.FechaCreacion, &cc.FechaModificacion)
		list = append(list, cc)
	}
	respondJSON(w, 200, list)
}

// GET BY ID crear contraseña
func GetCrearContraseñaByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var cc models.CrearContraseña

	err := config.DB.QueryRow(
		`SELECT id_usuario, contrasena, confirmar_contrasena, activo, fecha_creacion, fecha_modificacion 
		FROM "Usuario_seguridad"."Crear_contrasena" WHERE id_usuario=$1`, id,
	).Scan(&cc.IDUsuario, &cc.Contrasena, &cc.ConfirmarContrasena, &cc.Activo, &cc.FechaCreacion, &cc.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Registro no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, cc)
}

// POST crear contraseña
func CreateCrearContraseña(w http.ResponseWriter, r *http.Request) {
	var cc models.CrearContraseña
	if err := json.NewDecoder(r.Body).Decode(&cc); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO "Usuario_seguridad"."Crear_contrasena" (id_usuario, contrasena, confirmar_contrasena, activo) 
		VALUES ($1, $2, $3, $4) RETURNING fecha_creacion, fecha_modificacion`,
		cc.IDUsuario, cc.Contrasena, cc.ConfirmarContrasena, cc.Activo,
	).Scan(&cc.FechaCreacion, &cc.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, cc)
}

// PUT actualizar crear contraseña
func UpdateCrearContraseña(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var cc models.CrearContraseña
	json.NewDecoder(r.Body).Decode(&cc)

	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Crear_contrasena" 
		SET contrasena=$1, confirmar_contrasena=$2, activo=$3, fecha_modificacion=now() 
		WHERE id_usuario=$4`,
		cc.Contrasena, cc.ConfirmarContrasena, cc.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Contraseña actualizada correctamente"})
}

// DELETE crear contraseña (baja lógica)
func DeleteCrearContraseña(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(
		`UPDATE "Usuario_seguridad"."Crear_contrasena" SET activo=FALSE, fecha_modificacion=now() WHERE id_usuario=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Registro eliminado (baja lógica)"})
}
