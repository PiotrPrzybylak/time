package time

import (
	"testing"
)

func TestContains(t *testing.T) {

	var tests = []TimeSpanTest{
		{"10:00-11:22", "2001-10-01 10:00", "2001-10-01 11:22", true},
		{"10:00-11:22", "2001-10-01 10:20", "2001-10-01 10:50", true},
		{"10:00-11:22", "2001-10-01 10:00", "2001-10-01 10:30", true},
		{"10:10-11:22", "2001-10-01 10:30", "2001-10-01 11:22", true},
		{"10:10-11:22", "2001-10-01 09:59", "2001-10-01 11:22", false},
		{"10:10-11:22", "2001-10-01 10:00", "2001-10-01 11:23", false},
		{"10:10-11:22", "2001-10-01 05:00", "2001-10-01 06:00", false},
		{"10:00-00:00", "2001-10-01 10:00", "2001-10-01 11:22", true},
		{"10:00-00:00", "2001-10-01 10:20", "2001-10-01 10:50", true},
		{"10:00-00:00", "2001-10-01 10:00", "2001-10-01 10:30", true},
		{"10:10-00:00", "2001-10-01 10:30", "2001-10-01 11:22", true},
		{"10:10-00:00", "2001-10-01 10:30", "2001-10-02 00:00", true},
		{"10:10-00:00", "2001-10-01 10:30", "2001-10-02 00:01", false},
		{"10:10-00:00", "2001-10-01 09:59", "2001-10-01 11:22", false},
		{"10:10-00:00", "2001-10-01 10:00", "2001-10-01 11:23", false},
		{"10:10-00:00", "2001-10-01 05:00", "2001-10-01 06:00", false},
	}

	for _, test := range tests {
		testTimeSpanMethod(t, test, LocalTimeSpan.Contains)
	}

}

func TestOverlapsDateTimeSpan(t *testing.T) {

	var tests = []TimeSpanTest{
		{"10:00-11:22", "2001-10-01 10:00", "2001-10-01 11:22", true},
		{"10:00-11:22", "2001-10-01 10:20", "2001-10-01 10:50", true},
		{"10:00-11:22", "2001-10-01 10:00", "2001-10-01 10:30", true},
		{"10:10-11:22", "2001-10-01 10:30", "2001-10-01 11:22", true},
		{"10:10-11:22", "2001-10-01 09:59", "2001-10-01 11:22", true},
		{"10:10-11:22", "2001-10-01 10:00", "2001-10-01 11:23", true},
		{"10:10-11:22", "2001-10-01 05:00", "2001-10-01 06:00", false},
		{"10:10-11:22", "2001-10-01 05:00", "2001-10-01 10:10", false},
		{"10:00-00:00", "2001-10-01 10:00", "2001-10-01 11:22", true},
		{"10:00-00:00", "2001-10-01 10:20", "2001-10-01 10:50", true},
		{"10:00-00:00", "2001-10-01 10:00", "2001-10-01 10:30", true},
		{"10:10-00:00", "2001-10-01 10:30", "2001-10-01 11:22", true},
		{"10:10-00:00", "2001-10-01 09:59", "2001-10-01 11:22", true},
		{"10:10-00:00", "2001-10-01 10:00", "2001-10-01 11:23", true},
		{"10:10-00:00", "2001-10-01 05:00", "2001-10-01 06:00", false},
		{"10:10-00:00", "2001-10-01 05:00", "2001-10-01 06:00", false},
	}

	for _, test := range tests {

		testTimeSpanMethod(t, test, LocalTimeSpan.OverlapsDateTimeSpan)

	}

}

func testTimeSpanMethod(
	t *testing.T,
	test TimeSpanTest,
	timeSpanMethod func(timeSpan LocalTimeSpan, dateTimeSpan DateTimeSpan) bool,
) {
	timeSpan := MustParseTimeSpan(test.timeSpan)
	datetimeStart, _ := ParseLocalDateTime(test.datetimeStart)
	datetimeEnd, _ := ParseLocalDateTime(test.datetimeEnd)
	dateTimeSpan := NewDateTimeSpan(datetimeStart, datetimeEnd)
	if timeSpanMethod(timeSpan, dateTimeSpan) != test.want {
		t.Errorf(
			"Wanted : %v "+
				"for timeSpan: %v,"+
				" datetimeStart: %v,"+
				" datetimeEnd: %v",
			test.want,
			test.timeSpan,
			test.datetimeStart,
			test.datetimeEnd)
	}
}

type TimeSpanTest struct {
	timeSpan      string
	datetimeStart string
	datetimeEnd   string
	want          bool
}
