package internal_test

import (
	"testing"
		e "github.com/Fact0RR/AVITO/entity"
)

func TestCheckType(t *testing.T){
	cases:=[]e.PValidate{}
	cases = append(cases, e.PValidate{Required: true,Exist: false})
	cases = append(cases, e.PValidate{Required: false,Exist: false})
	cases = append(cases, e.PValidate{Required: false,Exist: true,PType: e.Pint,Content: "123s"})
	cases = append(cases, e.PValidate{Required: false,Exist: true,PType: e.Pint,Content: "123"})
	cases = append(cases, e.PValidate{Required: true,Exist: true,PType: e.Pstring,Content: "123"})
	cases = append(cases, e.PValidate{Required: true,Exist: true,PType: e.Pbool,Content: "false"})
	cases = append(cases, e.PValidate{Required: true,Exist: false,PType: e.Pbool,Content: "true"})
	cases = append(cases, e.PValidate{Required: true,Exist: true,PType: e.Pbool,Content: "true"})

	exp1 := []bool{false,true,false,true,true,true,false,true}
	exp2 := []bool{false,true,true,true,true,true,false,true}



	for i,c := range cases{
		if exp1[i] != c.CheckType(){
			t.Errorf("Result: %t, expected: %t",c.CheckType(),exp1[i])
		}
	}

	for i,c := range cases{
		if exp2[i] != c.IsNoExistRequiredParametr(){
			t.Errorf("Result: %t, expected: %t",c.CheckType(),exp2[i])
		}
	}
}