package store

import (
	"sync"
	"time"
)

type Banners struct {
	Id         int
	Title      string
	Text       string
	Url        string
	Feature_id int
	Visible    bool
	CreateTime time.Time
	UpdateTime time.Time
}

type Features struct {
	Id int
	Name string
	Description string
}

type Tags struct {
	Id int
	Name string
	Description string 	
}

type B_T struct {
	Id        int
	Banner_id int
	Tag_id    int
}

type AllBannersData struct{
	BannersMap map[int]*Banners
	FeaturesMap map[int]*Features
	TagsMap map[int]*Tags
	B_Tmap map[int]*B_T
	TagsRefOnBannersMap map[int][]int
}

func(s *Store) GetAllDataAboutBannersToLocal() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func(){
		s.setAllBanners()
		defer wg.Done()
	}()
	go func(){
		s.setAllFeatures()
		defer wg.Done()
	}()
	go func(){
		s.setAllTags()
		defer wg.Done()
	}()
	go func(){
		s.setAllB_T()
		defer wg.Done()
	}()
	wg.Wait()

	for k,_ := range s.LocalDB.B_Tmap{
		delete(s.LocalDB.TagsRefOnBannersMap,k)
	}

	for _,v := range s.LocalDB.B_Tmap{
		s.LocalDB.TagsRefOnBannersMap[v.Tag_id] = append(s.LocalDB.TagsRefOnBannersMap[v.Tag_id], v.Banner_id)
	}
}

func (s *Store) setAllBanners(){
	rows, err := s.DB.Query("select * from Banners")
    if err != nil {
        panic("Banners error:  -=-=-=-="+err.Error())
    }
	defer rows.Close()
	for rows.Next(){
        b := Banners{}
        err := rows.Scan(&b.Id, &b.Title,&b.Text,&b.Url,&b.Feature_id,&b.Visible,&b.CreateTime,&b.UpdateTime)
        if err != nil{
            panic(err.Error())
        }
		s.Mutex.Lock()
        s.LocalDB.BannersMap[b.Id] = &b
		s.Mutex.Unlock()
    }
}

func (s *Store) setAllTags(){
	rows, err := s.DB.Query("select * from Tags")
    if err != nil {
        panic("Tags error:  -=-=-=-="+err.Error())
    }
	defer rows.Close()
	for rows.Next(){
        t := Tags{}
        err := rows.Scan(&t.Id, &t.Name,&t.Description)
        if err != nil{
            panic(err.Error())
        }
		s.Mutex.Lock()
        s.LocalDB.TagsMap[t.Id] = &t
		s.Mutex.Unlock()
    }
}

func (s *Store) setAllFeatures(){
	rows, err := s.DB.Query("select * from Features")
    if err != nil {
        panic("Features error:  -=-=-=-="+err.Error())
    }
	defer rows.Close()
	for rows.Next(){
        f := Features{}
        err := rows.Scan(&f.Id, &f.Name,&f.Description)
        if err != nil{
            panic(err.Error())
        }
		s.Mutex.Lock()
        s.LocalDB.FeaturesMap[f.Id] = &f
		s.Mutex.Unlock()
    }
}

func (s *Store) setAllB_T(){
	rows, err := s.DB.Query("select * from B_T")
    if err != nil {
        panic("B_t error  "+err.Error())
    }
	defer rows.Close()
	for rows.Next(){
        bt := B_T{}
        err := rows.Scan(&bt.Id, &bt.Banner_id,&bt.Tag_id)
        if err != nil{
            panic(err.Error())
        }
		s.Mutex.Lock()
        s.LocalDB.B_Tmap[bt.Id] = &bt
		s.Mutex.Unlock()
    }
	
}