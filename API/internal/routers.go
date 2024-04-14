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
	setAdminParams()
	router := mux.NewRouter()
	router.HandleFunc("/user_banner", m.UserMiddleWare(s.Config.TokenAdmin, s.Config.TokenUser, m.ValidateBannerMiddleware(parametrsForUser,s.UserBannerHandler))).Methods(http.MethodGet)
	router.HandleFunc("/banner",m.AdminMiddleWare(s.Config.TokenAdmin, s.Config.TokenUser,m.ValidateBannerMiddleware(parametrsForAdmin,s.BannerHandlerGet))).Methods(http.MethodGet)
	router.HandleFunc("/banner",m.AdminMiddleWare(s.Config.TokenAdmin, s.Config.TokenUser,s.BannerHandlerPost)).Methods(http.MethodPost)
	router.HandleFunc("/banner/{id:[0-9]+}",m.AdminMiddleWare(s.Config.TokenAdmin, s.Config.TokenUser,s.BannerIdHandlerPatch)).Methods(http.MethodPatch)
	router.HandleFunc("/banner/{id:[0-9]+}",m.AdminMiddleWare(s.Config.TokenAdmin, s.Config.TokenUser,s.BannerIdHandlerDelete)).Methods(http.MethodDelete)
	router.HandleFunc("/stress",s.StressHandler)

	s.Router = router
}

func setUserParms(){
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "tag_id", PType: "int",Required: true})
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "feature_id", PType: "int",Required: true})
	parametrsForUser = append(parametrsForUser,e.PValidate{Name: "use_last_revision", PType: "bool"})
}

func setAdminParams(){
	parametrsForAdmin = append(parametrsForAdmin,e.PValidate{Name: "tag_id", PType: "int"})
	parametrsForAdmin = append(parametrsForAdmin,e.PValidate{Name: "feature_id", PType: "int"})
	parametrsForAdmin = append(parametrsForAdmin,e.PValidate{Name: "limit", PType: "int"})
	parametrsForAdmin = append(parametrsForAdmin,e.PValidate{Name: "offset", PType: "int"})
}
