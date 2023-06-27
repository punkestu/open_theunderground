package db

import (
	"database/sql"
	"github.com/punkestu/open_theunderground/cerror/invalid"
	"github.com/punkestu/open_theunderground/domain"
	"github.com/savsgio/gotils/uuid"
	"golang.org/x/crypto/bcrypt"
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
			return nil, invalid.New("userId", "Id user is not found")
		}
		return nil, err
	}
	return &mUser, nil
}

func (u UserDB) GetByUsername(username string) (*domain.User, error) {
	var mUser domain.User
	err := u.conn.QueryRow("SELECT * FROM users WHERE username=?", username).Scan(&mUser.ID, &mUser.Fullname, &mUser.Username, &mUser.Password, &mUser.Email)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, invalid.New("username", "username is not found")
		}
		return nil, err
	}
	return &mUser, nil
}

func (u UserDB) Create(fullname, username, password, email string) (*domain.User, error) {
	if mUser, err := u.GetByUsername(username); mUser != nil {
		return nil, invalid.New("username", "username is used")
	} else if err != nil {
		if iErr := invalid.Parse(err); iErr == nil {
			return nil, err
		}
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}
	id := uuid.V4()
	_, err = u.conn.Exec("INSERT INTO users(id, fullname, username, password, email) VALUE (?, ?, ?, ?, ?)", id, fullname, username, string(bytes), email)
	if err != nil {
		return nil, err
	}

	return u.GetByUsername(username)
}

func (u UserDB) Update(userId, fullname, username, email string) (*domain.User, error) {
	if mUser, err := u.GetByUsername(username); mUser != nil {
		if mUser.ID != userId {
			return nil, invalid.New("username", "username is used")
		}
	} else if err != nil {
		if iErr := invalid.Parse(err); iErr == nil {
			return nil, err
		}
	}

	if _, err := u.conn.Exec("UPDATE users SET fullname = ?, username = ?, email = ? WHERE id = ?", fullname, username, email, userId); err != nil {
		return nil, err
	}

	return u.GetByUsername(username)
}
