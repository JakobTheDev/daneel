package models

type Domain struct {
	// Database fields
	ID        int64
	ProgramId int64
	Name      string
	IsInScope bool
	IsActive  bool
	// Extra fields
	ProgramName string
}
