package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalTextToWeekday(t *testing.T) {

	var w Weekday

	w.UnmarshalText([]byte("monday"))
	assert.Equal(t, Monday, w)

	w.UnmarshalText([]byte("tuesday"))
	assert.Equal(t, Tuesday, w)

	w.UnmarshalText([]byte("wednesday"))
	assert.Equal(t, Wednesday, w)

	w.UnmarshalText([]byte("Thursday"))
	assert.Equal(t, Thursday, w)

	w.UnmarshalText([]byte("Friday"))
	assert.Equal(t, Friday, w)

	w.UnmarshalText([]byte("Saturday"))
	assert.Equal(t, Saturday, w)

	w.UnmarshalText([]byte("sunday"))
	assert.Equal(t, Sunday, w)

	w.UnmarshalText([]byte("Sunday"))
	assert.Equal(t, Sunday, w)

	w.UnmarshalText([]byte("SUNDAY"))
	assert.Equal(t, Sunday, w)

	w.UnmarshalText([]byte("sUnDAY"))
	assert.Equal(t, Sunday, w)

	err := w.UnmarshalText([]byte("not a weekday"))
	assert.Error(t, err)
	assert.Equal(t, NotAWeekday, w)

}

func TestConversionToTimeWeekday(t *testing.T) {
	assert.Equal(t, Monday, Weekday(time.Monday))
	assert.Equal(t, time.Tuesday, time.Weekday(Tuesday))
}
