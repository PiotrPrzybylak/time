package mocks

import (
	"context"
	"github.com/RnDity/time"
	gotime "time"
)

type StubClock struct {
	Date          time.LocalDate
	DateTime      time.LocalDateTime
	WallTime      gotime.Time
	UnixTimeStamp time.UnixTimeStamp
}

func (c StubClock) Today(_ context.Context) time.LocalDate {
	return c.Date
}

func (c StubClock) Now(_ context.Context) time.LocalDateTime {
	return c.DateTime
}

func (c StubClock) GoTime(_ context.Context) gotime.Time {
	return c.WallTime
}

func (c StubClock) ToUnixTimestamp(_ context.Context, _ time.LocalDateTime) time.UnixTimeStamp {
	return 0
}
