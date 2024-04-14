package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	e "github.com/Fact0RR/AVITO/entity"
)


func (s *APIserver) BannerHandlerPost(w http.ResponseWriter, r *http.Request) {
	s.Logger.Debugln("Обработчик добавления баннеров")
	var pb e.PostBanner
	b,err := io.ReadAll(r.Body)
	if err!=nil{
		s.Logger.Warningln(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(e.QuickErrToJson(err.Error()))
        return
	}
	if err = json.Unmarshal(b,&pb);err !=nil{
		s.Logger.Errorln(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e.QuickErrToJson(err.Error()))
        return
	}

	s.Logger.Debugln("Проверка на пустые поля")

	str,empty := pb.IsWithEmptyField()
	if empty{
		s.Logger.Warningln(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(e.QuickErrToJson("Отстуствует поле: "+str))
		return 
	}

	s.Logger.Debugln("Запрос в бд на добавление")

	banner_id, err := s.Store.PostBannerToDB(&pb)
	if err !=nil{
		s.Logger.Errorln(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e.QuickErrToJson(err.Error()))
        return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	s.Logger.Debugln("Успешное добавление в бд")
	w.Write([]byte("{\"banner_id\":"+strconv.Itoa(banner_id)+"}"))
	
}