package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterCrearContrasenaRoutes(r *mux.Router) {
	r.HandleFunc("/crear-contrasena", controllers.GetAllCrearContraseña).Methods("GET")
	r.HandleFunc("/crear-contrasena/{id}", controllers.GetCrearContraseñaByID).Methods("GET")
	r.HandleFunc("/crear-contrasena", controllers.CreateCrearContraseña).Methods("POST")
	r.HandleFunc("/crear-contrasena/{id}", controllers.UpdateCrearContraseña).Methods("PUT")
	r.HandleFunc("/crear-contrasena/{id}", controllers.DeleteCrearContraseña).Methods("DELETE")
}
