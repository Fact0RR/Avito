package internal_test

import (
	e "github.com/Fact0RR/AVITO/API/entity"
	"testing"
)

func TestQuickErrToJson(t *testing.T) {
	dataError := "dataError"
	exp_json := "{\"error\":\"" + dataError + "\"}"

	res := e.QuickErrToJson(dataError)

	if string(res) != exp_json {
		t.Errorf("Result: %s, expected: %s", string(res), exp_json)
	}
}
