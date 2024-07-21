package types

type Role int8

const (
	// NullRole indicates a unknown user role.
	NullRole Role = iota

	// UserRole indicates a standard user role.
	UserRole

	// AdminRole indicates an administrative user role.
	AdminRole
)

func (r Role) String() string {
	return [...]string{"null", "user", "admin"}[r]
}
