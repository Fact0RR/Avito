package main

import (
	"log"

	"github.com/Fact0RR/AVITO/config"
	"github.com/Fact0RR/AVITO/internal"
)

func main() {

	conf := config.GetConfig()
	s:=internal.New(conf)
	if err:=s.Start();err!=nil{
		log.Fatal(err)
	}
}
