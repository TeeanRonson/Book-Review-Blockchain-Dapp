package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "fmt"
    "os"
)

func main() {

    db := make(map[string]int32, 0)

    db["Principles"] = 0
    db["CRA"] = 1

    for title, id := range db {
        fmt.Println(title)
        fmt.Println(id)
    }

}

func keys() {

    RongPrivate, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        fmt.Println(err.Error)
        os.Exit(1)
    }
    RongPublic := &RongPrivate.PublicKey

    JasonPrivate, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        fmt.Println(err.Error)
        os.Exit(1)
    }
    JasonPublic := &JasonPrivate.PublicKey
    fmt.Println("Rong Private Key:", RongPrivate)
    fmt.Println("Rong Public Key:", RongPublic)
    fmt.Println("Jason Private Key:", JasonPrivate)
    fmt.Println("Jason Public Key:", JasonPublic)


    message := []byte("the code must be like a piece of music")
    label := []byte("")
    hash := sha256.New()
    ciphertext, err := rsa.EncryptOAEP(
        hash,
        rand.Reader,
        JasonPublic,
        message,
        label,
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("OAEP encrypted [%s] to\n[%x]\n", string(message), ciphertext)


}
