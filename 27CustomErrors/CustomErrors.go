package main

//自定义错误
import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int    //触发错误的参数值
	message string //错误描述
}

// 实现 error 接口的 Error() 方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)

	var ae *argError
	if errors.As(err, &ae) { //erros.is的升级版，检测和匹配特殊的error
		fmt.Println(ae.arg)
		fmt.Println(ae.message)

	} else {
		fmt.Println("err doesn't match argError")
	}
}
