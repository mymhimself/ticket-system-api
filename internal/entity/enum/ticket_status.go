package enum

import "errors"

type TicketStatus uint

const (
	Active = iota
	All    //all status
	Replied
	Closed
	Unread
)

var stringMap []string = []string{"active", "replied", "closed"}

func (s TicketStatus) String() string {
	return stringMap[s]
}

func (s *TicketStatus) UnmarshalText(b []byte) error {
	for i, v := range stringMap {
		if v == string(b) {
			*s = TicketStatus(i)
			return nil
		}
	}
	return errors.New("invalid ticket status")
}
