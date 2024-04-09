package store

type UserBanner struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}

func (s *Store) GetUserBanner(tag_id, feature_id int) *UserBanner {
	ub := UserBanner{}
	for _, v := range s.LocalDB.TagsRefOnBannersMap[tag_id] {
		if s.LocalDB.BannersMap[v].Feature_id == feature_id && s.LocalDB.BannersMap[v].Visible {
			b := s.LocalDB.BannersMap[v]
			ub.Text = b.Text
			ub.Title = b.Title
			ub.Url = b.Url
		}
	}
	return &ub
}