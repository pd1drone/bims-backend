package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type LoginResponse struct {
	ID          int64  `json:"ID"`
	UserName    string `json:"Username"`
	FullName    string `json:"FullName"`
	ProfileLink string `json:"ProfileLink"`
	IsAdmin     bool   `json:"IsAdmin"`
}

func Login(db sqlx.Ext, username string, password string) (*LoginResponse, error) {

	counter := 0
	var id int64
	var user string
	var pass string
	var fullname string
	var profilelink string
	var isadmin bool

	rows, err := db.Queryx(`SELECT u.ID, u.Username, u.Password, u.FullName, u.ProfileLink, u.IsAdmin FROM Users as u
	WHERE BINARY u.Username=? AND BINARY u.Password=?`,
		username, password)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &user, &pass, &fullname, &profilelink, &isadmin)
		if err != nil {
			return nil, err
		}
		counter++
	}

	if counter == 0 {
		return nil, fmt.Errorf("User does not exists")
	}

	return &LoginResponse{
		ID:          id,
		UserName:    user,
		FullName:    fullname,
		ProfileLink: profilelink,
		IsAdmin:     isadmin,
	}, nil
}
