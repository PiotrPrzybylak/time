package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	before := NewLocalDateTime(time.Date(2017, 1, 1, 10, 30, 0, 0, time.UTC))
	after := NewLocalDateTime(time.Date(2017, 1, 1, 22, 45, 0, 0, time.UTC))

	assert.Equal(t, after, before.Add(12*Hour+15*Minute))
}

func TestSub(t *testing.T) {

	before := NewLocalDateTime(time.Date(2017, 1, 1, 10, 30, 0, 0, time.UTC))
	after := NewLocalDateTime(time.Date(2017, 1, 1, 22, 45, 0, 0, time.UTC))

	assert.Equal(t, 12*Hour+15*Minute, after.Sub(before))
}

func TestUnmarshalJSONOnLocalDateTime(t *testing.T) {

	expected := NewLocalDateTime(time.Date(2016, 1, 1, 12, 30, 0, 0, time.UTC))
	json := `"2016-01-01 12:30"`

	var actual LocalDateTime
	error := actual.UnmarshalJSON([]byte(json))
	assert.Nil(t, error)
	assert.Equal(t, expected, actual)

	error = actual.UnmarshalJSON([]byte(`"2016-01-01 12:30:12"`))
	assert.IsType(t, &time.ParseError{}, error)
}
