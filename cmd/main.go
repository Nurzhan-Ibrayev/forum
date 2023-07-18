package main

import (
	"fmt"
	"log"
	"net/http"

	"01.alem.school/git/btussupb/forum/internal/config"
	"01.alem.school/git/btussupb/forum/internal/handler"
	"01.alem.school/git/btussupb/forum/internal/repository"
)

func main() {
	fmt.Println("Server is listening in port localhost:1999")

	cfg, err := config.InitConfig("./config/config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := repository.OpenDb(cfg.Database)
	if err != nil {
		fmt.Println(err, "with main, opendb")
		return
	}

	defer db.Close()

	storage := repository.NewStorage(db)
	router := handler.NewHandler(storage)

	if err := http.ListenAndServe(":1999", router.Router()); err != nil {
		log.Fatal(err)
	}
}
