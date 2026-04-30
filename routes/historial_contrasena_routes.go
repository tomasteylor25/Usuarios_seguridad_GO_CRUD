package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterHistorialContrasenaRoutes(r *mux.Router) {
	r.HandleFunc("/historial-contrasena", controllers.GetAllHistorialContrasena).Methods("GET")
	r.HandleFunc("/historial-contrasena/{id}", controllers.GetHistorialContrasenaByID).Methods("GET")
	r.HandleFunc("/historial-contrasena", controllers.CreateHistorialContrasena).Methods("POST")
	r.HandleFunc("/historial-contrasena/{id}", controllers.UpdateHistorialContrasena).Methods("PUT")
	r.HandleFunc("/historial-contrasena/{id}", controllers.DeleteHistorialContrasena).Methods("DELETE")
}
