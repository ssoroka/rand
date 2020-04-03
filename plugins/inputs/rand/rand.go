// https://github.com/danielnelson/telegraf-plugins
package rand

import (
	"context"
	"math/rand"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/config"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type RandomNumberGenerator struct {
	ValueName       string          `toml:"value_name"`
	Min             int64           `toml:"min"`
	Max             int64           `toml:"max"`
	SampleFrequency config.Duration `toml:"sample_frequency"`
	ctx             context.Context
	cancel          context.CancelFunc
}

func init() {
	inputs.Add("rand", func() telegraf.Input {
		return &RandomNumberGenerator{
			ValueName:       "value",
			Min:             0,
			Max:             100,
			SampleFrequency: config.Duration(1 * time.Second),
		}
	})
}

func (r *RandomNumberGenerator) Init() error {
	return nil
}

func (r *RandomNumberGenerator) SampleConfig() string {
	return `
  ## Generates random numbers
	[inputs.rand]
	# the name of the measurement to write out to.
	# value_name = "value"
	# min = 0
	# max = 100
	# sample_frequency = "1000ms"
`
}

func (r *RandomNumberGenerator) Description() string {
	return "Generates a random number"
}

func (r *RandomNumberGenerator) Gather(a telegraf.Accumulator) error {
	r.sendMetric(a)
	return nil
}

// // provide the extra functions so we can also run as a service input.
// func (r *RandomNumberGenerator) Start(a telegraf.Accumulator) error {
// 	println("Started as service")
// 	r.ctx, r.cancel = context.WithCancel(context.Background())
// 	go func() {
// 		t := time.NewTicker(r.SampleFrequency)
// 		for {
// 			select {
// 			case <-r.ctx.Done():
// 				t.Stop()
// 				return
// 			case <-t.C:
// 				r.sendMetric(a)
// 			}
// 		}
// 	}()
// 	return nil
// }

func (r *RandomNumberGenerator) Stop() {
	r.cancel()
}

func (r *RandomNumberGenerator) sendMetric(a telegraf.Accumulator) {
	n := rand.Int63n(r.Max-r.Min) + r.Min
	a.AddFields("random",
		map[string]interface{}{
			r.ValueName: n,
		},
		nil,
	)
}
