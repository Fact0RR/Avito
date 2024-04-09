package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/Fact0RR/AVITO/internal/middlewares"
)

func (s *APIserver) configureRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/user_banner", m.UserMiddleWare(s.Config.TokenAdmin, s.Config.TokenUser, m.ValidateUserBannerMiddleware(s.UserBannerHandler))).Methods(http.MethodGet)

	s.Router = router
}