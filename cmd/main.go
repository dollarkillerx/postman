package main

import (
	"log"

	"github.com/dollarkillerx/postman/internal/conf"
	"github.com/dollarkillerx/postman/internal/server"
)

func main() {
	log.Println("Postman Init Config: ", conf.Conf)

	ser := server.NewServer()
	err := ser.Run(conf.Conf.PostmanAddr)
	if err != nil {
		log.Fatalln(err)
	}
}
