package domain

import (
	"github.com/JakobTheDev/daneel/internal/database"
	"github.com/JakobTheDev/daneel/internal/models"
)

func GetDomainFromDb(domainName string) (domain models.Domain, err error) {
	err = database.DB.QueryRow(`SELECT d.*,
									(SELECT p.[Name]
									FROM Program p
									WHERE p.[Id] = d.[ProgramId]) AS ProgramName
								FROM [Domain] d
								WHERE [Name] = ?`, domainName).Scan(&domain.ID, &domain.ProgramId, &domain.Name, &domain.IsInScope, &domain.IsActive, &domain.ProgramName)
	if err != nil {
		return domain, err
	}

	return domain, nil
}

func GetDomainListFromDb(programName string, showOutOfScope bool) ([]models.Domain, error) {
	rows, err := database.DB.Query(`SELECT d.*,
										(SELECT p.[Name]
										FROM Program p
										WHERE p.[Id] = d.[ProgramId]) AS ProgramName
									FROM [Domain] d
									JOIN [Program] p on d.[ProgramId] = p.[Id]
									WHERE ([IsInScope] = 1 OR ? = 1) AND
										d.[IsActive] = 1 AND
										(p.[Name] = ? OR ? = '')
									ORDER BY [Name] ASC`, showOutOfScope, programName, programName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var domains []models.Domain

	for rows.Next() {
		var domain models.Domain

		err := rows.Scan(&domain.ID, &domain.ProgramId, &domain.Name, &domain.IsInScope, &domain.IsActive, &domain.ProgramName)
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

func AddDomainToDb(domain models.Domain) error {
	var err error

	err = database.DB.QueryRow(`SELECT [Id] 
								FROM [Program]
								WHERE [Name] = ?`, domain.ProgramName).Scan(&domain.ProgramId)
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
			WHERE [Name] = ?) 
		INSERT INTO [Domain] ([ProgramId], [Name], [IsInScope]) VALUES (?, ?, ?) 
		ELSE UPDATE [Domain] 
			 SET [IsInScope] = ? 
			 WHERE [Name] = ?`, domain.Name, domain.ProgramId, domain.Name, domain.IsInScope, domain.IsInScope, domain.Name)
	if err != nil {
		return err
	}

	return nil
}

func RemoveDomainFromDb(domain models.Domain) error {
	_, err := database.DB.Exec("UPDATE [Domain] SET [IsActive] = 0 WHERE [Name] = ?", domain.Name)
	if err != nil {
		return err
	}

	return nil
}
