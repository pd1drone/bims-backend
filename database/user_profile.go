package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func ChangePassword(db sqlx.Ext, ID int64, user string, pass string, hashedNewPassword string) error {

	var username string
	var password string

	rows, err := db.Queryx(`SELECT Username, Password FROM Users
	WHEREUsername=? AND Password=?`,
		user, pass)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&username, &password)
		if err != nil {
			return err
		}
	}

	if pass != password {
		return fmt.Errorf("Wrong password!")
	}

	_, err = db.Exec(`UPDATE Users SET Password= ? WHERE ID= ?`,
		hashedNewPassword,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func UploadUserPicture(db sqlx.Ext, ID int64, ProfileLink string) error {

	_, err := db.Exec(`UPDATE Users SET ProfileLink= ? WHERE ID= ?`,
		ProfileLink,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}
