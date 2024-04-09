package middlewares

import (
	"net/http"
	"net/url"
	"strconv"
)

type Params struct{
	Tag_id int
	Feature_id int
	UseLastRevision bool
}

func ValidateUserBannerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		req_qp_Arr := []string{"tag_id","feature_id"}
		noReq_QP_Arr := []string{"use_last_revision"}
		if !checkOnExistParams(queryParams,req_qp_Arr){
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Отсутствуют обязательные параметры в url"))
			return
		} 
		if !checkOnTypeParams(queryParams,req_qp_Arr,noReq_QP_Arr){
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Типы параметров не соблюдены"))
			return
		}
		
		p := Params{}
		tag_id, _ := strconv.Atoi(queryParams.Get(req_qp_Arr[0]))
		p.Tag_id = tag_id
		feature_id,_:=strconv.Atoi(queryParams.Get(req_qp_Arr[1]))
		p.Feature_id = feature_id

		if queryParams.Has(noReq_QP_Arr[0]){
			p.UseLastRevision=convertStringToBool(queryParams.Get(noReq_QP_Arr[0]))
		}

		if !checkOnDataParams(&p){
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id должны быть больше 0"))
			return
		}
		next(w,r)
	}
}

func checkOnExistParams(queryParam url.Values, qpArr []string) bool{

	for _,s := range qpArr{
		if !queryParam.Has(s){
			return false
		}
	}
	return true
}

func checkOnDataParams(p *Params) bool{
	return (p.Feature_id > 0 && p.Tag_id > 0)
}

func checkOnTypeParams(queryParam url.Values, req_qp_Arr,noreq_qp_Arr []string) bool{

	for _,s := range req_qp_Arr{
		if !isStingConvertibleToInt(queryParam.Get(s)){
			return false
		}
	}
	for _,s := range noreq_qp_Arr{
		if queryParam.Has(s){
			if !isStringConvertibleToBool(queryParam.Get(s)){
				return false
			}
		}
	}
	return true
}

func isStingConvertibleToInt(s string) bool{
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func isStringConvertibleToBool(s string) bool{
	return (s == "true" || s == "false")
}

func convertStringToBool(s string) bool{
	return s == "true"
}

