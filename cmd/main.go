package main

import (
	"atlant1da-404/server"
	"log"
)

func main() {

	serv := server.NewServer()
	if err := serv.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
