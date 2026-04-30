package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterIntentoLoginRoutes(r *mux.Router) {
	r.HandleFunc("/intentos-login", controllers.GetAllIntentosLogin).Methods("GET")
	r.HandleFunc("/intentos-login/{id}", controllers.GetIntentoLoginByID).Methods("GET")
	r.HandleFunc("/intentos-login", controllers.CreateIntentoLogin).Methods("POST")
	r.HandleFunc("/intentos-login/{id}", controllers.UpdateIntentoLogin).Methods("PUT")
	r.HandleFunc("/intentos-login/{id}", controllers.DeleteIntentoLogin).Methods("DELETE")
}
