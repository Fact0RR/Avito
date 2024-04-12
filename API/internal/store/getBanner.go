package store

import "github.com/Fact0RR/AVITO/API/entity"

func (s *Store) GetUserBanner(tag_id, feature_id int) *entity.UserBanner {
	ub := entity.UserBanner{}
	s.Mutex.Lock()
	for _, v := range s.LocalDB.TagsRefOnBannersMap[tag_id] {
		if s.LocalDB.BannersMap[v].Feature_id == feature_id && s.LocalDB.BannersMap[v].Visible {
			b := s.LocalDB.BannersMap[v]
			ub.Text = b.Text
			ub.Title = b.Title
			ub.Url = b.Url
		}
	}
	s.Mutex.Unlock()
	return &ub
}

func (s *Store) GetUserBannerFromDB(tag_id, feature_id int) *entity.UserBanner {
	ub := entity.UserBanner{}
	s.DB.QueryRow("select b.title, b.text, b.url from banners b "+
		"join b_t bt on bt.banner_id = b.id "+
		"join tags t on t.id = bt.tag_id "+
		"where b.visible = true and t.id =$1 and b.feature_id = $2", tag_id, feature_id).Scan(&ub.Title, &ub.Text, &ub.Url)

	return &ub
}

func (s *Store) GetAdminBannerFromDB(query string) ([]entity.AdminBanner, error) {
	abs := []entity.AdminBanner{}
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ab := entity.AdminBanner{}
		err := rows.Scan(&ab.Banner_id, &ab.Feature_id, &ab.UserBanner.Title, &ab.UserBanner.Text, &ab.UserBanner.Url, &ab.Is_active, &ab.Сreated_at, &ab.Updated_at)
		if err != nil {
			return nil, err
		}
		abs = append(abs, ab)
	}

	for i := 0; i < len(abs); i++ {
		rows, err := s.DB.Query("select tag_id from b_t where banner_id = $1", abs[i].Banner_id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var tag_id int
			err := rows.Scan(&tag_id)
			if err != nil {
				return nil, err
			}
			abs[i].Tag_ids = append(abs[i].Tag_ids, tag_id)
		}
	}

	return abs, nil
}