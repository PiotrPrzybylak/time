package time

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeInWarsaw(t *testing.T) {
	t1, err := time.Parse(time.RFC3339, "2018-01-02T15:04:05Z")
	assert.NoError(t, err)
	t2, err := time.Parse(time.RFC3339, "2018-07-02T15:04:05Z")
	assert.NoError(t, err)
	warsaw, err := time.LoadLocation("Europe/Warsaw")
	assert.NoError(t, err)
	t1InWarsaw := t1.In(warsaw)
	t2InWarsaw := t2.In(warsaw)
	// Winter time - CET
	assert.Equal(t, 16, t1InWarsaw.Hour())
	// Summer time - CEST
	assert.Equal(t, 17, t2InWarsaw.Hour())

}
