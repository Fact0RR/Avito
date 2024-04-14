package internal

import (
	"net/http"
	"strconv"

	e "github.com/Fact0RR/AVITO/entity"
	"github.com/gorilla/mux"
)

func (s *APIserver) BannerIdHandlerDelete(w http.ResponseWriter, r *http.Request) {
	s.Logger.Debugln("Обработчик удаления баннеров")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		s.Logger.Warningln(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(e.QuickErrToJson(err.Error()))
		return
	}
	
	s.Logger.Debugln("Проверка на существование баннера с id = ",id)

	if !s.Store.IsExistBanner(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.Logger.Debugln("Запрос в базу данных на удаление")

	err = s.Store.DeleteBannerToDB(id)
	if err != nil {
		s.Logger.Errorln(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e.QuickErrToJson(err.Error()))
		return
	}
	s.Logger.Debugln("Успешное удаление")
	w.WriteHeader(http.StatusNoContent)
}