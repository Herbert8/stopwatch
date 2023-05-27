package stopwatch

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Stopwatch struct {
	name                string
	startTime           time.Time
	lastCheckpointIndex int
	lastCheckpoint      time.Time
}

// 日志
var logger *log.Logger

// 初始化日志
func init() {
	//logger = log.New(os.Stderr, "", log.LstdFlags)
	//logger = log.New(os.Stderr, "", log.Llongfile|log.LstdFlags)
	logger = log.New(os.Stderr, "", log.Lshortfile|log.LstdFlags)
}

// NewStopwatch 通过 名字 新建计时器
func NewStopwatch(name string) *Stopwatch {
	now := time.Now()
	return &Stopwatch{
		name:                name,
		startTime:           now,
		lastCheckpointIndex: 0,
		lastCheckpoint:      now,
	}
}

// Name 返回计时器名称
func (receiver *Stopwatch) Name() string {
	return receiver.name
}

// StartTime 返回计时开始时间
func (receiver *Stopwatch) StartTime() time.Time {
	return receiver.startTime
}

// SinceStart 当前时间距离开始时间的间隔
func (receiver *Stopwatch) SinceStart() time.Duration {
	return time.Since(receiver.startTime)
}

// LastCheckpoint 上一次的检查点
func (receiver *Stopwatch) LastCheckpoint() time.Time {
	return receiver.lastCheckpoint
}

// LastCheckpointIndex 上一次的检查点序号
func (receiver *Stopwatch) LastCheckpointIndex() int {
	return receiver.lastCheckpointIndex
}

// SinceLastCheckpoint 距离上次检查点的时间间隔
func (receiver *Stopwatch) SinceLastCheckpoint() time.Duration {
	now := time.Now()
	duration := now.Sub(receiver.lastCheckpoint)
	var mutex sync.Mutex
	mutex.Lock()
	receiver.lastCheckpointIndex++
	receiver.lastCheckpoint = now
	defer mutex.Unlock()
	return duration
}

// PrintDurationSinceLastCheckpoint 显示距离上次检查点的时间间隔
func (receiver *Stopwatch) PrintDurationSinceLastCheckpoint() {
	duration := receiver.SinceLastCheckpoint()
	s := fmt.Sprintf("[%s] [since checkpoint %d-%d]: %s", receiver.name, receiver.lastCheckpointIndex, receiver.lastCheckpointIndex-1, duration)
	printString(s)
}

// 使用 logger 打印字符串，显示代码行号
func printString(s string) {
	_ = logger.Output(3, s)
}

// PrintDurationSinceStart 显示距离开始时间的时间间隔
func (receiver *Stopwatch) PrintDurationSinceStart() {
	duration := receiver.SinceStart()
	s := fmt.Sprintf("[%s] since start: %s", receiver.name, duration)
	printString(s)
}
