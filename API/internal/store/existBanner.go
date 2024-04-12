package store

func (s *Store) IsExistBanner(banner_id int) bool {
	var count int
	s.DB.QueryRow("select count(*) from banners where id = $1",banner_id).Scan(&count)
	return count>0
}