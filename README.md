# gtimer - An easy way to test time cost in golang

Example:

```go
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
```