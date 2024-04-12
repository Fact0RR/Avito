package store

import (

	"github.com/Fact0RR/AVITO/API/entity"
	"github.com/lib/pq"
)

func (s *Store) PatchBannerToDB(id int,pb *entity.PostBanner) error {
	_,err := s.DB.Exec("select updateBannerWithTags($1,$2,$3,$4,$5,$6,$7)",id, pq.Array(*pb.Tag_ids), pb.Feature_id, pb.UserBanner.Title, pb.UserBanner.Text, pb.UserBanner.Url, pb.Is_active)
	if err != nil {
		return err
	}
	
	return nil
}