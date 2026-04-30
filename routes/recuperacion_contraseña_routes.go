package routes

import (
	"Usuarios_seguridad_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterRecuperacionContrasenaRoutes(r *mux.Router) {
	r.HandleFunc("/recuperacion-contrasena", controllers.GetAllRecuperacionContraseña).Methods("GET")
	r.HandleFunc("/recuperacion-contrasena/{id}", controllers.GetRecuperacionContraseñaByID).Methods("GET")
	r.HandleFunc("/recuperacion-contrasena", controllers.CreateRecuperacionContraseña).Methods("POST")
	r.HandleFunc("/recuperacion-contrasena/{id}", controllers.UpdateRecuperacionContraseña).Methods("PUT")
	r.HandleFunc("/recuperacion-contrasena/{id}", controllers.DeleteRecuperacionContraseña).Methods("DELETE")
}
