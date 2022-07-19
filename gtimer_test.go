package gtimer

import (
	"testing"
	"time"
)

func testFunc() {
	Tick("enter testFunc")
	time.Sleep(150 * time.Millisecond)
	Tick("exit testFunc")
}
func TestTimer(t *testing.T) {
	defer Start("test1").End()
	time.Sleep(100 * time.Millisecond)
	testFunc()
	time.Sleep(50 * time.Millisecond)
}
