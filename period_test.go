package time

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeriodString(t *testing.T) {
	period, _ := NewPeriod(NewLocalDate(2010, 11, 12), NewLocalDate(2011, 12, 13))
	assert.Equal(t, "[2010-11-12 - 2011-12-13]", period.String())
	period, _ = NewPeriod(NewLocalDate(10, 10, 10), NewLocalDate(11, 11, 11))
	assert.Equal(t, "[0010-10-10 - 0011-11-11]", period.String())
	period, _ = NewOpenPeriodFrom(NewLocalDate(2010, 11, 12))
	assert.Equal(t, "[2010-11-12 - indefinitely]", period.String())
	period, _ = NewOneDayPeriod(NewLocalDate(2010, 11, 12))
	assert.Equal(t, "[2010-11-12 - 2010-11-12]", period.String())
}
