package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLocalTimeEquality(t *testing.T) {
	time1, _ := NewLocalTime(11, 22)
	time2, _ := NewLocalTime(11, 22)
	time3, _ := NewLocalTime(11, 23)
	assert.Equal(t, time1, time2)
	assert.NotEqual(t, time1, time3)
}

func TestMarshalJSONOnLocalTime(t *testing.T) {
	time1, _ := NewLocalTime(12, 34)
	JSON, _ := time1.MarshalJSON()
	assert.Equal(t, []byte(`"12:34"`), JSON)

	time2, _ := NewLocalTime(1, 1)
	JSON2, _ := time2.MarshalJSON()
	assert.Equal(t, []byte(`"01:01"`), JSON2)
}

func TestUnmarshalJSONOnLocalTime(t *testing.T) {
	time1 := NullLocalTime
	var time1p = &time1
	err := time1p.UnmarshalJSON([]byte(`"23:45"`))
	assert.Nil(t, err)
	expected, _ := NewLocalTime(23, 45)
	assert.Equal(t, expected, time1)
}

func TestScanLocalTime(t *testing.T) {
	time1 := NullLocalTime
	var time1p = &time1
	err := time1p.Scan(time.Date(0, 0, 0, 23, 45, 0, 0, time.UTC))
	assert.Nil(t, err)
	expected, _ := NewLocalTime(23, 45)
	assert.Equal(t, expected, time1)
}
