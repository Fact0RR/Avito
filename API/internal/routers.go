package internal

import (
	"net/http"

	e "github.com/Fact0RR/AVITO/entity"
	m "github.com/Fact0RR/AVITO/internal/middlewares"
	"github.com/gorilla/mux"
)

var parametrsForUser []e.PValidate 
var parametrsForAdmin []e.PValidate

func (s *APIserver) configureRouter() {
	setUserParms()
	

	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/user_banner", m.UserMiddleWare(s.Config.TokenAdmin, s.Config.TokenUser, m.ValidateUserBannerMiddleware(parametrsForUser,s.UserBannerHandler))).Methods(http.MethodGet)
	router.HandleFunc("/banner",m.AdminMiddleWare(s.Config.TokenAdmin, s.Config.TokenUser,m.ValidateUserBannerMiddleware(parametrsForAdmin,s.BannerHandlerGet))).Methods(http.MethodGet)

	s.Router = router
}

func setUserParms(){
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "tag_id", PType: "int",Required: true})
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "feature_id", PType: "int",Required: true})
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "use_last_revision", PType: "bool"})
}

func setAdmiParams(){
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "tag_id", PType: "int"})
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "feature_id", PType: "int"})
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "limit", PType: "int"})
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "offset", PType: "int"})
}