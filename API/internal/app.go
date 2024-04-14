package internal

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

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
	level,err := logrus.ParseLevel(s.Config.LogLevel)
	if err != nil{
		return err
	}
	s.Logger.SetLevel(level)

	s.Logger.Info("Конфигурируем роутер")

	s.configureRouter()

	s.Logger.Info("открываем соединение с бд и запускаем горутину, которая скачивает данные из бд в ОЗУ раз в 5 мин")
	
	if err:= s.Store.Open();err != nil{
		return err
	}

	if s.Config.LoggingHandler{

		s.Logger.Info("Запуск сервера с логированием обработчиков")
		return http.ListenAndServe(s.Config.Port, handlers.LoggingHandler(os.Stdout, s.Router))
	}
	s.Logger.Info("Сервер запущен без логирования обработчиков")
	return http.ListenAndServe(s.Config.Port, s.Router)
}

func (s *APIserver) TryConnectToServer(maxWaitTime int,urlConn string)error{
	start := time.Now()
	for{
		duration := time.Since(start)
		if duration.Seconds() > float64(maxWaitTime) {
			return errors.New(("Не удается подключиться к серверу (время ожидания "+strconv.Itoa(int(duration.Seconds()))+" секунд)"))
		}
		_, err := http.Get(urlConn)
		if err != nil {
			time.Sleep(time.Millisecond*200)
			continue
		}
		break
	}
	return nil
}