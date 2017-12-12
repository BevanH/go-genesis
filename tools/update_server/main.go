package main

import (
	"path/filepath"

	"log"

	"github.com/AplaProject/go-apla/tools/update_server/config"
	"github.com/AplaProject/go-apla/tools/update_server/database"
	"github.com/AplaProject/go-apla/tools/update_server/web"
)

func main() {
	p := config.NewParser(filepath.Join(".", "resources"))
	c, err := p.Do()
	if err != nil {
		log.Fatalf("Config parsing error: %s", err.Error())
	}

	db, err := database.NewDatabase(c.DBPath)
	if err != nil {
		log.Fatalf("Creation database error: %s", err.Error())
	}

	s := web.Server{
		Db:   &db,
		Conf: &c,
	}

	log.Fatalf("Server running error: %s", s.Run().Error())
}
