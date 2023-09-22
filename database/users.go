package database

import (
	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID           int64  `json:"ID"`
	FullName     string `json:"FullName"`
	FirstName    string `json:"FirstName"`
	MiddleName   string `json:"MiddleName"`
	LastName     string `json:"LastName"`
	PositionID   int64  `json:"PositionID"`
	PositionName string `json:"PositionName"`
	Email        string `json:"Email"`
	Username     string `json:"Username"`
	IsAdmin      bool   `json:"IsAdmin"`
	ProfileLink  string `json:"ProfileLink"`
}

func CreateUser(db sqlx.Ext, FullName string, FirstName string, MiddleName string, LastName string, PositionID int64, Email string,
	Username string, Password string, ProfileLink string) error {

	_, err := db.Exec(`INSERT INTO Users (
		FullName,
		FirstName,
		MiddleName,
		LastName,
		PositionID,
		Email,
		Username,
		Password,
		IsAdmin,
		ProfileLink
	)
	Values(?,?,?,?,?,?,?,?,?,?)`,
		FullName,
		FirstName,
		MiddleName,
		LastName,
		PositionID,
		Email,
		Username,
		Password,
		false,
		ProfileLink,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Users WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(db sqlx.Ext, ID int64, FullName string, FirstName string, MiddleName string, LastName string, PositionID int64, Email string,
	Username string, ProfileLink string) error {

	_, err := db.Exec(`UPDATE Residents SET 
		FullName = ?,
		FirstName = ?,
		MiddleName = ?,
		LastName = ?,
		PositionID = ?,
		Email = ?,
		Username,
		ProfileLink = ? WHERE ID= ?`,
		FullName,
		FirstName,
		MiddleName,
		LastName,
		PositionID,
		Email,
		Username,
		ProfileLink,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadUsers(db sqlx.Ext) ([]*Users, error) {

	usersArray := make([]*Users, 0)
	var ID int64
	var FullName string
	var FirstName string
	var MiddleName string
	var LastName string
	var PositionID int64
	var Email string
	var Username string
	var IsAdmin bool
	var ProfileLink string

	rows, err := db.Queryx(`SELECT ID,
				FullName,
				FirstName,
				MiddleName,
				LastName,
				PositionID,
				Email,
				Username,
				IsAdmin,
				ProfileLink FROM Users`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &FullName, &FirstName, &MiddleName, &LastName, &PositionID, &Email, &Username, &IsAdmin, &ProfileLink)
		if err != nil {
			return nil, err
		}

		posName, err := GetPositionName(db, PositionID)
		if err != nil {
			return nil, err
		}

		usersArray = append(usersArray, &Users{
			ID:           ID,
			FullName:     FullName,
			FirstName:    FirstName,
			MiddleName:   MiddleName,
			LastName:     LastName,
			PositionID:   PositionID,
			Email:        Email,
			Username:     Username,
			IsAdmin:      IsAdmin,
			ProfileLink:  ProfileLink,
			PositionName: posName,
		})

	}
	return usersArray, nil
}

func GetPositionName(db sqlx.Ext, PositionID int64) (string, error) {
	var PositionName string
	rows, err := db.Queryx(`SELECT Description FROM Positions WHERE ID = ?`, PositionID)

	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&PositionName)
		if err != nil {
			return "", err
		}
	}

	return PositionName, nil
}
