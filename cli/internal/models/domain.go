package models

import (
	"github.com/JakobTheDev/daneel/internal/database"
)

type Domain struct {
	ID          int64
	ProgramId   int64
	ProgramName string
	DomainName  string
	IsInScope   bool
	IsActive    bool
}

func ListDomain(programName string, showOutOfScope bool) ([]Domain, error) {
	rows, err := database.DB.Query(`SELECT d.*,
										(SELECT p.[DisplayName]
										FROM Program p
										WHERE p.[Id] = d.[ProgramId]) AS ProgramName
									FROM [Domain] d
									JOIN [Program] p on d.[ProgramId] = p.[Id]
									WHERE ([IsInScope] = 1 OR ? = 1) AND
										d.[IsActive] = 1 AND
										(p.[DisplayName] = ? OR ? = '')
									ORDER BY [DomainName] ASC`, showOutOfScope, programName, programName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var domains []Domain

	for rows.Next() {
		var domain Domain

		err := rows.Scan(&domain.ID, &domain.ProgramId, &domain.DomainName, &domain.IsInScope, &domain.IsActive, &domain.ProgramName)
		if err != nil {
			return nil, err
		}

		domains = append(domains, domain)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return domains, nil
}

func AddDomain(domain Domain) error {
	var err error

	err = database.DB.QueryRow(`SELECT [Id] 
								FROM [Program]
								WHERE [DisplayName] = ?`, domain.ProgramName).Scan(&domain.ProgramId)
	if err != nil {
		return err
	}
	if domain.ProgramId == 0 {
		return nil
	}

	// Insert domain if not exists, else set to active
	_, err = database.DB.Exec(`
		IF NOT EXISTS (
			SELECT 1 
			FROM [Domain] 
			WHERE [DomainName] = ?) 
		INSERT INTO [Domain] ([ProgramId], [DomainName], [IsInScope]) VALUES (?, ?, ?) 
		ELSE UPDATE [Domain] 
			 SET [IsInScope] = ? 
			 WHERE [DomainName] = ?`, domain.DomainName, domain.ProgramId, domain.DomainName, domain.IsInScope, domain.IsInScope, domain.DomainName)
	if err != nil {
		return err
	}

	return nil
}

func RemoveDomain(d Domain) error {
	_, err := database.DB.Exec("UPDATE [Domain] SET [IsActive] = 0 WHERE [Domain] = ?", d.DomainName)
	if err != nil {
		return err
	}

	return nil
}
