package entity

type UserBanner struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}

type AdminBanner struct{
	Banner_id int `json:"banner_id"`
	Tag_ids []int `json:"tag_ids"`
	Feature_id int `json:"feature_id"`
	UserBanner UserBanner `json:"content"`
	Is_active bool `json:"is_active"`
	Сreated_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`

}

type PostBanner struct{
	Tag_ids *[]int `json:"tag_ids"`
	Feature_id *int `json:"feature_id"`
	UserBanner *UserBanner `json:"content"`
	Is_active *bool `json:"is_active"`
} 

func (pb *PostBanner) IsWithEmptyField()(string,bool){
	switch {
	case pb.Tag_ids == nil:
		return "tag_idstag_ids",true
	case pb.Feature_id == nil:
		return "feature_id",true
	case pb.UserBanner == nil:
		return "content",true
	case pb.Is_active == nil:
		return "is_active",true
	}
	return "",false
}