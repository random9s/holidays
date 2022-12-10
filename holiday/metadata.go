package holiday

import "time"

type metadata struct {
	Month   time.Month
	Week    week
	Weekday time.Weekday
	Day     int
	Year    int
}

//NullMetadata ...
var NullMetadata = metadata{}

func (m metadata) String() string {
	return m.Date().String()
}

func (m metadata) Date() time.Time {
	return time.Date(m.Year, m.Month, m.Day, 0, 0, 0, 0, time.Now().Location())
}
