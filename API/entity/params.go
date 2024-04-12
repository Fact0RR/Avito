package entity

import (
	"net/url"
	"strconv"
)

const (
	Pstring = "string"
	Pint    = "int"
	Pbool   = "bool"
)

type PValidate struct {
	Name     string
	PType    string
	Content  string
	Required bool
	Exist    bool
}

type UserParams struct {
	Tag_id          int
	Feature_id      int
	UseLastRevision bool
}

type ExistInt struct{
	Exist bool
	Int int
}

type AdminParams struct {
	Tag_id     ExistInt
	Feature_id ExistInt
	Offset     ExistInt
	Limit      ExistInt
}

func (pv *PValidate) IsNoExistRequiredParametr() bool {
	if pv.Required {
		return pv.Exist
	} else {
		return true
	}
}

func (pv *PValidate) CheckType() bool {
	switch {
	case pv.PType == Pstring && pv.Exist:
		return true
	case pv.PType == Pint && pv.Exist:
		return isStingConvertibleToInt(pv.Content)
	case pv.PType == Pbool && pv.Exist:
		return isStringConvertibleToBool(pv.Content)
	case !pv.Required && !pv.Exist:
		return true
	}
	return false
}

func isStingConvertibleToInt(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func CopySlice(oldSlice []PValidate) []PValidate {
	var newSlice []PValidate
	for _, v := range oldSlice {
		newSlice = append(newSlice, v)
	}
	return newSlice
}

func isStringConvertibleToBool(s string) bool {
	return (s == "true" || s == "false")
}

func convertStringToBool(s string) bool {
	return s == "true"
}

func Fill(qp url.Values, pvs []PValidate) {
	for i := 0; i < len(pvs); i++ {
		if qp.Has(pvs[i].Name) {
			pvs[i].Exist = true
			pvs[i].Content = qp.Get(pvs[i].Name)
		}
	}
}
