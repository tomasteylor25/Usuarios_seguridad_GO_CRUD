package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterCredencialesRoutes(r *mux.Router) {
	r.HandleFunc("/credenciales", controllers.GetAllCredenciales).Methods("GET")
	r.HandleFunc("/credenciales/{id}", controllers.GetCredencialesByID).Methods("GET")
	r.HandleFunc("/credenciales", controllers.CreateCredenciales).Methods("POST")
	r.HandleFunc("/credenciales/{id}", controllers.UpdateCredenciales).Methods("PUT")
	r.HandleFunc("/credenciales/{id}", controllers.DeleteCredenciales).Methods("DELETE")
}
