package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) //协议名
	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	P, _ := u.User.Password()
	fmt.Println(P)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host) //接着拆
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)     //路径
	fmt.Println(u.Fragment) //锚点，对应 URL 的 # 后内容，常用于页面内定位

	fmt.Println(u.RawQuery)            //查询语句
	m, _ := url.ParseQuery(u.RawQuery) //解析成map
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
