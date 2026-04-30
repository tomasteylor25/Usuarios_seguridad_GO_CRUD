package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterHistorialContrasenaRoutes(r *mux.Router) {
	r.HandleFunc("/historial-contrasena", controllers.GetAllHistorialContraseña).Methods("GET")
	r.HandleFunc("/historial-contrasena/{id}", controllers.GetHistorialContraseñaByID).Methods("GET")
	r.HandleFunc("/historial-contrasena", controllers.CreateHistorialContraseña).Methods("POST")
	r.HandleFunc("/historial-contrasena/{id}", controllers.UpdateHistorialContraseña).Methods("PUT")
	r.HandleFunc("/historial-contrasena/{id}", controllers.DeleteHistorialContraseña).Methods("DELETE")
}
