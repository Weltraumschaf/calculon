package config

import "github.com/urfave/cli/v2"

type Options interface {
	InitOptions(c *cli.Context)
	IsColor() bool
}

type options struct {
	color bool
}

func (o *options) InitOptions(c *cli.Context) {
	o.color = c.Bool("color")
	o.color = !c.Bool("nocolor")
}

func (o *options) IsColor() bool {
	return o.color
}

var applicationOptions Options = &options{}

func ApplicationOptions() Options {
	return applicationOptions
}
