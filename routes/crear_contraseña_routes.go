package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterCrearContrasenaRoutes(r *mux.Router) {
	r.HandleFunc("/crear-contrasena", controllers.GetAllCrearContrasena).Methods("GET")
	r.HandleFunc("/crear-contrasena/{id}", controllers.GetCrearContrasenaByID).Methods("GET")
	r.HandleFunc("/crear-contrasena", controllers.CreateCrearContrasena).Methods("POST")
	r.HandleFunc("/crear-contrasena/{id}", controllers.UpdateCrearContrasena).Methods("PUT")
	r.HandleFunc("/crear-contrasena/{id}", controllers.DeleteCrearContrasena).Methods("DELETE")
}
