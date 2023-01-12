package models

// RespuestaConsultaRelacion es la respuesta que se devuelve al consultar la relacion entre dos usuarios
type RespuestaConsultaRelacion struct {
	Status bool `json:"status,omitempty"`
}
