package constants

type UserRole string

const (
	UserRoleUser  UserRole = "user"
	UserRoleAdmin UserRole = "admin"
)

type SortOrder string

const (
	ASC  SortOrder = "asc"
	DESC SortOrder = "desc"
)
