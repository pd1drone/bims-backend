package database

import "github.com/jmoiron/sqlx"

type Positions struct {
	ID          int64  `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func ReadPositions(db sqlx.Ext) ([]*Positions, error) {

	positionArray := make([]*Positions, 0)
	var ID int64
	var Name string
	var Description string

	rows, err := db.Queryx(`SELECT ID,
				Name,
				Description FROM Positions`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &Name, &Description)
		if err != nil {
			return nil, err
		}
		positionArray = append(positionArray, &Positions{
			ID:          ID,
			Name:        Name,
			Description: Description,
		})
	}
	return positionArray, nil
}
