package holiday

import "time"

type country int

const (
	unknownCountry country = iota
	usa
	canada
)

var tzToCountry = map[string]country{
	"America/New_York": usa,
	"America/Toronto":  canada,
}

func countryFromTZ(loc *time.Location) (c country) {
	var ok bool
	if c, ok = tzToCountry[loc.String()]; !ok {
		c = unknownCountry
	}
	return c
}
