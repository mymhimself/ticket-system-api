package enum

import "errors"

type TicketThreadStatus uint

const (
	Active TicketThreadStatus = iota
	Suspend
	Closed
)

var stringMap []string = []string{"active", "closed"}

func (s TicketThreadStatus) String() string {
	return stringMap[s]
}

func (s *TicketThreadStatus) UnmarshalText(b []byte) error {
	for i, v := range stringMap {
		if v == string(b) {
			*s = TicketThreadStatus(i)
			return nil
		}
	}
	return errors.New("invalid ticket status")
}
