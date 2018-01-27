package mocks

import (
	"github.com/RnDity/time"
)

func Period(from, to string) time.Period {
	period, err := time.NewPeriod(Date(from), Date(to))
	if err != nil {
		panic(err)
	}
	return period
}

func Date(date string) time.LocalDate {
	return time.MustParseLocalDate(date)
}
