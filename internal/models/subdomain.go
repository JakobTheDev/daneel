package models

type Subdomain struct {
	// Database fields
	ID        int64
	DomainId  uint64
	Name      string
	IsInScope bool
	IsActive  bool
	// Extra fields
	DomainName string
}
