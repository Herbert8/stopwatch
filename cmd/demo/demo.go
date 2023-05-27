package main

import (
	"github.com/Herbert8/stopwatch"
	"time"
)

func main() {
	stopwatch.Enabled = true
	sinceStartTester := stopwatch.NewStopwatch("SinceStartTester")
	sinceStartTester.Enabled = false
	time.Sleep(time.Second)
	sinceStartTester.PrintDurationSinceStart()
	time.Sleep(time.Second * 2)
	sinceStartTester.PrintDurationSinceStart()

	checkpointTester := stopwatch.NewStopwatch("CheckpointTester")
	time.Sleep(time.Second)
	checkpointTester.PrintDurationSinceLastCheckpoint("CheckpointA")
	time.Sleep(time.Second * 2)
	checkpointTester.PrintDurationSinceLastCheckpoint("CheckpointB")
}
