package main

import (
	"fmt"
	"log"
	"net/http"
	"personas/controller"
	"personas/utils"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/api/persona", controller.CreatePersona).Methods(http.MethodPost)
	router.HandleFunc("/api/persona", controller.GetPersona).Methods(http.MethodGet)
	router.HandleFunc("/api/persona", controller.UpdatePersona).Methods(http.MethodPut)
	router.HandleFunc("/api/persona/{identificador}", controller.DeletePersona).Methods(http.MethodDelete)
	router.HandleFunc("/api/persona/html", controller.GetPersonasHTML).Methods(http.MethodGet)

	router.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./view/swagger.yaml") })
	router.PathPrefix("/console").Handler(http.StripPrefix("/console/", http.FileServer(http.Dir("./view/api"))))

	log.Printf("Servidor escuchando por el puerto %s\n", utils.GetConfig().ApiServerPort)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", utils.GetConfig().ApiServerPort), router))
}

func init() {
	utils.InitConf()
	utils.GetConnection()
}
