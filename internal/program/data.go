package program

import (
	"github.com/JakobTheDev/daneel/internal/database"
	"github.com/JakobTheDev/daneel/internal/models"
)

func GetProgramListFromDb() ([]models.Program, error) {
	rows, err := database.DB.Query(`SELECT *,
									(SELECT p.[Name] FROM Platform p WHERE p.[Id] = pr.[PlatformId]) as PlatformName
									FROM [Program] pr
									WHERE pr.[IsActive] = 1
									ORDER BY Name ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programs []models.Program

	for rows.Next() {
		var platform models.Program

		err := rows.Scan(&platform.Id, &platform.PlatformId, &platform.Name, &platform.IsPrivate, &platform.IsActive, &platform.PlatformName)
		if err != nil {
			return nil, err
		}

		programs = append(programs, platform)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return programs, nil
}

func AddProgramToDb(program models.Program) error {
	var err error

	err = database.DB.QueryRow(`SELECT [Id] 
								FROM [Platform]
								WHERE [Name] = ?`, program.PlatformName).Scan(&program.PlatformId)
	if err != nil {
		return err
	}
	if program.PlatformId == 0 {
		return nil
	}

	_, err = database.DB.Exec(`
		IF NOT EXISTS (
			SELECT 1 
			FROM [Program] 
			WHERE [Name] = ?) 
		INSERT INTO [Program] ([PlatformId], [Name], [IsPrivate]) VALUES (?, ?, ?) 
		ELSE UPDATE [Program] 
			 SET [IsActive] = 1,
			 	 [IsPrivate] = ?
			 WHERE [Name] = ?`, program.Name, program.PlatformId, program.Name, program.IsPrivate, program.IsPrivate, program.Name)
	if err != nil {
		return err
	}

	return nil
}

func RemoveProgramFromDb(p models.Program) error {
	_, err := database.DB.Exec("UPDATE [Program] SET [IsActive] = 0 WHERE [DisplayName] = ?", p.Name)
	if err != nil {
		return err
	}

	return nil
}
