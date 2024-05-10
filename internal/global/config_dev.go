//go:build dev

package global

import (
	"github.com/Mmx233/config"
)

func initConfig() {
	// load config from yaml
	c := config.NewConfig(&config.Options{
		Config:    &Config,
		Default:   &Config,
		Overwrite: true,
	})
	if err := c.Load(); err != nil {
		panic(err)
	}
}
