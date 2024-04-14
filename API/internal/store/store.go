package store

import (
	"database/sql"
	"errors"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type Store struct {
	Connection *string
	DB         *sql.DB
	Mutex      *sync.Mutex
	LocalDB    *AllBannersData
}

func New(connection string) *Store {

	abd := &AllBannersData{}
	abd.BannersMap = make(map[int]*Banners)
	abd.B_Tmap = make(map[int]*B_T)
	abd.TagsMap = make(map[int]*Tags)
	abd.FeaturesMap = make(map[int]*Features)
	abd.TagsRefOnBannersMap = make(map[int][]int)
	store := Store{
		Connection: &connection,
		Mutex:      &sync.Mutex{},
		LocalDB:    abd,
	}
	return &store
}

func (s *Store) Open() error {
	err := TryConnectToDB(15,*s.Connection)
	if err != nil{
		return err
	}

	db, err := sql.Open("postgres", *s.Connection)
	if err != nil {
		return err
	}
	
	s.DB = db

	go s.startUpdateLocalData()

	return nil
}

func (s *Store) startUpdateLocalData() {
	s.GetAllDataAboutBannersToLocal()
	for {
		time.Sleep(time.Minute * 5)
		s.GetAllDataAboutBannersToLocal()
	}

}

func (s *Store) Close() {
	s.DB.Close()
}

func TryConnectToDB(waitTime int,urlConn string)error{
	start := time.Now()
	for{
		duration := time.Since(start)
		if duration.Seconds() > float64(waitTime) {
			return errors.New(("Не удается подключиться к бд (время ожидания "+strconv.Itoa(int(duration.Seconds()))+" секунд)"))
		}
		db, err := sql.Open("postgres", urlConn)
		if err != nil {
			time.Sleep(time.Millisecond*200)
			continue
		}
		if err := db.Ping(); err != nil {
			time.Sleep(time.Millisecond*200)
			continue
		}
		break
	}
	return nil
}