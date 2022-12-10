package holiday

// Week implies which week the holiday usually falls on.
// Eg. Memorial Day falls on the last monday of the month.
// - or -
// Thanksgiving falls on the fourth Thursday of the month.
type week int

const (
	noWeek week = iota
	firstWeek
	secondWeek
	thirdWeek
	fourthWeek
	fifthWeek
	secondToLastWeek
	lastWeek
)

func (w week) String() string {
	var s string
	switch w {
	case firstWeek:
		s = "first week"
	case secondWeek:
		s = "second week"
	case thirdWeek:
		s = "third week"
	case fourthWeek:
		s = "fourth week"
	case fifthWeek:
		s = "fifth week"
	case secondToLastWeek:
		s = "second to last week"
	case lastWeek:
		s = "last week"
	}
	return s
}
