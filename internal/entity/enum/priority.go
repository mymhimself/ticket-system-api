package enum

type Priority uint

const (
	Normal = iota
	Urgent
)

func (p Priority) String() string {
	return []string{"Normal", "Urgent"}[p]
}
