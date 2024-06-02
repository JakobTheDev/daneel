package models

import (
	"github.com/JakobTheDev/daneel/internal/database"
)

type Platform struct {
	ID          int64
	DisplayName string
	IsActive    bool
}

func ListPlatforms(showInactive bool) ([]Platform, error) {
	rows, err := database.DB.Query(`SELECT * 
									FROM [Platform] 
									WHERE ([IsActive] = 1 OR ? = 1)
									ORDER BY [DisplayName] ASC`, showInactive)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var platforms []Platform

	for rows.Next() {
		var platform Platform

		err := rows.Scan(&platform.ID, &platform.DisplayName, &platform.IsActive)
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
	// Insert platform if not exists, else set to active
	_, err := database.DB.Exec(`
		IF NOT EXISTS (
			SELECT 1 
			FROM [Platform] 
			WHERE [DisplayName] = ?) 
		INSERT INTO [Platform] ([DisplayName]) VALUES (?) 
		ELSE UPDATE [Platform] 
			 SET [IsActive] = 1 
			 WHERE [DisplayName] = ?`, p.DisplayName, p.DisplayName, p.DisplayName)
	if err != nil {
		return err
	}

	return nil
}

func RemovePlatform(p Platform) error {
	_, err := database.DB.Exec("UPDATE [Platform] SET [IsActive] = 0 WHERE [DisplayName] = ?", p.DisplayName)
	if err != nil {
		return err
	}

	return nil
}
