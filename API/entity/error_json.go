package entity

import "encoding/json"

type JsonError struct {
	Error string `json:"error"`
}

func QuickErrToJson(err string)[]byte{
	je := JsonError{Error: err}
	return je.toJson()
}


func (je *JsonError) toJson() []byte{
	b, err := json.Marshal(je)
	if err != nil {
		panic(err)
	}
	return b
}