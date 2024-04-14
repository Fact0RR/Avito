package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	e "github.com/Fact0RR/AVITO/entity"
	"github.com/gorilla/mux"
)

func (s *APIserver) BannerIdHandlerPatch(w http.ResponseWriter, r *http.Request) {
	s.Logger.Debugln("Обработчик обновления баннеров")
	var pb e.PostBanner
	params := mux.Vars(r)
	id,err :=strconv.Atoi(params["id"])
	if err !=nil{
		s.Logger.Warningln(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
        return
	}
	b,err := io.ReadAll(r.Body)
	if err !=nil{
		s.Logger.Warningln(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
        return
	}
	if err = json.Unmarshal(b,&pb);err !=nil{
		s.Logger.Warningln(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
        return
	}
	s.Logger.Debugln("Проверка на отсутствие пустых полей")
	str,empty := pb.IsWithEmptyField()
	if empty{
		s.Logger.Warningln(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Отстуствует поле: "+str))
		return 
	}
	s.Logger.Debugln("Проверка на существование баннера с id = ",id)
	if !s.Store.IsExistBanner(id){
		w.WriteHeader(http.StatusNotFound)
		return
	}
	s.Logger.Debugln("Запрос в бд на изменение")
	err = s.Store.PatchBannerToDB(id,&pb) 
	if err !=nil{
		s.Logger.Errorln(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e.QuickErrToJson(err.Error()))
        return
	}
	s.Logger.Debugln("Успешное изменение")
	w.WriteHeader(http.StatusOK)
}