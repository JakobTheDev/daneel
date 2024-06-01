package models

import (
	"github.com/JakobTheDev/daneel/internal/database"
)

type Platform struct {
	ID          int64
	DisplayName string
}

func ListPlatforms() ([]Platform, error) {
	rows, err := database.DB.Query("SELECT * FROM Platform ORDER BY DisplayName ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var platforms []Platform

	for rows.Next() {
		var platform Platform

		err := rows.Scan(&platform.ID, &platform.DisplayName)
		if err != nil {
			return nil, err
		}

		platforms = append(platforms, platform)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return platforms, nil
}

func AddPlatform(p Platform) error {
	_, err := database.DB.Exec("INSERT INTO Platform (DisplayName) VALUES (?)", p.DisplayName)
	if err != nil {
		return err
	}

	return nil
}
