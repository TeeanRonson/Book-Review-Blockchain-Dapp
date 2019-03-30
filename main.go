package main

import (
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {

	router := p3.NewRouter()
	if len(os.Args) > 1 {
	    fmt.Println("Here with", os.Args[1], os.Args[2])
		log.Fatal(http.ListenAndServe(":" + os.Args[1], router))
	} else {
		log.Fatal(http.ListenAndServe(":6686", router))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func doSomething(s string) {
    fmt.Println("doing something", s)
}

func startPolling1() {
    for {
        time.Sleep(2 * time.Second)
        go doSomething("from polling 1")
    }
}

func startPolling2() {
    for {
        <-time.After(2 * time.Second)
        go doSomething("from polling 2")
    }
}



