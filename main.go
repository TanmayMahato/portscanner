package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"text/template"
	"time"
)

var c1 chan net.Conn

var wg sync.WaitGroup

type output struct {
	IP   string
	PORT int
	Net  string
	Time string
}

type input struct {
	Startport string
	Endport   string
	Address   string
}

var inputdata input

var outdata []output

func handle(portno int) {
	defer wg.Done()
	port := strconv.Itoa(portno)

	addr := inputdata.Address + ":" + port
	res, err := net.Dial("tcp", addr)
	if err != nil {
	}
	if err == nil {
		c1 <- res

	}
}

func data() {

	for v := range c1 {
		fmt.Println("", v.RemoteAddr().String(), " ", v.RemoteAddr().Network(), " ")
		var o1 output
		t := time.Now()
		ftime := t.Format("2006-01-02 , 15:04:05.000000000")
		o1.Time = ftime
		o1.IP = v.RemoteAddr().String()

		rAddr := v.RemoteAddr()
		o1.PORT = rAddr.(*net.TCPAddr).Port

		o1.Net = v.RemoteAddr().Network()
		outdata = append(outdata, o1)

	}

}

// http request handler
func httphandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		c1 = make(chan net.Conn)
		//Form data
		inputdata.Startport = r.FormValue("startPort")
		inputdata.Endport = r.FormValue("endPort")
		inputdata.Address = r.FormValue("ip")

		//ports str to int
		sport, errsp := strconv.Atoi(inputdata.Startport)
		if errsp != nil {
			fmt.Println(errsp)
			return
		}
		eport, errep := strconv.Atoi(inputdata.Endport)
		if errep != nil {
			fmt.Println(errep)
			return
		}

		//calling the go routines
		for i := sport; i <= eport; i++ {
			wg.Add(1)
			go handle(i)
		}

		//Calling for channel data
		go data()

		//wait till waitGroup is zero
		wg.Wait()

		fmt.Println("done bhaisab")

		temp, err := template.ParseGlob("*.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		errt := temp.ExecuteTemplate(w, "indexdata.html", outdata)
		if errt != nil {
			fmt.Println(errt)
			return
		}

		return
	}

	temp, err := template.ParseGlob("*.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	errt := temp.ExecuteTemplate(w, "index.html", nil)
	if errt != nil {
		fmt.Println(errt)
		return
	}
}

func main() {
	http.HandleFunc("/", httphandle)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
