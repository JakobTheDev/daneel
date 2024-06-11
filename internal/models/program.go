package models

type Program struct {
	// Database fields
	Id         int64
	PlatformId int64
	Name       string
	IsPrivate  bool
	IsActive   bool
	// Extra fields
	PlatformName string
}
