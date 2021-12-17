package enum

import "errors"

type Priority int64

const (
	PriorityLevelLow = iota + 1
	PriorityLevelNormal
	PriorityLevelHigh
	PriorityLevelCritical
)

func (p Priority) String() string {
	if val, ok := priorityStringMap[p]; ok {
		return val
	}
	return stringMap[PriorityLevelLow]
}

func IsValidValue(priority uint) bool {
	_, ok := priorityStringMap[Priority(priority)]
	return ok
}

var priorityStringMap map[Priority]string = map[Priority]string{
	PriorityLevelLow:      "Low",
	PriorityLevelNormal:   "Normal",
	PriorityLevelHigh:     "High",
	PriorityLevelCritical: "Critical",
}

func (s *Priority) UnmarshalText(b []byte) error {
	for i, v := range priorityStringMap {
		if v == string(b) {
			*s = Priority(i)
			return nil
		}
	}
	return errors.New("invalid priority value")
}
