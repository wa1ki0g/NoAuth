package main

import (
	"flag"
	"fmt"
	"noauth/lib"
	"os"
	"runtime"
	"strings"
)

var (
	u     string
	n     string
	t     int
	h     bool
	a     string
	debug int
)

func init() {
	flag.BoolVar(&h, "h", false, "This help")
	flag.StringVar(&u, "u", "", "A target url(Please add http or https)")
	flag.StringVar(&n, "n", "", "An interface without authentication, such as /login")
	flag.StringVar(&a, "a", "", "An interface that requires authentication, such as /admin/adduser")
	flag.IntVar(&t, "t", runtime.NumCPU(), "Thread Num")
	flag.IntVar(&debug, "debug", 0, "choose start debug, such -debug 1")
	flag.Usage = usage
}

func checkFlags() {
	if u == "" || n == "" || a == "" {
		fmt.Println("Error: Missing parameter. Please use the -h  to view the required parameters.")
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `noauth version: 1.0.0
Usage:  [-unat] [-u url] [-n interface without authentication] [-a interface An interface that requires authentication] [-t thread] [-debug choose start debug] [-h help]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	lib.Logo()
	flag.Parse()

	if h {
		flag.Usage()
		os.Exit(0)
	}
	checkFlags()

	res1 := strings.Contains(u, "http://")
	res2 := strings.Contains(u, "https://")

	if !res1 && !res2 {
		fmt.Println(lib.Red("[-] Please add http or https for url !!!"))
		os.Exit(0)

	}
	lib.GetStart(u, n, a, t, debug)
	lib.PostStart(u, n, a, t, debug)

	//baseURL := "/admin/adduse/123"
	//baseURL2 := "/login/1/2/3/4"
	//
	//poc.Summary(baseURL2, baseURL)

}
