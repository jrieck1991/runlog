package app

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func Do() {
	for {
		time.Sleep(1 * time.Second)
		makeLogs()
	}
}

func makeLogs() {
	l := log.WithField("foo", "bar")
	l.Debug("message")
}
