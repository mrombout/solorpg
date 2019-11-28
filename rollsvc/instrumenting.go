package rollsvc

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RollCount   metrics.Counter
	RollLatency metrics.Histogram
	Next        RollService
}

func (mw InstrumentingMiddleware) Roll(diceNotation string) (result RollResult, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "roll", "error", fmt.Sprint(err != nil)}
		mw.RollCount.With(lvs...).Add(1)
		mw.RollLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	result, err = mw.Next.Roll(diceNotation)
	return
}
