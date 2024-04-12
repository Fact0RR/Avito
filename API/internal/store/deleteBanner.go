package store


func (s *Store) DeleteBannerToDB(id int) error {
	_,err := s.DB.Exec("delete from banners where id = $1",id)
	if err != nil {
		return err
	}
	
	return nil
}