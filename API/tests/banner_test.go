package internal_test

import (
	"testing"
	e "github.com/Fact0RR/AVITO/API/entity"
)

func TestIsWithEmptyField(t *testing.T){
	feature_id := 1
	tag_ids := &[]int{1,2,3}
	userBanner := &e.UserBanner{Text: "text",Title: "title", Url: "url"}
	is_active := true

	pbs := []e.PostBanner{}
	
	pbs = append(pbs, e.PostBanner{Tag_ids: tag_ids})
	pbs = append(pbs, e.PostBanner{Tag_ids: tag_ids, Feature_id:  &feature_id})
	pbs = append(pbs, e.PostBanner{Tag_ids: tag_ids, Feature_id:  &feature_id, UserBanner: userBanner})
	pbs = append(pbs, e.PostBanner{Tag_ids: tag_ids, Feature_id:  &feature_id, UserBanner: userBanner, Is_active: &is_active})
	pbs = append(pbs, e.PostBanner{Tag_ids: tag_ids, Feature_id:  &feature_id, UserBanner: userBanner, Is_active: &is_active})

	expected:= []bool{false,false,false,true,true}


	for i,pb :=range pbs{
		s,res := pb.IsWithEmptyField()
		if res == expected[i]{
			t.Errorf("Result: %t, expected: %t, parametr %s",res,expected[i], s)
		}
	}
}