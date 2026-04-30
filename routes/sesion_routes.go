package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterSesionRoutes(r *mux.Router) {
	r.HandleFunc("/sesiones", controllers.GetAllSesiones).Methods("GET")
	r.HandleFunc("/sesiones/{id}", controllers.GetSesionByID).Methods("GET")
	r.HandleFunc("/sesiones", controllers.CreateSesion).Methods("POST")
	r.HandleFunc("/sesiones/{id}", controllers.UpdateSesion).Methods("PUT")
	r.HandleFunc("/sesiones/{id}", controllers.DeleteSesion).Methods("DELETE")
}
