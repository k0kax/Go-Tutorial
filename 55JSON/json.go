package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string //导出的字段必须以大写字母开头
}

type response2 struct {
	Page   int      `json:"page"` //用标签声明
	Fruits []string `json:""fruits`
}

func main() {
	fmt.Println("基本数据类型：")
	bolB, _ := json.Marshal(true) //Marshal编排
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	alcB, _ := json.Marshal(slcD)
	fmt.Println(string(alcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	fmt.Println("自定义数据类型：")
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//通用数据结构 json数据解码到Go的值中
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{} //存放解码后的数据

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1,"fruits":["apple","peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)             // 创建指向标准输出的 JSON 编码器
	d := map[string]int{"apple": 5, "lettuce": 7} // 定义待编码的映射数据
	enc.Encode(d)                                 // 流式编码并输出到标准输出

	dec := json.NewDecoder(strings.NewReader(str)) // 从字符串创建读取器（模拟网络/文件输入）
	res1 := response2{}                            // 声明接收解码结果的结构体
	dec.Decode(&res1)                              // 流式解码数据到结构体
	fmt.Println(res1)
}
