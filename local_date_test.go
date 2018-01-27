package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, NewLocalDate(2000, 1, 1), NewLocalDate(2000, 1, 1))
	assert.NotEqual(t, NewLocalDate(2000, 1, 1), NewLocalDate(2000, 1, 2))
}

func TestConversionToTime(t *testing.T) {
	assert.Equal(t,
		NewLocalDate(2005, 6, 7).GetStartOfDayUTC(),
		time.Date(2005, 6, 7, 0, 0, 0, 0, time.UTC))
}

func TestMarshalText(t *testing.T) {
	d, _ := NewLocalDate(22345, 01, 01).MarshalText()
	assert.Equal(t, []byte("22345-01-01"), d)
}

func TestMarshalJSON(t *testing.T) {
	d, _ := NewLocalDate(22345, 01, 01).MarshalJSON()
	assert.Equal(t, []byte(`"22345-01-01"`), d)
}

func TestUnmarshalText(t *testing.T) {
	d := NullLocalDate
	var dp = &d
	err := dp.UnmarshalText([]byte("2000-09-10"))
	assert.Nil(t, err)
	assert.Equal(t, NewLocalDate(2000, 9, 10), d)
}

func TestUnmarshalJSON(t *testing.T) {
	d := NullLocalDate
	var dp = &d
	err := dp.UnmarshalJSON([]byte(`"2000-09-10"`))
	assert.Nil(t, err)
	assert.Equal(t, NewLocalDate(2000, 9, 10), d)
}
