package middlewares

import (
	"net/http"

	e "github.com/Fact0RR/AVITO/entity"
)

// req_qp_Arr := []string{"tag_id","feature_id"}
// noReq_QP_Arr := []string{"use_last_revision","offset","limit"}
func ValidateUserBannerMiddleware(pvsGlobal []e.PValidate, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		qp := r.URL.Query()
		pvs := e.CopySlice(pvsGlobal)
		e.Fill(qp, pvs)
		for _, pv := range pvs {
			if !pv.CheckRequired() {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Отсутствуют обязательные параметры в url: " + pv.Name))
				return
			}
			if !pv.CheckType() {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Типы параметров не соблюдены: " + pv.Name))
				return
			}
		}
		next(w, r)
	}
}



