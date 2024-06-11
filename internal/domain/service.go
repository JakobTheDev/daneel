package domain

import (
	"fmt"

	"github.com/JakobTheDev/daneel/internal/models"
)

func GetDomain(domainName string) (domain models.Domain, err error) {
	domain, err = GetDomainFromDb(domainName)
	if err != nil {
		return models.Domain{}, fmt.Errorf("failed to get domain: %w", err)
	}
	return domain, nil
}

func ListDomains(programName string, showOutOfScope bool) (domains []models.Domain, err error) {
	domains, err = GetDomainListFromDb(programName, showOutOfScope)
	if err != nil {
		return nil, fmt.Errorf("failed to get domain list: %w", err)
	}
	return domains, nil
}

func AddDomain(domain models.Domain) (err error) {
	err = AddDomainToDb(domain)
	if err != nil {
		return fmt.Errorf("failed to add domain: %w", err)
	}
	return nil
}

func RemoveDomain(domain models.Domain) (err error) {
	err = RemoveDomainFromDb(domain)
	if err != nil {
		return fmt.Errorf("failed to remove domain: %w", err)
	}
	return nil
}
