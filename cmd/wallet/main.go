package main

import (
	"github.com/bahodurnazarov/wallet/internal/app"
	"github.com/bahodurnazarov/wallet/pkg/config"
	"github.com/bahodurnazarov/wallet/pkg/db"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalln("Failed init cfg", err.Error())
	}
	a := app.New(cfg)
	defer db.CloseDB()

	log.Fatalln(a.Run(cfg))
}
