package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RubenStark/twitter/bd"
	"github.com/RubenStark/twitter/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request){
	var mensaje models.GraboTweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID: IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}