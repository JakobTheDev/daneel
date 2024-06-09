package domain

type Domain struct {
	// Database fields
	ID         int64
	ProgramId  int64
	DomainName string
	IsInScope  bool
	IsActive   bool
	// Extra fields
	ProgramName string
}
