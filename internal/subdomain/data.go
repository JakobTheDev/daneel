package subdomain

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/database"
	"github.com/JakobTheDev/daneel/internal/models"
)

func GetSubdomainListFromDb(programName string, domainName string, showOutOfScope bool) ([]models.Subdomain, error) {
	rows, err := database.DB.Query(`SELECT s.*,
										(SELECT d.[Name]
										FROM Domain d
										WHERE d.[Id] = s.[DomainId]) AS DomainName
									FROM [Subdomain] s
									JOIN [Domain] d on s.[DomainId] = d.[Id]
									JOIN [Program] p on d.[ProgramId] = p.[Id]
									WHERE (s.[IsInScope] = 1 OR ? = 1) AND
										s.[IsActive] = 1 AND
										(p.[Name] = ? OR d.[Name] = ?)
									ORDER BY [Name] ASC`, showOutOfScope, programName, domainName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subdomains []models.Subdomain

	for rows.Next() {
		var subdomain models.Subdomain

		err := rows.Scan(&subdomain.ID, &subdomain.DomainId, &subdomain.Name, &subdomain.IsInScope, &subdomain.IsActive, &subdomain.DomainName)
		if err != nil {
			return nil, err
		}

		subdomains = append(subdomains, subdomain)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subdomains, nil
}

func AddSubdomainToDb(subdomain models.Subdomain) (isInserted bool, err error) {
	err = database.DB.QueryRow(`SELECT [Id] 
								FROM [Domain]
								WHERE [Name] = ?`, subdomain.DomainName).Scan(&subdomain.DomainId)
	if err != nil {
		return false, err
	}
	if subdomain.DomainId == 0 {
		return false, fmt.Errorf("domain not found")
	}

	// The mssql driver doesn't support the OUTPUT clause, so we have to do this in two steps
	// First, check if the subdomain exists, then either insert or update
	// Returns 1 if inserted, 0 if updated
	row := database.DB.QueryRow(`
		DECLARE @IsInserted BIT
		IF NOT EXISTS (
			SELECT 1
			FROM [Subdomain]
			WHERE [Name] = ?)
		BEGIN
			INSERT INTO [Subdomain] ([DomainId], [Name], [IsInScope]) VALUES (?, ?, ?);
			SET @IsInserted = 1;
		END
		ELSE 
		BEGIN
			UPDATE [Subdomain]
				SET [IsInScope] = ?
				WHERE [Name] = ?;
			SET @IsInserted = 0;
		END
		SELECT @IsInserted`, subdomain.Name, subdomain.DomainId, subdomain.Name, subdomain.IsInScope, subdomain.IsInScope, subdomain.Name)

	err = row.Scan(&isInserted)
	if err != nil {
		return false, err
	}

	return isInserted, nil
}

func RemoveSubdomainFromDb(subdomain models.Subdomain) error {
	_, err := database.DB.Exec("UPDATE [Subdomain] SET [IsActive] = 0 WHERE [Name] = ?", subdomain.Name)
	if err != nil {
		return err
	}

	return nil
}
