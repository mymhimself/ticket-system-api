package enum

type Role uint

const (
	User Role = iota
	Admin
)

func (r Role) String() string {
	return []string{"User", "Admin"}[r]
}
