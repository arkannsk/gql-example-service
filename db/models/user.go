package models

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type User struct {
	tableName struct{} `pg:"users,alias:u"`
	ID        int      `pg:"id"`
	Phone     string   `pg:"phone"`
}

type UserCriteria struct {
	ID    int
	Phone string
}

func GetUserByCriteria(c UserCriteria, db *pg.DB) (*User, error) {
	var user User
	err := db.Model(&user).
		WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			if c.Phone != "" {
				q = q.Where("phone = ?", c.Phone)
			}
			if c.ID > 0 {
				q = q.Where("id = ?", c.ID)
			}
			return q, nil
		}).
		Where("deleted_at IS NULL").
		Select()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func IsUserExistByCriteria(c UserCriteria, db *pg.DB) (bool, error) {
	exst, err := db.Model(new(User)).
		WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			if c.Phone != "" {
				q = q.Where("phone = ?", c.Phone)
			}
			if c.ID > 0 {
				q = q.Where("id = ?", c.ID)
			}
			return q, nil
		}).
		Where("deleted_at IS NULL").
		Exists()
	if err != nil {
		return false, err
	}
	return exst, nil
}

func InsertUser(phone string, tx *pg.Tx) error {
	user := User{
		Phone: phone,
	}
	_, err := tx.Model(&user).OnConflict("DO NOTHING").Insert()
	if err != nil {
		return err
	}
	return nil
}
