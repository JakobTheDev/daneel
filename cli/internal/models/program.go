package models

import (
	"github.com/JakobTheDev/daneel/internal/database"
)

type Program struct {
	Id           int64
	PlatformId   int64
	PlatformName string
	DisplayName  string
	IsPrivate    bool
	IsActive     bool
}

func ListPrograms() ([]Program, error) {
	rows, err := database.DB.Query(`SELECT *,
									(SELECT p.[DisplayName] FROM Platform p WHERE p.[Id] = pr.[PlatformId]) as PlatformName
									FROM [Program] pr
									WHERE pr.[IsActive] = 1
									ORDER BY DisplayName ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programs []Program

	for rows.Next() {
		var platform Program

		err := rows.Scan(&platform.Id, &platform.PlatformId, &platform.DisplayName, &platform.IsPrivate, &platform.IsActive, &platform.PlatformName)
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

func AddProgram(program Program) error {
	var err error

	err = database.DB.QueryRow(`SELECT [Id] 
								FROM [Platform]
								WHERE [DisplayName] = ?`, program.PlatformName).Scan(&program.PlatformId)
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
			WHERE [DisplayName] = ?) 
		INSERT INTO [Program] ([PlatformId], [DisplayName], [IsPrivate]) VALUES (?, ?, ?) 
		ELSE UPDATE [Program] 
			 SET [IsActive] = 1,
			 	 [IsPrivate] = ?
			 WHERE [DisplayName] = ?`, program.DisplayName, program.PlatformId, program.DisplayName, program.IsPrivate, program.IsPrivate, program.DisplayName)
	if err != nil {
		return err
	}

	return nil
}

func RemoveProgram(p Program) error {
	_, err := database.DB.Exec("UPDATE [Program] SET [IsActive] = 0 WHERE [DisplayName] = ?", p.DisplayName)
	if err != nil {
		return err
	}

	return nil
}
