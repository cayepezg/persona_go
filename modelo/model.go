package modelo

// Persona
//
// Representación de la entidad Persona.
//
// swagger:model Persona
type Persona struct {
	ID            *int   `json:"id,omitempty"`
	Identificador string `json:"identificador,omitempty"`
	Nombre        string `json:"nombre,omitempty"`
	Apellido      string `json:"apellido,omitempty"`
	Sexo          string `json:"sexo,omitempty"`
}

// Persona Mapea el nombre relacional, de la entidad.
func (Persona) TableName() string {
	return "persona"
}
