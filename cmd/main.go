package main

import (
	"log"

	"github.com/kenedyCO/todo"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("error occured while runnung http server: %s", err.Error())
	}
}
