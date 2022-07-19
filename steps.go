package gtimer

import "time"

type (
	StepInfo struct {
		Name string
		At   time.Time
	}
	Steps []StepInfo
)

func (s Steps) DurationBetween(i, j int) time.Duration {
	if i < 0 {
		return 0
	}
	if i >= j {
		return 0
	}
	if j >= len(s) {
		return 0
	}
	return s[j].At.Sub(s[i].At)
}

func (s Steps) TotalDuration() time.Duration {
	return s.DurationBetween(0, len(s)-1)
}
