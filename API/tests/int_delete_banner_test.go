package internal_test

import (
	"fmt"

	"net/http"
	"net/http/httptest"
	
	"testing"

	"github.com/Fact0RR/AVITO/config"
	"github.com/Fact0RR/AVITO/internal"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationDeleteBannerHandle(t *testing.T) {

	testCases := []struct{
		banner_id int
		tag_id int
		feature_id int
		use_last_revision bool
		want[]byte
	}{
		{
			banner_id: 1,
			tag_id: 1,
			feature_id: 1,
			use_last_revision: true,
			want: nil,
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

	tc := testCases[0]
	s.Store.DeleteBannerToDB(tc.banner_id)

	handler := http.HandlerFunc(s.UserBannerHandler)
	rec := httptest.NewRecorder()
	req,err := http.NewRequest("GET",fmt.Sprintf("/user_banner?use_last_revision=%t&tag_id=%d&feature_id=%d",tc.use_last_revision,tc.tag_id,tc.feature_id),nil)
	if err != nil{
		t.Error(err.Error())
		return
	}
	req.Header.Set("token", conf.TokenAdmin)
	handler.ServeHTTP(rec,req)
	
	assert.Equal(t,tc.want,rec.Body.Bytes())

}