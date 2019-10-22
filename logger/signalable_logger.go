package logger

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/bosh-utils/logger"
)

func NewSignalableLogger(log logger.Logger, sigCh chan os.Signal) (logger.Logger, chan bool) {
	doneChannel := make(chan bool, 1)

	go func() {
		for {
			<-sigCh
			fmt.Println("Received SIGHUP - toggling debug output")
			log.ToggleForcedDebug()
			doneChannel <- true
		}
	}()

	return log, doneChannel
}
