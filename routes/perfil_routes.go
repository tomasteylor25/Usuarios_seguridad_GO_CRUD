package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterPerfilRoutes(r *mux.Router) {
	r.HandleFunc("/perfiles", controllers.GetAllPerfiles).Methods("GET")
	r.HandleFunc("/perfiles/{id}", controllers.GetPerfilByID).Methods("GET")
	r.HandleFunc("/perfiles", controllers.CreatePerfil).Methods("POST")
	r.HandleFunc("/perfiles/{id}", controllers.UpdatePerfil).Methods("PUT")
	r.HandleFunc("/perfiles/{id}", controllers.DeletePerfil).Methods("DELETE")
}
