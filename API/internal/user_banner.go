package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Fact0RR/AVITO/internal/middlewares"
	"github.com/Fact0RR/AVITO/internal/store"
)

func (s *APIserver) UserBannerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qp := r.URL.Query()
	p := middlewares.Params{}
	tag_id, _ := strconv.Atoi(qp.Get("tag_id"))
	p.Tag_id = tag_id
	feature_id,_:=strconv.Atoi(qp.Get("feature_id"))
	p.Feature_id = feature_id
	
	if qp.Has("use_last_revision"){
		p.UseLastRevision=qp.Get("use_last_revision")=="true"
	}
	//if p.UseLastRevision == true{
		//TODO 
	//}
	banner := s.Store.GetUserBanner(p.Tag_id,p.Feature_id)
	
	if *banner==(store.UserBanner{}){
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(banner)
    if err != nil {
        return
    }
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}