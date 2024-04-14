package internal_test

import (
	"testing"

	"github.com/Fact0RR/AVITO/API/config"
	"github.com/Fact0RR/AVITO/API/internal/store"
	_ "github.com/lib/pq"
)

func TestPingDB(t *testing.T) {
	conf := config.GetConfig("../config/config_test.json")

	err:=store.TryConnectToDB(15,conf.DataBaseString)
	if err != nil{
		t.Error(err.Error())
		return
	}

	store_test := store.New(conf.DataBaseString)
	if err := store_test.Open();err != nil {
		t.Error("Проблема с инициализацией базы данных: "+err.Error())
	}
}

