package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	e "github.com/Fact0RR/AVITO/API/entity"

)

func (s *APIserver) UserBannerHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Debugln("Обработчик получения баннеров для пользователя")
	w.Header().Set("Content-Type", "application/json")
	qp := r.URL.Query()
	p := e.UserParams{}
	tag_id, _ := strconv.Atoi(qp.Get("tag_id"))
	p.Tag_id = tag_id
	feature_id,_:=strconv.Atoi(qp.Get("feature_id"))
	p.Feature_id = feature_id
	
	s.Logger.Debugln("Проверка присутствия параметра use_last_revision")

	if qp.Has("use_last_revision"){
		s.Logger.Debugln("use_last_revision присутствует")
		p.UseLastRevision=qp.Get("use_last_revision")=="true"
	}

	var banner *e.UserBanner
	if p.UseLastRevision{
		s.Logger.Debugln("Получение данных из БД")
		banner = s.Store.GetUserBannerFromDB(p.Tag_id,p.Feature_id)
	}else{
		s.Logger.Debugln("Получение данных из ОЗУ")
		banner = s.Store.GetUserBanner(p.Tag_id,p.Feature_id)
	}
	if *banner==(e.UserBanner{}){
		s.Logger.Debugln("Баннеры с такими свойствами отсутствуют")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(banner)
    if err != nil {
		//s.Logger.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e.QuickErrToJson(err.Error()))
        return
    }
	s.Logger.Debugln("Успешная отправка списка баннеров")
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}