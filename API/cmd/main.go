package main

import (
	"log"

	"github.com/Fact0RR/AVITO/API/config"
	"github.com/Fact0RR/AVITO/API/internal"
)

func main() {

	conf := config.GetConfig("API/config/config.json")
	s:=internal.New(conf)
	if err:=s.Start();err!=nil{
		log.Fatal(err)
	}
}
