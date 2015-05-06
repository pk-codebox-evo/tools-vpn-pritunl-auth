// Commands available in cli.
package cmd

import (
	"github.com/Sirupsen/logrus"
	"math/rand"
	"os"
	"strconv"
)

type Options struct {
	Debug        bool
	Port         int
	GoogleId     string
	GoogleSecret string
}

func getOpts() (opts *Options) {
	opts = &Options{}

	debugStr := os.Getenv("DEBUG")
	if debugStr == "" {
		opts.Debug = true
	} else {
		opts.Debug, _ = strconv.ParseBool(debugStr)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("cmd: Failed to parse port")
		panic(err)
	}
	opts.Port = port

	opts.GoogleId = os.Getenv("GOOGLE_ID")
	opts.GoogleSecret = os.Getenv("GOOGLE_SECRET")

	return
}