package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	e "github.com/Fact0RR/AVITO/entity"
)


func (s *APIserver) BannerHandlerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qp := r.URL.Query()
	p := e.AdminParams{}

	if qp.Has("tag_id"){
		p.Tag_id.Int,_=strconv.Atoi(qp.Get("tag_id"))
		p.Tag_id.Exist = true
	}
	if qp.Has("feature_id"){
		p.Feature_id.Int,_=strconv.Atoi(qp.Get("feature_id"))
		p.Feature_id.Exist = true
	}
	if qp.Has("limit"){
		p.Limit.Int,_=strconv.Atoi(qp.Get("tag_id"))
		p.Limit.Exist = true
	}
	if qp.Has("offset"){
		p.Offset.Int,_=strconv.Atoi(qp.Get("offset"))
		p.Offset.Exist = true
	}
	var abs []e.AdminBanner
	switch {
	case !p.Tag_id.Exist && !p.Feature_id.Exist:
		limit,offset := getQueryLimitAndOffset(p.Limit,p.Offset)
		query := "select b.id,b.feature_id, b.title, b.text, b.url, b.visible, b.create_time, b.update_time from banners b join b_t on b_t.banner_id = b.id "+limit+" "+offset

		absl,err := s.Store.GetAdminBannerFromDB(query)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}
		abs = absl
	case p.Tag_id.Exist && !p.Feature_id.Exist:

		limit,offset := getQueryLimitAndOffset(p.Limit,p.Offset)
		query := "select b.id, b.feature_id, b.title, b.text, b.url, b.visible, b.create_time, b.update_time from banners b join b_t on b_t.banner_id = b.id where b_t.tag_id = "+strconv.Itoa(p.Tag_id.Int)+" "+limit+" "+offset

		absl,err := s.Store.GetAdminBannerFromDB(query)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return 
		}
		abs = absl
	case !p.Tag_id.Exist && p.Feature_id.Exist:

		limit,offset := getQueryLimitAndOffset(p.Limit,p.Offset)
		query := "select b.id,b.feature_id, b.title, b.text, b.url, b.visible, b.create_time, b.update_time from banners b where  b.feature_id = "+strconv.Itoa(p.Feature_id.Int)+" "+limit+" "+offset

		absl,err := s.Store.GetAdminBannerFromDB(query)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return 
		}
		abs = absl

	case p.Tag_id.Exist && p.Feature_id.Exist:
		
		limit,offset := getQueryLimitAndOffset(p.Limit,p.Offset)
		query := "select b.id,b.feature_id, b.title, b.text, b.url, b.visible, b.create_time, b.update_time from banners b join b_t on b_t.banner_id = b.id where b_t.tag_id = "+strconv.Itoa(p.Tag_id.Int)+" and b.feature_id = "+strconv.Itoa(p.Feature_id.Int)+" "+limit+" "+offset

		absl,err := s.Store.GetAdminBannerFromDB(query)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return 
		}
		abs = absl
	}

	if len(abs)==0{
		w.WriteHeader(http.StatusNotFound)
        return
	}

	b, err := json.Marshal(abs)
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
        return
    }
	w.Write(b)
	w.WriteHeader(http.StatusOK)

}


func getQueryLimitAndOffset(lim,off e.ExistInt)(string,string){
	switch{
	case !lim.Exist && !off.Exist:
		return "",""
	case lim.Exist && !off.Exist:
		return " limit "+strconv.Itoa(lim.Int),""
	case !lim.Exist && off.Exist:
		return "", "offset "  + strconv.Itoa(off.Int)
	case lim.Exist && off.Exist:
		return " limit "+strconv.Itoa(lim.Int)," offset " + strconv.Itoa(off.Int)
	}
	return "",""
	
}