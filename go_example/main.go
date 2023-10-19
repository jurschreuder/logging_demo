package main

import (
	"os"
	"time"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

func main() {

	// construct a ElasticSearch compatible logger
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	sugar := logger.Sugar()

	counter := 0
	for true {
		// info
		sugar.Infof("Logging a counter: %d", counter)
		sugar.Infow("Logging something more advanced",
			"counter", counter,
		)
		// debug
		sugar.Debugf("Debug a counter: %d", counter)
		sugar.Debugw("Debug something more advanced",
			"counter", counter,
		)
		// warning
		sugar.Warnf("Warning on a counter: %d", counter)
		sugar.Warnw("Warning on something more advanced",
			"counter", counter,
		)
		// error
		sugar.Errorf("Error on a counter: %d", counter)
		sugar.Errorw("Error on something more advanced",
			"counter", counter,
		)

		// log an example every 10 seconds
		time.Sleep(5 * time.Second)
		counter++
	}

}
