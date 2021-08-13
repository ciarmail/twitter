package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ciarmail/twitter/bd"
	"github.com/ciarmail/twitter/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Error, el email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Error, contraseña de mínimo 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Error, ya existe un usuario con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error, al registrar al usuario "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Error, no se pudo registrar al usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
