package db

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
)

func NewDBConnection(url string) *pg.DB {
	opt, err := pg.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	db := pg.Connect(opt)

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	return db
}
