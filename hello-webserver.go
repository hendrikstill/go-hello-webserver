package main

import (
	"fmt"
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
	// ifaceString
	output *string
)

func fetchAllIfaces() (*string, error) {
	result := "Hello Knative!!!\n"

	ifaces, err := net.Interfaces()
	if err != nil {
		return &result, err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return &result, err
		}

		for _, addr := range addrs {
			result = fmt.Sprint(result, addr.String()+"\n")
		}
	}

	return &result, nil
}

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
	io.WriteString(w, "Hello from knative! My addresses are the following:\n")
	io.WriteString(w, *output)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Alive")
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	var err error
	output, err = fetchAllIfaces()
	if err != nil {
		log.Fatal(err)
	}

	Info.Println("Started")
	http.HandleFunc("/", responseWithIPs)
	http.HandleFunc("/health", healthCheck)
	http.ListenAndServe(":8000", nil)
}
