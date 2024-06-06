package models

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/database"
)

type Subdomain struct {
	ID            int64
	DomainId      int64
	DomainName    string
	SubdomainName string
	IsInScope     bool
	IsActive      bool
}

func ListSubdomains(programName string, domainName string, showOutOfScope bool) ([]Subdomain, error) {
	rows, err := database.DB.Query(`SELECT s.*,
										(SELECT d.[DomainName]
										FROM Domain d
										WHERE d.[Id] = s.[DomainId]) AS DomainName
									FROM [Subdomain] s
									JOIN [Domain] d on s.[DomainId] = d.[Id]
									JOIN [Program] p on d.[ProgramId] = p.[Id]
									WHERE (s.[IsInScope] = 1 OR ? = 1) AND
										s.[IsActive] = 1 AND
										(p.[DisplayName] = ? OR d.[DomainName] = ?)
									ORDER BY [SubdomainName] ASC`, showOutOfScope, programName, domainName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subdomains []Subdomain

	for rows.Next() {
		var subdomain Subdomain

		err := rows.Scan(&subdomain.ID, &subdomain.DomainId, &subdomain.SubdomainName, &subdomain.IsInScope, &subdomain.IsActive, &subdomain.DomainName)
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

func AddSubdomain(subdomain Subdomain) (err error, isInserted bool) {
	err = database.DB.QueryRow(`SELECT [Id] 
								FROM [Domain]
								WHERE [DomainName] = ?`, subdomain.DomainName).Scan(&subdomain.DomainId)
	if err != nil {
		return err, false
	}
	if subdomain.DomainId == 0 {
		return fmt.Errorf("Domain not found"), false
	}

	// Insert domain if not exists, else set to active
	row := database.DB.QueryRow(`
		DECLARE @IsInserted BIT
		IF NOT EXISTS (
			SELECT 1
			FROM [Subdomain]
			WHERE [SubdomainName] = ?)
		BEGIN
			INSERT INTO [Subdomain] ([DomainId], [SubdomainName], [IsInScope]) VALUES (?, ?, ?);
			SET @IsInserted = 1;
		END
		ELSE 
		BEGIN
			UPDATE [Subdomain]
				SET [IsInScope] = ?
				WHERE [SubdomainName] = ?;
			SET @IsInserted = 0;
		END
		SELECT @IsInserted`, subdomain.SubdomainName, subdomain.DomainId, subdomain.SubdomainName, subdomain.IsInScope, subdomain.IsInScope, subdomain.SubdomainName)

	err = row.Scan(&isInserted)
	if err != nil {
		return err, false
	}

	fmt.Println("isInserted", isInserted)

	return nil, isInserted
}

func RemoveSubdomain(subdomain Subdomain) error {
	_, err := database.DB.Exec("UPDATE [Subdomain] SET [IsActive] = 0 WHERE [SubdomainName] = ?", subdomain.SubdomainName)
	if err != nil {
		return err
	}

	return nil
}
