# gtimer - An easy way to test time cost in golang

Example:

```go
// gtimer_test.go

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

/* Output:
Time cost for topic 'test1' total 302.772167ms: 
  |- [begin-enter testFunc]: 101.1695ms
  |- [enter testFunc-exit testFunc]: 151.044334ms
  |- [exit testFunc-end]: 50.558333ms
*/
```