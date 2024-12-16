package utils

import (
	"time"
)

func Repeatable(fun func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fun(); err == nil {
			return nil
		}

		attempts--
		time.Sleep(delay)
	}

	return
}
