package console

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs"
)

func init() {
	outputs.RegisterOutputPlugin("console", plugin{})
}

type plugin struct{}

func (p plugin) NewOutput(
	config *outputs.MothershipConfig,
	topologyExpire int,
) (outputs.Outputer, error) {
	pretty := config.Pretty != nil && *config.Pretty
	c := newConsole(pretty)

	// check stdout actually being available
	if _, err := c.out.Stat(); err != nil {
		return nil, fmt.Errorf("console output initialization failed with: %v", err)
	}

	return c, nil
}

type console struct {
	pretty bool
	out    *os.File
}

func newConsole(pretty bool) *console {
	return &console{pretty: pretty, out: os.Stdout}
}

func (c *console) PublishEvent(
	s outputs.Signaler,
	opts outputs.Options,
	event common.MapStr,
) error {
	var jsonEvent []byte
	var err error

	if c.pretty {
		jsonEvent, err = json.MarshalIndent(event, "", "  ")
	} else {
		jsonEvent, err = json.Marshal(event)
	}
	if err != nil {
		logp.Err("Fail to convert the event to JSON: %s", err)
		outputs.SignalCompleted(s)
		return err
	}

	if err = c.writeBuffer(jsonEvent); err != nil {
		goto fail
	}
	if err = c.writeBuffer([]byte{'\n'}); err != nil {
		goto fail
	}

	outputs.SignalCompleted(s)
	return nil
fail:
	if opts.Guaranteed {
		logp.Critical("Unable to publish events to console: %v", err)
	}
	outputs.SignalFailed(s, err)
	return err
}

func (c *console) writeBuffer(buf []byte) error {
	written := 0
	for written < len(buf) {
		n, err := c.out.Write(buf[written:])
		if err != nil {
			return err
		}

		written += n
	}
	return nil
}
