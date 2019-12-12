package rollsvc

import (
	"time"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware logs all calls to a RollService.
type LoggingMiddleware struct {
	Logger log.Logger
	Next   RollService
}

// Roll logs the call to the Roll method on a RollService.
func (mw LoggingMiddleware) Roll(diceNotation string, seed int64) (result RollResult, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "roll",
			"input", diceNotation,
			"output", result.Result,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	result, err = mw.Next.Roll(diceNotation, seed)
	return
}
