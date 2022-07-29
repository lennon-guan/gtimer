package gtimer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/petermattis/goid"
)

type (
	ResultWriter func(topic string, steps Steps)
	timer        struct {
		gid   int64
		topic string
		steps []StepInfo
	}
)

var timerMap sync.Map

//Start create a timer and begin recording
func Start(topic string) *timer {
	t := &timer{
		gid:   goid.Get(),
		topic: topic,
		steps: []StepInfo{{At: time.Now(), Name: "begin"}},
	}
	timerMap.Store(t.gid, t)
	return t
}

//End end up the recording, and write the result by the default result writer
func (t *timer) End() {
	t.EndWiteWriter(defaultResultWriter)
}

//End end up the recording, and write the result by the given result writer
func (t *timer) EndWiteWriter(w ResultWriter) {
	t.tick("end")
	if w != nil {
		w(t.topic, t.steps)
	}
	timerMap.Delete(t.gid)
}

func (t *timer) tick(step string) {
	t.steps = append(t.steps, StepInfo{At: time.Now(), Name: step})
}

//Tick
func Tick(step string) {
	t, found := timerMap.Load(goid.Get())
	if !found {
		return
	}
	timer, isTimer := t.(*timer)
	if !isTimer {
		return
	}
	timer.tick(step)
}

func defaultWriteResult(topic string, steps Steps) {
	rc := rand.Intn(10000)
	fmt.Printf("Time cost for topic '%s@%04d' total %s: \n", topic, rc, steps.TotalDuration())
	for i := 1; i < len(steps); i++ {
		fmt.Printf("  |- [%s@%04d: \"%s\" ~ \"%s\"]: %s\n",
			topic, rc, steps[i-1].Name, steps[i].Name, steps.DurationBetween(i-1, i))
	}
}

var defaultResultWriter = defaultWriteResult

//SetDefaultWriter set default result writer
func SetDefaultWriter(w ResultWriter) {
	defaultResultWriter = w
}
