package main

import (
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3"
    "log"
    "net/http"
    "os"
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