//Package holiday contains algorithms to determine bank holidays for a given year
package holiday

import (
	"time"
)

/*
MLK Day – 3rd Monday of January
President’s Day – 3rd Monday of February
Memorial Day – Last Monday of May
Juneteenth – Always June 19th
Independence Day – Always July 4th
Labor Day – 1st Monday in September
Indigenous People’s Day – 2nd Monday in October
Veteran's Day – Always November 11th
Thanksgiving – 4th Thursday in November
Christmas – Always December 25th
New Years – Always January 1st
*/

// Holiday is an enumeration of all valid bank holidays globally
type Holiday int

const (
	MartinLutherKingDay Holiday = iota
	PresidentsDay
	MemorialDay
	Juneteenth
	IndependenceDay
	LaborDay
	IndigenousPeoplesDay
	VeteransDay
	ThanksgivingDay
	ChristmasDay
	NewYearsDay
	GoodFriday
	VictoriaDay
	CanadaDay
	LabourDay
	TruthAndReconciliationDay
	RemembranceDay
	BoxingDay
	GregorianEaster
)

func (h Holiday) String() string {
	var name string
	switch h {
	case MartinLutherKingDay:
		name = "Marting Luther King Jr Day"
	case PresidentsDay:
		name = "Presidents Day"
	case MemorialDay:
		name = "Memorial Day"
	case Juneteenth:
		name = "Juneteenth"
	case IndependenceDay:
		name = "Independence Day"
	case LaborDay:
		name = "Labor Day"
	case IndigenousPeoplesDay:
		name = "Indigenous People's Day"
	case VeteransDay:
		name = "Veterans Day"
	case ThanksgivingDay:
		name = "Thanksgiving Day"
	case ChristmasDay:
		name = "Christmas Day"
	case NewYearsDay:
		name = "New Years Day"
	case GoodFriday:
		name = "Good Friday"
	case VictoriaDay:
		name = "Victoria Day"
	case CanadaDay:
		name = "Canada Day"
	case LabourDay:
		name = "Labour Day"
	case TruthAndReconciliationDay:
		name = "Truth and Reconciliation Day"
	case RemembranceDay:
		name = "Remembrance Day"
	case BoxingDay:
		name = "Boxing Day"
	case GregorianEaster:
		name = "Easter"
	}
	return name
}

//Of the below, each should read the following way:
//  This holiday falls on the month provided
//  This holiday falls on the given week of the given month
//  This holiday falls on the given weekday of the given month
//  This holiday falls on the given day of the given month
// Ex:
// Memorial Day falls on the last Monday of May. Because the day is -1, no day is used to determine the actual date
// Juneteenth always falls on the 19th of June every year.
var usaObservedHolidays = map[Holiday]metadata{
	NewYearsDay:          {time.January, noWeek, time.Weekday(-1), 1, 0},
	MartinLutherKingDay:  {time.January, thirdWeek, time.Monday, -1, 0},
	PresidentsDay:        {time.February, thirdWeek, time.Monday, -1, 0},
	MemorialDay:          {time.May, lastWeek, time.Monday, -1, 0},
	Juneteenth:           {time.June, noWeek, time.Weekday(-1), 19, 0},
	IndependenceDay:      {time.July, noWeek, time.Weekday(-1), 4, 0},
	LaborDay:             {time.September, firstWeek, time.Monday, -1, 0},
	IndigenousPeoplesDay: {time.October, secondWeek, time.Monday, -1, 0},
	VeteransDay:          {time.November, noWeek, time.Weekday(-1), 11, 0},
	ThanksgivingDay:      {time.November, fourthWeek, time.Thursday, -1, 0},
	ChristmasDay:         {time.December, noWeek, time.Weekday(-1), 25, 0},
	GregorianEaster:      NullMetadata,
}

var canadaObservedHolidays = map[Holiday]metadata{
	NewYearsDay:               {time.January, noWeek, time.Weekday(-1), 1, 0},
	GoodFriday:                NullMetadata,
	VictoriaDay:               NullMetadata,
	CanadaDay:                 {time.July, noWeek, time.Weekday(-1), 1, 0},
	LabourDay:                 {time.September, firstWeek, time.Monday, -1, 0},
	TruthAndReconciliationDay: {time.September, noWeek, time.Weekday(-1), 30, 0},
	ThanksgivingDay:           {time.October, secondWeek, time.Monday, -1, 0},
	RemembranceDay:            {time.November, noWeek, time.Weekday(-1), 11, 0},
	ChristmasDay:              {time.December, noWeek, time.Weekday(-1), 25, 0},
	BoxingDay:                 {time.December, noWeek, time.Weekday(-1), 26, 0},
	GregorianEaster:           NullMetadata,
}

var customHolidayCalculation = map[Holiday]func(year int) time.Time{
	GregorianEaster: determineEasterDate,
	GoodFriday:      determineGoodFriday,
	VictoriaDay:     determineVictoriaDay,
}

func determineVictoriaDay(year int) time.Time {
	t1 := time.Date(year, time.May, 18, 0, 0, 0, 0, time.Now().Location())
	for {
		if t1.Weekday() == time.Monday {
			break
		}
		t1 = t1.AddDate(0, 0, 1)
	}

	return t1
}

func determineGoodFriday(year int) time.Time {
	easter := determineEasterDate(year)
	return easter.AddDate(0, 0, -2)
}

func determineEasterDate(year int) time.Time {
	a := year % 19

	b := year / 100
	c := year % 100

	d := b / 4
	e := b % 4

	f := (b + 8) / 25
	g := (b - f + 1) / 3

	h := ((19 * a) + b - d - g + 15) % 30

	i := c / 4
	k := c % 4

	l := (32 + (2 * e) + (2 * i) - h - k) % 7
	m := (a + (11 * h) + (22 * l)) / 451

	o := (h + l - (7 * m) + 114)

	n := o / 31
	p := o % 31

	return time.Date(year, time.Month(n), p+1, 0, 0, 0, 0, time.Now().Location())
}

func determineOccurencesOfWeekdayInMonth(weekday time.Weekday, month time.Month, year int) week {
	var occurs int
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().Location())
	for t := start; t.Month() == month; t = t.AddDate(0, 0, 1) {
		if t.Weekday() == weekday {
			occurs++
		}
	}
	return week(occurs)
}

func determineDayOfObserveration(weekday time.Weekday, w week, month time.Month, year int) int {
	var day int
	var occurs int

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().Location())
	for t := start; t.Month() == month; t = t.AddDate(0, 0, 1) {
		if t.Weekday() == weekday {
			occurs++
		}

		if occurs == int(w) {
			day = t.Day()
			break
		}
	}

	return day
}

func generateHolidays(observedHolidays map[Holiday]metadata, year int) map[Holiday]metadata {
	output := make(map[Holiday]metadata)
	for holiday, info := range observedHolidays {
		var curr time.Time
		if info == NullMetadata {
			curr = customHolidayCalculation[holiday](year)
		} else {
			info.Year = year
			if info.Week == lastWeek {
				info.Week = determineOccurencesOfWeekdayInMonth(info.Weekday, info.Month, info.Year)
			}

			if info.Day == -1 {
				info.Day = determineDayOfObserveration(info.Weekday, info.Week, info.Month, info.Year)
			}

			curr = time.Date(info.Year, info.Month, info.Day, 0, 0, 0, 0, time.Now().Location())
		}

		switch curr.Weekday() {
		case time.Saturday:
			curr = curr.AddDate(0, 0, -1) // Celebrate on Friday
			break
		case time.Sunday:
			curr = curr.AddDate(0, 0, 1) // Celebrate on Monday
			break
		}

		info.Month = curr.Month()
		info.Year = curr.Year()
		info.Day = curr.Day()
		output[holiday] = info
	}
	return output
}

// WhenIs reports when a given holiday will occur for the current year
func WhenIs(h Holiday, tz *time.Location) time.Time {
	t := time.Now()

	output := make(map[Holiday]metadata)
	c := countryFromTZ(tz)
	switch c {
	case usa:
		output = generateHolidays(usaObservedHolidays, t.Year())
	case canada:
		output = generateHolidays(canadaObservedHolidays, t.Year())
	}

	return output[h].Date()
}

// WhenIsInYear reports when a given holiday will occur for the current year
func WhenIsInYear(h Holiday, year int, tz *time.Location) time.Time {
	output := make(map[Holiday]metadata)
	c := countryFromTZ(tz)
	switch c {
	case usa:
		output = generateHolidays(usaObservedHolidays, year)
	case canada:
		output = generateHolidays(canadaObservedHolidays, year)
	}

	return output[h].Date()
}

// IsToday determines if the given holiday is occuring today
func IsToday(h Holiday, tz *time.Location) bool {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()) == WhenIs(h, tz)
}

// Generate creates the holiday map based on a given location and year
func Generate(tz *time.Location, year int) map[Holiday]time.Time {
	m := make(map[Holiday]time.Time)
	output := make(map[Holiday]metadata)

	c := countryFromTZ(tz)
	switch c {
	case usa:
		output = generateHolidays(usaObservedHolidays, year)
	case canada:
		output = generateHolidays(canadaObservedHolidays, year)
	}

	for k, v := range output {
		m[k] = v.Date()
	}

	return m
}
