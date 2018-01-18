package model

import (
	"database/sql"
	"fmt"
	"crypto/md5"
	"meshalka/database"
	"encoding/hex"
)

type User struct {
	UserId     uint64
	Login      string
	Password   string
}

func encodePassword(rawPass string) string {
	sum := md5.Sum([]byte(rawPass))
	return hex.EncodeToString(sum[:])
}

const maxLoginLength = 255

type UserRepository interface {
	SelectUserByLoginInfo(login string, rawPass string) (*User, error)
	SelectUserById(userId uint64) (*User, error)
	AddUser(login string, rawPass string) (bool, error)
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepository{db}
}

type userRepository struct {
	db database.Database
}

func (ur *userRepository) AddUser(login string, rawPass string) (bool, error) {
	if len(login) >= maxLoginLength {
		return false, fmt.Errorf("login too long")
	}

	return getBoolResult(getIntFunctionResult(ur.db, func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT add_user(?, ?)`, login, encodePassword(rawPass))
	}))
}

func (ur *userRepository) SelectUserByLoginInfo(login string, rawPass string) (*User, error) {
	return ur.selectUser(func(con *sql.DB) (*sql.Rows, error) {
		q := `SELECT user_id, login, password AS create_date FROM user WHERE login=? AND password=? LIMIT 1`
		return con.Query(q, login, encodePassword(rawPass))
	})
}

func (ur *userRepository) SelectUserById(userId uint64) (*User, error) {
	if userId == 0 {
		return nil, fmt.Errorf(`user with id 0 not exists`)
	}

	return ur.selectUser(func(con *sql.DB) (*sql.Rows, error) {
		q := `SELECT user_id, login, password AS create_date FROM user WHERE user_id=? LIMIT 1`
		return con.Query(q, userId)
	})
}

func (ur *userRepository) selectUser(q querier) (*User, error) {
	con, err := ur.db.Connection()
	if err != nil {
		return nil, err
	}

	rows, err := q(con)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if user, ok := fetchUser(rows); ok {
		return user, nil
	}

	return nil, fmt.Errorf("user not exists")
}

func fetchUser(rows *sql.Rows) (*User, bool) {
	var id uint64
	var name string
	var pass string

	if rows.Next() {
		if rows.Scan(&id, &name, &pass) == nil {
			return &User{
				UserId:     id,
				Login:      name,
				Password:   pass,
			}, true
		}
	}

	return nil, false
}
