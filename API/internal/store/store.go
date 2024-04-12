package store

import (
	"database/sql"
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
	db, err := sql.Open("postgres", *s.Connection)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.DB = db

	go s.startUpdateLocalData()

	return nil
}

func (s *Store) startUpdateLocalData() {
	for {
		s.getAllDataAboutBannersToLocal()
		time.Sleep(time.Minute * 5)
	}

}

func (s *Store) Close() {
	s.DB.Close()
}
