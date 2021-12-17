package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
	"time"
)

type PhoneAuthRequest struct {
	tableName struct{} `pg:"phone_auth_request,alias:par"`
	ID        int
	Phone     string
	Code      string
	ExpiredAt sql.NullTime `pg:"expired_at"`
}

type PhoneAuthRequestStat struct {
	tableName struct{}      `pg:"phone_auth_request,alias:par"`
	RequestID int           `pg:"request_id"`
	Attempts  int           `pg:"attempts"`
	Code      string        `pg:"code"`
	UserId    sql.NullInt64 `pg:"user_id"`
}

//CheckPhoneAuthRequest check request that have not expired
func CheckPhoneAuthRequest(phone string, db *pg.DB) (*PhoneAuthRequestStat, error) {
	var result PhoneAuthRequestStat
	err := db.Model(&result).
		ColumnExpr("par.code").
		ColumnExpr("par.id as request_id").
		ColumnExpr("u.id as user_id").
		ColumnExpr("count(para.id) as attempts").
		Where("par.phone = ?", phone).
		Where("expired_at > now()").
		Where("par.success = false").
		Join("LEFT JOIN phone_auth_request_attempts para on par.id = para.phone_auth_requests_id").
		Join("LEFT JOIN users u on par.phone = u.phone").
		GroupExpr("par.code, par.id, u.id").
		Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func SetSuccessPhoneAuthRequest(id int, tx *pg.Tx) (bool, error) {
	_, err := tx.Model(new(PhoneAuthRequest)).Set("success = true").Where("id = ?", id).Update()
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreatePhoneAuthRequest(phone string, code string, codeTTL int, tx *pg.Tx) error {
	req := PhoneAuthRequest{
		Phone: phone,
		Code:  code,
		ExpiredAt: sql.NullTime{
			Time:  time.Now().Add(time.Second * time.Duration(codeTTL)),
			Valid: true,
		},
	}
	_, err := tx.Model(&req).Insert()
	if err != nil {
		return err
	}
	return nil
}

func CreateRequestSignInByCode(phone string, code string, codeTTL int, db *pg.DB) error {
	currentAuthRequest, err := CheckPhoneAuthRequest(phone, db)
	if err != nil {
		return err
	}
	// already have not expired code
	if currentAuthRequest != nil {
		return errors.New("user must wait new code or enter correct code")
	} else {
		if err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
			exist, err := IsUserExistByCriteria(UserCriteria{Phone: phone}, db)
			if err != nil {
				return err
			}
			if !exist {
				err = InsertUser(phone, tx)
				if err != nil {
					return err
				}
			}
			err = CreatePhoneAuthRequest(phone, code, codeTTL, tx)
			if err != nil {
				return err
			}
			log.Infof("code for phone: %s is %s", phone, code)
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}
