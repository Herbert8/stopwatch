package main

import (
	"github.com/Herbert8/stopwatch"
	"time"
)

func main() {
	sinceStartTester := stopwatch.NewStopwatch("SinceStartTester")
	time.Sleep(time.Second)
	sinceStartTester.PrintDurationSinceStart()
	time.Sleep(time.Second * 2)
	sinceStartTester.PrintDurationSinceStart()

	checkpointTester := stopwatch.NewStopwatch("CheckpointTester")
	time.Sleep(time.Second)
	checkpointTester.PrintDurationSinceLastCheckpoint()
	time.Sleep(time.Second * 2)
	checkpointTester.PrintDurationSinceLastCheckpoint()
}
