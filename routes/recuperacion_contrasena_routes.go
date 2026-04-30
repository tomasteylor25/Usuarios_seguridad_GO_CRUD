package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterRecuperacionContrasenaRoutes(r *mux.Router) {
	r.HandleFunc("/recuperacion-contrasena", controllers.GetAllRecuperacionContrasena).Methods("GET")
	r.HandleFunc("/recuperacion-contrasena/{id}", controllers.GetRecuperacionContrasenaByID).Methods("GET")
	r.HandleFunc("/recuperacion-contrasena", controllers.CreateRecuperacionContrasena).Methods("POST")
	r.HandleFunc("/recuperacion-contrasena/{id}", controllers.UpdateRecuperacionContrasena).Methods("PUT")
	r.HandleFunc("/recuperacion-contrasena/{id}", controllers.DeleteRecuperacionContrasena).Methods("DELETE")
}
