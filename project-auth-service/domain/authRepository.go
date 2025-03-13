package domain

import (
	"database/sql"
	appexception "project-common/exception"
	"project-common/logger"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindBy(username string, password string) (*Login, *appexception.AppError)
	GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *appexception.AppError)
	RefreshTokenExists(refreshToken string) *appexception.AppError
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func (d AuthRepositoryDb) RefreshTokenExists(refreshToken string) *appexception.AppError {
	sqlSelect := "select refresh_token from refresh_token_store where refresh_token = $1"
	var token string
	err := d.client.Get(&token, sqlSelect, refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return appexception.AuthenticationError("refresh token not registered in the store")
		} else {
			logger.Error("Unexpected database error: " + err.Error())
			return appexception.UnexpectedError("unexpected database error")
		}
	}
	return nil
}

func (d AuthRepositoryDb) GenerateAndSaveRefreshTokenToStore(authToken AuthToken) (string, *appexception.AppError) {
	// generate the refresh token
	var appErr *appexception.AppError
	var refreshToken string
	if refreshToken, appErr = authToken.newRefreshToken(); appErr != nil {
		return "", appErr
	}

	// store it in the store
	sqlInsert := "insert into refresh_token_store (refresh_token) values ($1)"
	_, err := d.client.Exec(sqlInsert, refreshToken)
	if err != nil {
		logger.Error("unexpected database error: " + err.Error())
		return "", appexception.UnexpectedError("unexpected database error")
	}
	return refreshToken, nil
}

func (d AuthRepositoryDb) FindBy(username, password string) (*Login, *appexception.AppError) {
	var login Login
	sqlVerify := `SELECT username, u.customer_id, role, string_agg(a.account_id::text, ',') as account_numbers FROM users u
                  LEFT JOIN accounts a ON a.customer_id = u.customer_id
                  WHERE username = $1 and password = $2
                  GROUP BY u.username, u.customer_id, u.role`
	err := d.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appexception.AuthenticationError("invalid credentials")
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, appexception.UnexpectedError("Unexpected database error")
		}
	}
	return &login, nil
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client}
}
