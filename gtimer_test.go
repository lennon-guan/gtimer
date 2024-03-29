package gtimer_test

import (
	"testing"
	"time"

	"github.com/lennon-guan/gtimer"
)

func testFunc() {
	gtimer.Tick("enter testFunc")
	time.Sleep(150 * time.Millisecond)
	gtimer.Tick("exit testFunc")
}

func TestTimer(t *testing.T) {
	defer gtimer.Start("test1").End()
	time.Sleep(100 * time.Millisecond)
	testFunc()
	time.Sleep(50 * time.Millisecond)
}
