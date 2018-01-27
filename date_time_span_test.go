package time

import (
	"testing"
)

func TestDateTimeSpan_Overlaps(t *testing.T) {
	var tests = []DateTimeSpanTest{
		{
			"2017-10-04 09:00", "2017-10-04 20:22",
			"2017-10-04 10:00", "2017-10-04 11:22",
			true},
		{
			"2017-10-03 09:00", "2017-10-03 20:22",
			"2017-01-04 10:50", "2017-01-04 23:00",
			false},
		{
			"2017-01-04 09:00", "2017-01-04 10:00",
			"2017-01-04 09:30", "2017-01-04 10:30",
			true},
		{
			"2017-01-04 09:00", "2017-01-04 10:00",
			"2017-01-04 10:00", "2017-01-04 11:00",
			false},
	}
	for _, test := range tests {
		testDateTimeSpanMethod(t, test, DateTimeSpan.Overlaps)
	}
}

func TestDateTimeSpanContains(t *testing.T) {
	var tests = []DateTimeSpanTest{
		{
			"2017-01-04 09:00", "2017-01-04 10:00",
			"2017-01-04 09:00", "2017-01-04 10:00",
			true,
		},
		{
			"2017-10-04 08:00", "2017-10-04 10:00",
			"2017-10-04 09:00", "2017-10-04 10:00",
			true,
		},
		{
			"2017-01-04 10:50", "2017-01-04 23:00",
			"2017-10-03 09:00", "2017-10-03 10:00",
			false,
		},
		{
			"2017-01-04 09:30", "2017-01-04 10:30",
			"2017-01-04 09:00", "2017-01-04 10:00",
			false,
		},
	}
	for _, test := range tests {
		testDateTimeSpanMethod(t, test, DateTimeSpan.Contains)
	}
}

func testDateTimeSpanMethod(
	t *testing.T,
	test DateTimeSpanTest,
	dateTimeSpanMethod func(datetimeSpan DateTimeSpan, dateTimeSpan DateTimeSpan) bool,
) {
	datetimeStart, _ := ParseLocalDateTime(test.datetimeStart)
	datetimeEnd, _ := ParseLocalDateTime(test.datetimeEnd)
	datetimeStartOther, _ := ParseLocalDateTime(test.datetimeStartOther)
	datetimeEndOther, _ := ParseLocalDateTime(test.datetimeEndOther)
	dateTimeSpan := NewDateTimeSpan(datetimeStart, datetimeEnd)
	dateTimeSpanOther := NewDateTimeSpan(datetimeStartOther, datetimeEndOther)
	if dateTimeSpanMethod(dateTimeSpan, dateTimeSpanOther) != test.want {
		t.Errorf(
			"Wanted : %v "+
				"for datetimeStart: %v,"+
				"datetimeEnd: %v,"+
				"datetimeStartOther: %v,"+
				"datetimeEndOther: %v,",
			test.want,
			test.datetimeStart,
			test.datetimeEnd,
			test.datetimeStartOther,
			test.datetimeEndOther)
	}
}

type DateTimeSpanTest struct {
	datetimeStart      string
	datetimeEnd        string
	datetimeStartOther string
	datetimeEndOther   string
	want               bool
}
