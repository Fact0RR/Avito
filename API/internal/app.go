package internal

import (
	"net/http"
	"os"

	"github.com/Fact0RR/AVITO/config"
	"github.com/Fact0RR/AVITO/internal/store"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	Config *config.Config
	Router *mux.Router
	Store *store.Store
	Logger *logrus.Logger
}


func New(config *config.Config) *APIserver {
	return &APIserver{
		Config: config,
		Logger: logrus.New(),
		Store: store.New(config.DataBaseString),
	}
}

func (s *APIserver) Start() error {

	s.configureRouter()

	//открываем соединение с бд
	if err:= s.Store.Open();err != nil{
		return err
	}

	
	level,err := logrus.ParseLevel(s.Config.LogLevel)
	if err != nil{
		return err
	}
	s.Logger.SetLevel(level)

	s.Logger.Info("Запуск сервера")

	return http.ListenAndServe(s.Config.Port, handlers.LoggingHandler(os.Stdout, s.Router))
}