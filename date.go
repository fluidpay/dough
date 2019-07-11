package dough

import (
	"errors"
	"strconv"
	"time"
)

var (
	ErrInvalidExpDateLen = errors.New("Invalid expiration date length")
	ErrInvalidMonth = errors.New("Invalid month in expiration date")
	ErrInvalidYear = errors.New("Invalid year in expiration date")
)

// ParseExpToTime takes a exp_date in MMYY format and returns time.time
func ParseExpToTime(exp string) (time.Time, error) {
	var t time.Time
	length := len(exp)
	if length != 4 && length != 5 {
		return t, ErrInvalidExpDateLen
	}
	var month int
	var year int
	if length == 4 {
		month, _ = strconv.Atoi(exp[:2])
		year, _ = strconv.Atoi(exp[2:4])
	}
	if length == 5 {
		month, _ = strconv.Atoi(exp[:2])
		year, _ = strconv.Atoi(exp[3:5])
	}

	if month < 1 || month > 12 {
		return t, ErrInvalidMonth
	}

	year = year + 2000
	if year <= 2000 || year > 2050 {
		return t, ErrInvalidYear
	}

	t = time.Date(year, time.Month(month), (time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()), 0, 0, 0, 0, time.UTC)
	return t, nil
}

// ParseTimeToExp takes a time.time and returns exp_date in MMYY format
func ParseTimeToExp(t time.Time) string {
	return t.Format("01/06")
}

// GetTimeZoneDetails takes in a location and returns zone, offsetIsHourly, offsetInMinutes, isDaylighSavings
func GetTimeZoneDetails(timezone string) (string, bool, int, bool, error) {
	daylightSavingTimeZones := []string{"ACDT", "ADT", "AEDT", "AKDT", "AMST", "ANAST", "AWDT", "AZOST", "AZST", "BRST", "BST", "CDT", "CEST", "CHADT", "CHOST", "CIDST", "CKST", "CLST", "EASST", "EDT", "EEST", "EGST", "FJST", "FKST", "HADTMDT", "HOVST", "IDT", "IRDT", "IRKST", "KRAST", "LHDT", "MAGST", "MSD", "NDT", "NOVST", "NZDT", "OMSST", "PDT", "PETST", "PMDT", "PYST", "RKST", "TOST", "ULAST", "UYST", "VLAST", "WARST", "WAST", "WEST", "WGST", "WST", "YAKST", "YEKST"}
	var offsetIsHourly = false
	var offsetInMinutes int
	var isDaylighSavings = false

	now := time.Now().UTC()

	location, err := time.LoadLocation(timezone)

	if err != nil {
		return "", false, 0, false, err
	}

	zone, offset := now.In(location).Zone()

	offsetInMinutes = offset / 60

	if offset%60 == 0 {
		offsetIsHourly = true
	}

	for _, v := range daylightSavingTimeZones {
		if v == zone {
			isDaylighSavings = true
			break
		}
	}

	return zone, offsetIsHourly, offsetInMinutes, isDaylighSavings, nil
}
