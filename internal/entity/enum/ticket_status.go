package enum

import (
	"errors"
	"fmt"
	"strings"
)

type TicketThreadStatus uint

const (
	Active TicketThreadStatus = iota
	Suspend
	Closed
)

var ticketStatusStringMap map[TicketThreadStatus]string = map[TicketThreadStatus]string{
	Active:  "Active",
	Closed:  "Closed",
	Suspend: "Suspend",
}

func (s TicketThreadStatus) String() string {
	val, ok := ticketStatusStringMap[s]
	if ok {
		return val
	}
	return fmt.Sprintf("TicketStatus(%d)", s)
}

func (s *TicketThreadStatus) UnmarshalText(b []byte) error {
	for i, v := range ticketStatusStringMap {
		if strings.ToLower(v) == strings.ToLower(string(b)) {
			*s = TicketThreadStatus(i)
			return nil
		}
	}
	return errors.New("invalid ticket status")
}

func (s TicketThreadStatus) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}
