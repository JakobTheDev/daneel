package platform

import (
	"github.com/JakobTheDev/daneel/internal/database"
	"github.com/JakobTheDev/daneel/internal/models"
)

func GetPlatformListFromDb(showInactive bool) ([]models.Platform, error) {
	rows, err := database.DB.Query(`SELECT * 
									FROM [Platform] 
									WHERE ([IsActive] = 1 OR ? = 1)
									ORDER BY [Name] ASC`, showInactive)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var platforms []models.Platform

	for rows.Next() {
		var platform models.Platform

		err := rows.Scan(&platform.ID, &platform.Name, &platform.IsActive)
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

func AddPlatformToDb(platform models.Platform) error {
	// Insert platform if not exists, else set to active
	_, err := database.DB.Exec(`
		IF NOT EXISTS (
			SELECT 1 
			FROM [Platform] 
			WHERE [Name] = ?) 
		INSERT INTO [Platform] ([Name]) VALUES (?) 
		ELSE UPDATE [Platform] 
			 SET [IsActive] = 1 
			 WHERE [Name] = ?`, platform.Name, platform.Name, platform.Name)
	if err != nil {
		return err
	}

	return nil
}

func RemovePlatformFromDb(platform models.Platform) error {
	_, err := database.DB.Exec("UPDATE [Platform] SET [IsActive] = 0 WHERE [Name] = ?", platform.Name)
	if err != nil {
		return err
	}

	return nil
}
