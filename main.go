package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "os"
)

func main() {

	//router := p3.NewRouter()
	//if len(os.Args) > 1 {
	//fmt.Println("Here with", os.Args[1], os.Args[2])
	//	log.Fatal(http.ListenAndServe(":" + os.Args[1], router))
	//} else {
	//	log.Fatal(http.ListenAndServe(":6686", router))
	//}

	i, _ := RandomHex()
	fmt.Println(i)

	j, _ := RandomHex()
	fmt.Println(j)


}

func RandomHex() (string, error) {

    bytes := make([]byte, 8)
    special := []byte(os.Args[1])
    //bytes = append(bytes, special)

    fmt.Println(bytes)
    fmt.Println(special)


    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}



