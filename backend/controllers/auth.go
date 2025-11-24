package controllers

import (
	"encoding/json"
	"net/http"

	"safebase/db"
	"safebase/models"
	"safebase/utils"
	"safebase/validations"

	"golang.org/x/crypto/bcrypt"
)

// ------ INSCRIPTION UTILISATEUR ------
func Register(w http.ResponseWriter, r *http.Request) {
	var req models.User
	// Décodage du JSON envoyé par le client
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Format JSON invalide")
		return
	}

	// ------VALIDATION------
	if err := validations.ValidatePassword(req.Password); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validations.ValidateEmail(req.Email); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validations.ValidateUsername(req.Username); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Hash du mot de passe
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Erreur interne du serveur")
		return
	}
	req.Password = string(hashed)

	// Insertion dans la base des données
	_, err = db.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		req.Username, req.Email, req.Password,
	)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "L'adresse e-mail existe déjà dans la base des données")
		return
	}

	// Response JSON pour confirmation
	utils.SendSuccess(w, http.StatusCreated, "Utilisateur enregistré avec succès!")
}