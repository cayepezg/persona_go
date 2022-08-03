package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"personas/modelo"
	"personas/servicio"

	"github.com/gorilla/mux"
)

// swagger:operation GET /api/persona Persona GetPersona
// ---
// summary: Entrega la Persona solicitada.
// description: Entrega la Persona solicitada.
// tags:
// - Persona
// parameters:
// - name: identificador
//   in: query
//   description: Identificador único de la persona.
//   type: string
//   required: false
// responses:
//   200:
//     description: Data de la Persona consultada
//     schema:
//       $ref: '#/definitions/Persona'
func GetPersona(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	identificador := r.URL.Query().Get("identificador")
	persona, err := servicio.GetPersona(identificador)

	if err != nil {
		log.Printf("Error consultando persona. %s", err.Error())
		if err.Error() == "404" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persona)

}

// swagger:operation POST /api/persona Persona CreatePersona
// ---
// summary: Crea una persona.
// description: Crea una persona.
// tags:
// - Persona
// parameters:
// - name: persona
//   in: body
//   schema:
//	   $ref: '#/definitions/Persona'
//   description: Persona a Crear.
//   required: true
// responses:
//   200:
//     description: Data de la Persona creada
//     schema:
//       $ref: '#/definitions/Persona'
func CreatePersona(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	persona := modelo.Persona{}

	err := json.NewDecoder(r.Body).Decode(&persona)

	if err != nil {
		log.Println("Un problema extrayendo del body a la persona.. Revisar y pasar error como corresponde")
	}

	persona, err = servicio.CreatePersona(persona)

	if err != nil {
		log.Printf("Error creando persona. %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persona)

}

// swagger:operation PUT /api/persona Persona UpdatePersona
// ---
// summary: Actualiza una persona.
// description: Actualiza una persona.
// tags:
// - Persona
// parameters:
// - name: persona
//   in: body
//   description: Persona a Actualizar.
//   schema:
//	   $ref: '#/definitions/Persona'
//   required: true
// responses:
//   200:
//     description: Data de la Persona actualizada
//     schema:
//       $ref: '#/definitions/Persona'
func UpdatePersona(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	persona := modelo.Persona{}

	err := json.NewDecoder(r.Body).Decode(&persona)

	if err != nil {
		log.Println("Un problema extrayendo del body a la persona.. Revisar y pasar error como corresponde")
	}

	persona, err = servicio.UpdatePersona(persona)

	if err != nil {
		log.Printf("Error actualizando persona. %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persona)

}

// swagger:operation DELETE /api/persona/{identificador} Persona DeletePersona
// ---
// summary: Elimina una persona.
// description: Elimina una persona.
// tags:
// - Persona
// parameters:
// - name: identificador
//   in: path
//   description: Identificador de la persona a eliminar.
//   type: Persona
//   required: true
// responses:
//   200:
func DeletePersona(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	err := servicio.DeletePersona(vars["identificador"])

	if err != nil {
		log.Printf("Error eliminando a la persona con identificador: %s", vars["identificador"])
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// swagger:operation GET /api/persona/html Persona GetPersonasHTML
// ---
// summary: Entrega un listado de personas en formato HTML.
// description: Entrega un listado de Personas en formato HTML.
// tags:
// - Persona
// responses:
//   200:
//     description: Data de la Persona consultada
//     schema:
func GetPersonasHTML(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	html, _ := servicio.GetPersonasHTML()

	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, html)

}
