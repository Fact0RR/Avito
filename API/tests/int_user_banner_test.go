package internal_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fact0RR/AVITO/API/config"
	e "github.com/Fact0RR/AVITO/API/entity"
	"github.com/Fact0RR/AVITO/API/internal"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUserBannerHandle(t *testing.T) {

	b1, _ := json.Marshal(e.UserBanner{Title: "Banner 1",Url: "http://url_for_banner1",Text: "some text for banner 1 etc.."})
	//expected2 := e.UserBanner{}
	//expected3 := e.UserBanner{}
	b4, _ := json.Marshal(e.UserBanner{Title: "Banner 3",Url: "http://url_for_banner3",Text: "some text for banner 3 "})

	testCases := []struct{
		tag_id int
		feature_id int
		use_last_revision bool
		want[]byte
	}{
		{//тестируем ответ из ОЗУ
			tag_id: 1,
			feature_id: 1,
			use_last_revision: false,
			want: b1,
		},
		{
			tag_id: 1,
			feature_id: 2,
			use_last_revision: false,
			want: nil,//видимость этого банера отключена
		},
		{
			tag_id: 3,
			feature_id: 2,
			use_last_revision: false,
			want: nil,//этот банер отсутствует
		},
		{
			tag_id: 1,
			feature_id: 3,
			use_last_revision: false,
			want: b4,
		},
		{//тестируем ответ из Postgres
			tag_id: 1,
			feature_id: 1,
			use_last_revision: true,
			want: b1,
		},
		{
			tag_id: 1,
			feature_id: 2,
			use_last_revision: true,
			want: nil,//видимость этого банера отключена
		},
		{
			tag_id: 3,
			feature_id: 2,
			use_last_revision: true,
			want: nil,//этот банер отсутствует
		},
		{
			tag_id: 1,
			feature_id: 3,
			use_last_revision: true,
			want: b4,
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
	for i,tc := range testCases{
		rec := httptest.NewRecorder()

		req,err := http.NewRequest("GET",fmt.Sprintf("/user_banner?use_last_revision=%t&tag_id=%d&feature_id=%d",tc.use_last_revision,tc.tag_id,tc.feature_id),nil)
		if err != nil{
			t.Error(err.Error())
			return
		}
		req.Header.Set("token", conf.TokenUser)
		handler.ServeHTTP(rec,req)

		
		fmt.Println("----------")
		fmt.Println(string(tc.want))
		fmt.Println("+_+_+-+_+_+_+_")
		fmt.Println(rec.Body.String())
		fmt.Println("Итерация: " , i)
		fmt.Println("----------")
		fmt.Println(fmt.Sprintf("/user_banner?use_last_revision=%t&tag_id=%d&feature_id=%d",tc.use_last_revision,tc.tag_id,tc.feature_id))

		assert.Equal(t,tc.want,rec.Body.Bytes())
		
	}

}
