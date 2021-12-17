package models

import (
	"context"
	"errors"
	"github.com/go-pg/pg/v10"
)

type PhoneAuthRequestAttempt struct {
	tableName           struct{} `pg:"phone_auth_request_attempts,alias:para"`
	ID                  int
	PhoneAuthRequestsID int `pg:"phone_auth_requests_id"`
	InputCode           string
	Success             bool
}

func CreatePhoneAuthRequestAttempt(reqID int, inputCode string, success bool, tx *pg.Tx) error {
	req := PhoneAuthRequestAttempt{
		PhoneAuthRequestsID: reqID,
		InputCode:           inputCode,
		Success:             success,
	}
	_, err := tx.Model(&req).Insert()
	if err != nil {
		return err
	}
	return nil
}

func SignInByPhoneAuthCode(phone string, inputCode string, maxAttemptsCount int, db *pg.DB) (int, error) {
	currentAuthRequest, err := CheckPhoneAuthRequest(phone, db)
	if err != nil {
		return 0, err
	}
	if currentAuthRequest == nil {
		return 0, errors.New("code was expired or not requested")
	}
	if currentAuthRequest.Attempts+1 > maxAttemptsCount {
		return 0, errors.New("too many tries enter code")
	}
	success := currentAuthRequest.Code == inputCode
	if err = db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		err = CreatePhoneAuthRequestAttempt(currentAuthRequest.RequestID, inputCode, success, tx)
		if err != nil {
			return err
		}
		if success {
			_, err = SetSuccessPhoneAuthRequest(currentAuthRequest.RequestID, tx)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return 0, err
	}
	if !success {
		return 0, errors.New("wrong code")
	}

	return int(currentAuthRequest.UserId.Int64), nil
}
