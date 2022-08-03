package servicio

import (
	"errors"
	"fmt"
	"log"
	"personas/modelo"
	"personas/utils"
	"time"
)

// CreatePersona Crea en BD, una Persona.
func CreatePersona(persona modelo.Persona) (modelo.Persona, error) {

	c := utils.GetConnection() //.LogMode(true)

	result := c.Create(&persona)

	if result.Error != nil {
		log.Println("Error Creando Persona", result.Error.Error())
		return persona, result.Error
	}

	return persona, nil

}

// GetPersona Consulta en BD, una Persona a través de su cédula de identidad
func GetPersona(identificador string) (modelo.Persona, error) {

	persona := modelo.Persona{}
	c := utils.GetConnection() //.LogMode(true)

	result := c.Raw(`
		select * from persona
			where identificador = ?
		limit 1`,
		identificador)

	if result.Error != nil {
		log.Println("Error consultando Persona", result.Error.Error())
		return persona, result.Error
	}

	err := result.Scan(&persona)
	if err.Error != nil {
		log.Printf("error consultando vista de Persona. %s", err.Error)
		return persona, err.Error
	}

	if persona.ID == nil {
		return persona, errors.New("404")
	}

	return persona, nil

}

// UpdatePersona Modifica en BD, una Persona.
func UpdatePersona(persona modelo.Persona) (modelo.Persona, error) {

	c := utils.GetConnection() //.LogMode(true)

	result := c.Save(&persona)

	if result.Error != nil {
		log.Println("Error Modificando Persona", result.Error.Error())
		return persona, result.Error
	}

	return persona, nil

}

// DeletePersona Elimina en BD, una Persona.
func DeletePersona(identificador string) error {

	c := utils.GetConnection() //.LogMode(true)

	result := c.Where("identificador = ?", identificador).Delete(&modelo.Persona{})

	if result.Error != nil {
		log.Println("Error Eliminando Persona", result.Error.Error())
		return result.Error
	}

	return nil

}

// GetPersonasHTML Entrega un listado de personas en formato html.
func GetPersonasHTML() (string, error) {

	personas := []modelo.Persona{}
	html := fmt.Sprintf(`
		<h1> Consultado por Microservicio: %s </h1>
		<table> 
			<thead>
				<tr>
					<th>ID</th>
					<th>Identificador</th>
					<th>Nombre</th>
					<th>Apellido</th>
				</tr>
			</thead>`, time.Now().Format("2006-01-02 15:04:05"))

	c := utils.GetConnection() //.LogMode(true)

	result := c.Find(&personas)

	if result.Error != nil {
		return "", result.Error
	}

	for _, persona := range personas {
		html += fmt.Sprintf(`
			<tr>
				<td>
					%d
				</td>
				<td>
					%s
				</td>
				<td>
					%s
				</td>
				<td>
					%s
				</td>
			</tr>
		`, *persona.ID, persona.Identificador, persona.Nombre, persona.Apellido)
	}

	html += "</table>"

	return string(html), nil

}
