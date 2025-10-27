package main

//枚举类型
//枚举是编程中用来管理固定值集合的工具，通俗来说就是​​给一组相关的选项起名字​​，让代码更易读、更安全。
//Go 语言中基于 iota 的枚举实现解析
//iota 是 Go 的预定义标识符，用于在 const 块中自动生成递增的整数值
import "fmt"

//简单的服务器状态机，包含状态转换逻辑和友好的字符串输出
type ServerState int

const (
	StateIdle      ServerState = iota //0
	StateConnected                    //1
	StateError                        //2
	StateRetrying                     //3
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "reeor",
	StateRetrying:  "retrying",
}

//打印状态名称
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

}

//状态转换逻辑
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
