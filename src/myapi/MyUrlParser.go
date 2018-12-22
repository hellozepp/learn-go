package myapi

import "fmt"
import "net"
import "net/url"

func MyUrlParser() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println("-----------1-----------------")
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	fmt.Println("-----------4-----------------")
	fmt.Println(u.Host)
	fmt.Println("---------5-------------------")
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println("--------------6--------------")

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println("------------7----------------")
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
