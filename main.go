package main

import (
	"github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3/data"
)
func main() {

	//router := p3.NewRouter()
	//if len(os.Args) > 1 {
	//	log.Fatal(http.ListenAndServe(":" + os.Args[1], router))
	//} else {
	//	log.Fatal(http.ListenAndServe(":6686", router))
	//}


	data.RebalanceHelper()

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, Sir Liew %q", html.EscapeString(r.URL.Path))
	//})
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", Index)
	//log.Fatal(http.ListenAndServe(":8080", router))


}

//func Index(w http.ResponseWriter, r *http.Request) {
//
//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//}
