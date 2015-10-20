package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	//Trace logger
	Trace *log.Logger
	//Info logger for Info logs
	Info *log.Logger
	//Warning logger for warning logs
	Warning *log.Logger
	//Error logger for error logs
	Error *log.Logger
)

//Init all different logger
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func responseWithIPs(w http.ResponseWriter, r *http.Request) {
	Info.Printf("Got %q Request on %q", r.Method, r.Host)
	io.WriteString(w, "Hello my addresses are the following:\n")

	ifaces, err := net.Interfaces()
	if err != nil {
		Error.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			http.Error(w, err.Error(), 500)
			Error.Println(err.Error())
		}

		for _, addr := range addrs {
			io.WriteString(w, addr.String()+"\n")
		}
	}
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Info.Println("Started")
	http.HandleFunc("/", responseWithIPs)
	http.ListenAndServe(":8000", nil)
}
