package main

//timer 定时器
import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C                   // 阻塞等待触发
	fmt.Println("Timer 1 fired") //并未停止，最后直接打印
	/*
		Timer Fired
		​​技术含义​​：定时器到达预设时间后，自动向其通道 C 发送一个值（当前时间），表示计时完成。
		​底层行为​​：定时器触发后，通道 C 仅发送一次数据，之后定时器失效，需调用 Reset() 重新激活。
	*/
	time.Sleep(2 * time.Second)
	timer2 := time.NewTimer(time.Second)
	go func() { // 启动子协程等待触发
		<-timer2.C // 阻塞等待触发
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop() // 主协程立即停止定时器
	if stop2 {             //stop2 必为 true
		fmt.Println("Timer2 2 Stopped")
	}

	time.Sleep(2 * time.Second)
}

/*
执行流程时序图​

主协程              Timer1              Timer2              Goroutine
  │                   │                   │                     │
  ├─创建Timer1(2s)─────>│                   │                     │
  │                   │ 开始计时           │                     │
  │                   │                   │                     │
  ├─阻塞等待Timer1触发 │                   │                     │
  │                   │                   │                     │
  │                   │ 2秒后触发          │                     │
  │<───────────────────┤                   │                     │
  ├─打印"Timer 1 fired"│                   │                     │
  │ Sleep(2秒)                  │                   │                     │
  ├─创建Timer2(1s)─────┼───────────────────>│                     │
  │                   │                   │ 开始计时            │
  │                   │                   │                     │
  ├─启动Goroutine──────┼───────────────────┼─────────────────────>│
  │                   │                   │                     ├─阻塞等待Timer2触发
  │                   │                   │                     │
  ├─调用timer2.Stop()─┼───────────────────>│                     │
  │                   │                   │ 计时停止[3,7](@ref)     │
  │                   │                   │ (若在触发前停止)     │
  ├─Stop返回true───────┼───────────────────┼─────────────────────>│
  │                   │                   │                     │ (Goroutine持续阻塞)
  ├─打印"Timer 2 stopped"                  │                     │
  │                   │                   │                     │
  ├─Sleep(2秒)────────┼───────────────────┼─────────────────────>│
  │                   │                   │                     │
  └─程序退出───────────┼───────────────────┼─────────────────────>│

*/
