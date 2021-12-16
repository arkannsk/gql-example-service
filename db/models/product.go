package models

import (
	"github.com/go-pg/pg/v10"
)

type Product struct {
	tableName struct{} `pg:"products,alias:pr"`
	ID        int      `pg:"id"`
	Name      string   `pg:"name"`
}

func GetProductByIDs(ids []int, db *pg.DB) ([]*Product, error) {
	var products []*Product
	err := db.Model(&products).
		WhereIn("id in (?)", ids).
		Select()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetRandomProducts(n int, db *pg.DB) ([]*Product, error) {
	var products []*Product
	err := db.Model(&products).
		Limit(n).
		OrderExpr("RANDOM()").
		Select()
	if err != nil {
		return nil, err
	}
	return products, nil
}
