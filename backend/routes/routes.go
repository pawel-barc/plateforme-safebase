package routes

import (
	"net/http"

	"safebase/middleware"

	"github.com/go-chi/chi/v5"
)

// SetupRouter configure toutes les routes de l'application
func SetupRouter() http.Handler {
	// Création du routeur principal(chi framework comme express)
	r := chi.NewRouter()

	// Activation du middleware CORS
	r.Use(middleware.CORSHandler())

	// -----ROUTES-----//

	
	// Retourne le routeur configuré comme 'http.Handler'
	return r
}