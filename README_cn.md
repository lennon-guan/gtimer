# gtimer - 一个简易的go程序耗时统计工具

* 特点：基于[goid](github.com/petermattis/goid)库，可以在同一个goroutine上的各个函数内，不传递上下文对象的情况下，对完整过程进行耗时统计
  
## 基本用法：

* 创建统计任务

```go
// 简易方法，以该行代码为统计起点，以该函数退出为统计终点
defer gtimer.Start("job name").End()

// 更为灵活的方法
timer := gtimer.Start("job name")
...
timer.End()
```

* 记录时间点

```go
gtimer.Tick("step name") // 自动获取当前goroutine上已经创建好的timer对象，进行打点记录。如果无已创建的timer，该打点无效
```

## 例子：

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