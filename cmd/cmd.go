package cmd

import (
	"log"
	"mangosteen/internal/router"
)

func RunServer() {
	r := router.New()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("r.Run 的下一行")
}
