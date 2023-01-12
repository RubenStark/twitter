package middlew

import (
	"net/http"

	"github.com/RubenStark/twitter/bd"
)

// ChequeDB is the middleware that checks the connection to the database
func ChequeDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == false {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
