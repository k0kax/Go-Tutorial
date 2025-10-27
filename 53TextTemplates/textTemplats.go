package main

//模板
import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	// .占位符
	t1, err := t1.Parse("Value is {{.}}\n") //解析模板
	if err != nil {
		panic(err)
	}

	//Must() 用于封装可能 panic 的操作，简化错误处理流程
	t1 = template.Must(t1.Parse("Value:{{.}}\n"))

	//动态参数渲染
	t1.Execute(os.Stdout, "some text") //注入模板
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	//工厂函数
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})
	//{{if .}}yes条件为真时执行，{{else}}条件为假时不执行
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range:{{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

}
