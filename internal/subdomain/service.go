package subdomain

import (
	"fmt"
	"log"

	"github.com/JakobTheDev/daneel/internal/domain"
	"github.com/JakobTheDev/daneel/internal/models"
	"github.com/JakobTheDev/daneel/internal/tools"
)

func ListSubdomains(programName string, domainName string, showOutOfScope bool) (subdomains []models.Subdomain, err error) {
	subdomains, err = GetSubdomainListFromDb(programName, domainName, showOutOfScope)
	if err != nil {
		return nil, fmt.Errorf("failed to get subdomain list: %w", err)
	}

	return subdomains, nil
}

func AddSubdomain(subdomain models.Subdomain) (isInserted bool, err error) {
	isInserted, err = AddSubdomainToDb(subdomain)
	if err != nil {
		return false, fmt.Errorf("failed to add subdomain: %w", err)
	}
	return isInserted, nil
}

func RemoveSubdomain(subdomain models.Subdomain) (err error) {
	err = RemoveSubdomainFromDb(subdomain)
	if err != nil {
		return fmt.Errorf("failed to remove subdomain: %w", err)
	}
	return nil
}

func EnumerateSubdomainsByDomain(domainName string) (subdomains []string, newSubdomains []string, err error) {
	var isInserted bool

	// Check the domain exists
	var d models.Domain
	d, err = domain.GetDomain(domainName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get domain: %v", err)
	}

	// Enumerate subdomains for the domain
	subdomains, err = tools.RunSubfinder(d.Name)
	if err != nil {
		return nil, nil, fmt.Errorf("error running subfinder: %v", err)
	}

	if len(subdomains) == 0 {
		log.Println("No subdomains found")
		return
	}

	log.Printf("Found %d subdomains\n", len(subdomains))

	// Insert into database
	for _, s := range subdomains {
		isInserted, err = AddSubdomainToDb(models.Subdomain{DomainName: d.Name, Name: s})
		if err != nil {
			log.Printf("Failed to subdomain %s to database: %v\n", s, err)
		}
		if isInserted {
			newSubdomains = append(newSubdomains, s)
		}
	}

	return subdomains, newSubdomains, nil

}
