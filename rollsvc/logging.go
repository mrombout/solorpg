package rollsvc

import (
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   RollService
}

func (mw LoggingMiddleware) Roll(diceNotation string) (result RollResult, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "roll",
			"input", diceNotation,
			"output", result.Result,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	result, err = mw.Next.Roll(diceNotation)
	return
}
