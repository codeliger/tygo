package simple

type UserRole = string

const (
	UserRoleDefault UserRole = "viewer"
	UserRoleEditor  UserRole = "editor" // Line comments are also kept
)
