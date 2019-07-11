package dough

import (
	"testing"
	"time"
)

func TestGetTimeZoneDetails(t *testing.T) {
	zone, offsetIsHourly, offsetInMinutes, isDaylighSavings, err := GetTimeZoneDetails("US/Arizona")
	if err != nil {
		t.Error(err)
	}
	if zone != "MST" {
		t.Errorf("Zone should be MST, instead of %s", zone )
	}
	if offsetIsHourly != true {
		t.Error("offsetIsHourly should be true")
	}
	if offsetInMinutes != -420 {
		t.Errorf("Zone should be -420, instead of %d", offsetInMinutes )
	}
	if isDaylighSavings != false {
		t.Error("isDaylighSavings should be false")
	}

	_, _, _, _, err = GetTimeZoneDetails("invalid")
	if err == nil {
		t.Errorf("Error should be 'unknown time zone invalid'")
	}

	_, _, _, isDaylighSavings, err = GetTimeZoneDetails("US/Alaska")
	if err != nil {
		t.Error(err)
	}

	if isDaylighSavings != true {
		t.Error("isDaylighSavings should be true")
	}
}

func TestParseTimeToExp(t *testing.T) {
	scenarios := map[string]string{
		"2012-11-01T00:00:00+00:00":"11/12",
		"2015-12-01T00:00:00+00:00":"12/15",
		"2013-04-01T00:00:00+00:00":"04/13",
		"2011-11-01T00:00:00+00:00":"11/11",
		"2020-12-01T00:00:00+00:00":"12/20",
		"2022-06-01T00:00:00+00:00":"06/22",
		"2022-11-01T00:00:00+00:00":"11/22",
	}

	for k, v := range scenarios {
		t1, err := time.Parse(
			time.RFC3339,
			k)
		if err != nil {
			t.Error(err)
		}
		result := ParseTimeToExp(t1)

		if result != v {
			t.Errorf("Parsed time should be %s, instead of %s", v, result)
		}
	}
}

func TestParseExpToTime(t *testing.T) {
	scenarios := map[string]string{
		"2012-11-30T00:00:00Z":"11/12",
		"2015-12-31T00:00:00Z":"12/15",
		"2013-04-30T00:00:00Z":"04/13",
		"2011-11-30T00:00:00Z":"11/11",
		"2020-12-31T00:00:00Z":"12/20",
		"2022-06-30T00:00:00Z":"06/22",
		"2022-11-30T00:00:00Z":"1122",
	}

	for k, v := range scenarios {
		result, err := ParseExpToTime(v)
		if err != nil {
			t.Error(err)
		}

		t1, err := time.Parse(
			time.RFC3339,
			k)
		if err != nil {
			t.Error(err)
		}

		if result != t1 {
			t.Errorf("Parsed time should be %v, instead of %v", t1, result)
		}
	}

	_, err := ParseExpToTime("")
	if err != ErrInvalidExpDateLen {
		t.Errorf("Error should be %v, instead of %v", ErrInvalidExpDateLen, err)
	}

	_, err = ParseExpToTime("15/12")
	if err != ErrInvalidMonth {
		t.Errorf("Error should be %v, instead of %v", ErrInvalidMonth, err)
	}

	_, err = ParseExpToTime("11/60")
	if err != ErrInvalidYear {
		t.Errorf("Error should be %v, instead of %v", ErrInvalidYear, err)
	}
}