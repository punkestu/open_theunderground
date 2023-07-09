package db

import (
	"database/sql"
	"github.com/punkestu/open_theunderground/shared/domain"
	"github.com/punkestu/open_theunderground/shared/exception"
)

type UserDB struct {
	conn *sql.DB
}

func NewUserDB(conn *sql.DB) *UserDB {
	return &UserDB{conn: conn}
}

func (u UserDB) GetByID(userId string) (*domain.User, error) {
	var mUser domain.User
	err := u.conn.QueryRow("SELECT * FROM users WHERE id=?", userId).Scan(&mUser.ID, &mUser.Fullname, &mUser.Username, &mUser.Password, &mUser.Email)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, exception.New("userId", "Id user is not found")
		}
		return nil, err
	}
	return &mUser, nil
}
