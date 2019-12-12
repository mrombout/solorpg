package rollsvc

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

// InstrumentingMiddleware counts the number of calls and latency to a RollService.
type InstrumentingMiddleware struct {
	RollCount   metrics.Counter
	RollLatency metrics.Histogram
	Next        RollService
}

// Roll counts the number of calls and latency for reach call.
func (mw InstrumentingMiddleware) Roll(diceNotation string, seed int64) (result RollResult, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "roll", "error", fmt.Sprint(err != nil)}
		mw.RollCount.With(lvs...).Add(1)
		mw.RollLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	result, err = mw.Next.Roll(diceNotation, seed)
	return
}
