package internal_test

import (
	"testing"

	e "github.com/Fact0RR/AVITO/API/entity"
	"github.com/Fact0RR/AVITO/API/internal"
)

type limoff struct{
	Off e.ExistInt
	Lim e.ExistInt
}
type expected struct{
	p1 string
	p2 string
}

func TestGetQueryLimitAndOffset(t *testing.T){
	loa := []limoff{}
	loa = append(loa, limoff{Off:e.ExistInt{true,1},Lim: e.ExistInt{true,1}})
	loa = append(loa, limoff{Off:e.ExistInt{false,1},Lim: e.ExistInt{true,1}})
	loa = append(loa, limoff{Off:e.ExistInt{true,1},Lim: e.ExistInt{false,1}})
	loa = append(loa, limoff{Off:e.ExistInt{false,1},Lim: e.ExistInt{false,1}})

	exps := []expected{}
	exps = append(exps, expected{" limit 1"," offset 1"})
	exps = append(exps, expected{" limit 1",""})
	exps = append(exps, expected{"","offset 1"})
	exps = append(exps, expected{"",""})

	for i,lo := range loa{
		p1,p2 := internal.GetQueryLimitAndOffset(lo.Lim,lo.Off)
		if !(exps[i].p1 == p1 && exps[i].p2 == p2){
			t.Errorf("Result: %s,%s, expected: %s,%s",exps[i].p1,exps[i].p2,p1,p2)
		}
	}
}