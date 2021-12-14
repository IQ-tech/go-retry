package retry

import (
	"math"
	"time"

	"github.com/IQ-tech/go-errors"
)

type Options struct {
	Attempts                  int
	InitialTimeBetweenRetries int
}

func Func(options Options, function func() (interface{}, error)) (out interface{}, err error) {
	for i := 0; i < options.Attempts; i++ {
		if i != 0 {
			aux := math.Pow(float64(options.InitialTimeBetweenRetries), float64(i))
			time.Sleep(time.Duration(aux) * time.Second)
		}
		out, err = function()
		if err == nil {
			return out, nil
		}
	}
	return nil, errors.Wrap(err)
}
