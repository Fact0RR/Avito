package internal_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/Fact0RR/AVITO/config"
	e "github.com/Fact0RR/AVITO/entity"
	"github.com/Fact0RR/AVITO/internal"
)

func TestIntegrationGetBannerHandle(t *testing.T) {

	ab := e.AdminBanner{Banner_id: 2,Feature_id: 2,Tag_ids: []int{2,1},UserBanner: e.UserBanner{Title: "Banner 2",Text: "some text for banner2 etc..",Url: "http://url_for_banner2"},Is_active: false}
	abs := []e.AdminBanner{ab}


	testCases := []struct{
		tag_id int
		feature_id int
		want[]e.AdminBanner
	}{
		{//тестируем ответ из ОЗУ
			tag_id: 1,
			feature_id: 2,
			want: abs,
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
	handler := http.HandlerFunc(s.BannerHandlerGet)
	for _,tc := range testCases{
		rec := httptest.NewRecorder()

		req,err := http.NewRequest("GET",fmt.Sprintf("/banner?tag_id=%d&feature_id=%d",tc.tag_id,tc.feature_id),nil)
		if err != nil{
			t.Error(err.Error())
			return
		}
		req.Header.Set("token", conf.TokenAdmin)
		handler.ServeHTTP(rec,req)

		b,err := io.ReadAll(rec.Body)
		if err != nil{
			t.Error(err.Error())
			return
		}
		var resAB []e.AdminBanner
		if err = json.Unmarshal(b,&resAB);err !=nil{
			t.Errorf(err.Error())
			return
		}


		if !isSliceAdminBannerEqual(resAB, tc.want) {
			t.Errorf(err.Error())
			return
		}
	}
}

func isSliceAdminBannerEqual(res,want []e.AdminBanner)bool{
	if len(res)!=len(want){
		return false
	}
	for i,r := range res{
		if want[i].Banner_id!=r.Banner_id{
			return false
		}
		if want[i].Feature_id!=r.Feature_id{
			return false
		}
		if want[i].Is_active!=r.Is_active{
			return false
		}
		if !isTagSlicesEqual(want[i].Tag_ids, r.Tag_ids){
			return false
		}
		if want[i].UserBanner!=r.UserBanner{
			return false
		}
	}
	return true
}

func isTagSlicesEqual(a,b []int) bool{
	if len(a)!=len(b){
		return false
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for i,a_ :=range a{
		if a_ != b[i]{
			return false
		}
	}
	return true
}