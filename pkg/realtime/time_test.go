package realtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	assert.NotEqual(t, Now(), time.Now())
}

func TestMockTestsTime(t *testing.T) {
	currentTime := time.Now()
	MockTestsTime(currentTime)

	assert.Equal(t, Now(), currentTime)
}

func MockTestsTime(customtime time.Time) {
	Now = func() time.Time { return customtime }
}

func ResetTestsTime() {
	Now = defaultNowFunction
}
