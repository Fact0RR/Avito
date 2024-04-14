package internal_test

import (
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/Fact0RR/AVITO/config"
	e "github.com/Fact0RR/AVITO/entity"
	"github.com/Fact0RR/AVITO/internal"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationPosthBannerHandle(t *testing.T) {
	a := 2
	b := true
	pb:=e.PostBanner{Tag_ids: &[]int{2}, Feature_id: &a,Is_active: &b, UserBanner: &e.UserBanner{Title: "path",Text: "patchtext",Url: "urlPatch"}}
	b1, _ := json.Marshal(&e.UserBanner{Title: "path",Text: "patchtext",Url: "urlPatch"})
	//expected2 := e.UserBanner{}
	//expected3 := e.UserBanner{}

	testCases := []struct{
		banner_id int
		tag_id int
		feature_id int
		use_last_revision bool
		postBanner *e.PostBanner
		want[]byte
	}{
		
		{//тестируем ответ из Postgres
			banner_id: 4,
			tag_id: 2,
			feature_id: 2,
			use_last_revision: true,
			postBanner: &pb,
			want: b1,
		},
		
	}

	conf := config.GetConfig("../config/config_test.json")
	s:=internal.New(conf)

	go func(){
		if err:=s.Start();err!=nil{
			t.Error(err.Error())
			return
		}
	}()
	err:=s.TryConnectToServer(10,"http://localhost"+conf.Port+"/")
	if err != nil{
		t.Error(err.Error())
		return
	}


	handler := http.HandlerFunc(s.UserBannerHandler)
	for _,tc := range testCases{

		if _,err = s.Store.PostBannerToDB(tc.postBanner);err != nil{
			t.Error(err.Error())
			return
		}

		rec := httptest.NewRecorder()

		req,err := http.NewRequest("GET",fmt.Sprintf("/user_banner?use_last_revision=%t&tag_id=%d&feature_id=%d",tc.use_last_revision,tc.tag_id,tc.feature_id),nil)
		if err != nil{
			t.Error(err.Error())
			return
		}
		req.Header.Set("token", conf.TokenUser)
		handler.ServeHTTP(rec,req)

		assert.Equal(t,tc.want,rec.Body.Bytes())
		
	}

}