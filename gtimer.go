package gtimer

import (
	"fmt"
	"sync"
	"time"

	"github.com/petermattis/goid"
)

type (
	StepInfo struct {
		name string
		at   time.Time
	}
	ResultWriter func(topic string, steps []StepInfo)
	timer        struct {
		gid   int64
		topic string
		steps []StepInfo
	}
)

var timerMap sync.Map

func Start(topic string) *timer {
	t := &timer{
		gid:   goid.Get(),
		topic: topic,
		steps: []StepInfo{{at: time.Now(), name: "begin"}},
	}
	timerMap.Store(t.gid, t)
	return t
}

func (t *timer) End() {
	t.EndWiteWriter(defaultResultWriter)
}

func (t *timer) EndWiteWriter(w ResultWriter) {
	t.tick("end")
	if w != nil {
		w(t.topic, t.steps)
	}
	timerMap.Delete(t.gid)
}

func (t *timer) tick(step string) {
	t.steps = append(t.steps, StepInfo{at: time.Now(), name: step})
}

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

func defaultWriteResult(topic string, steps []StepInfo) {
	fmt.Printf("%s耗时情况：\n", topic)
	for i := 1; i < len(steps); i++ {
		fmt.Printf("\t[%s-%s]耗时:%s\n", steps[i-1].name, steps[i].name, steps[i].at.Sub(steps[i-1].at))
	}
}

var defaultResultWriter = defaultWriteResult
