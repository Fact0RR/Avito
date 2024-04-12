package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	e "github.com/Fact0RR/AVITO/API/entity"
)


func (s *APIserver) BannerHandlerGet(w http.ResponseWriter, r *http.Request) {
	s.Logger.Debugln("Обработчик получения баннеров для админом")
	w.Header().Set("Content-Type", "application/json")
	qp := r.URL.Query()
	p := e.AdminParams{}

	s.Logger.Debugln("Проверка url параметров на наличие")

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
		s.Logger.Debugln("Нет tag_id и feature_id")
		limit,offset := getQueryLimitAndOffset(p.Limit,p.Offset)
		query := "select b.id,b.feature_id, b.title, b.text, b.url, b.visible, b.create_time, b.update_time from banners b join b_t on b_t.banner_id = b.id "+limit+" "+offset

		absl,err := s.Store.GetAdminBannerFromDB(query)
		if err != nil{
			s.Logger.Errorln(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(e.QuickErrToJson(err.Error()))
			return 
		}
		abs = absl
	case p.Tag_id.Exist && !p.Feature_id.Exist:
		s.Logger.Debugln("Есть tag_id и нет feature_id")
		limit,offset := getQueryLimitAndOffset(p.Limit,p.Offset)
		query := "select b.id, b.feature_id, b.title, b.text, b.url, b.visible, b.create_time, b.update_time from banners b join b_t on b_t.banner_id = b.id where b_t.tag_id = "+strconv.Itoa(p.Tag_id.Int)+" "+limit+" "+offset

		absl,err := s.Store.GetAdminBannerFromDB(query)
		if err != nil{
			s.Logger.Errorln(err.Error())

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(e.QuickErrToJson(err.Error()))
			return 
		}
		abs = absl
	case !p.Tag_id.Exist && p.Feature_id.Exist:
		s.Logger.Debugln("Нет tag_id, но есть feature_id")
		limit,offset := getQueryLimitAndOffset(p.Limit,p.Offset)
		query := "select b.id,b.feature_id, b.title, b.text, b.url, b.visible, b.create_time, b.update_time from banners b where  b.feature_id = "+strconv.Itoa(p.Feature_id.Int)+" "+limit+" "+offset

		absl,err := s.Store.GetAdminBannerFromDB(query)
		if err != nil{
			s.Logger.Errorln(err.Error())

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(e.QuickErrToJson(err.Error()))
			return 
		}
		abs = absl

	case p.Tag_id.Exist && p.Feature_id.Exist:

		s.Logger.Debugln("Есть tag_id и feature_id")

		limit,offset := getQueryLimitAndOffset(p.Limit,p.Offset)
		query := "select b.id,b.feature_id, b.title, b.text, b.url, b.visible, b.create_time, b.update_time from banners b join b_t on b_t.banner_id = b.id where b_t.tag_id = "+strconv.Itoa(p.Tag_id.Int)+" and b.feature_id = "+strconv.Itoa(p.Feature_id.Int)+" "+limit+" "+offset

		s.Logger.Debugln("Запрос в бд на получение баннеров")

		absl,err := s.Store.GetAdminBannerFromDB(query)
		if err != nil{
			s.Logger.Errorln(err.Error())

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(e.QuickErrToJson(err.Error()))
			return 
		}
		abs = absl
	}

	if len(abs)==0{
		s.Logger.Debugln("Баннеров с такими свойствами отсутствуют")
		w.WriteHeader(http.StatusNotFound)
        return
	}

	b, err := json.Marshal(abs)
    if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e.QuickErrToJson(err.Error()))
        return
    }
	s.Logger.Debugln("Успешная отправка списка баннеров")
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