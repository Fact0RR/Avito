package store

import (
	"github.com/Fact0RR/AVITO/API/entity"
	"github.com/lib/pq"
)

func (s *Store) PostBannerToDB(pb *entity.PostBanner) (int, error) {
	var id int
	err := s.DB.QueryRow("select createBannerWithTags($1,$2,$3,$4,$5,$6)", pq.Array(*pb.Tag_ids), pb.Feature_id, pb.UserBanner.Title, pb.UserBanner.Text, pb.UserBanner.Url, pb.Is_active).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}