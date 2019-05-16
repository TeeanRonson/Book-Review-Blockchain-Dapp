package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "fmt"
    "os"
    "reflect"
)
func main() {


    mariaPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        fmt.Println(err.Error)
        os.Exit(1)
    }
    mariaPublicKey := &mariaPrivateKey.PublicKey
    raulPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        fmt.Println(err.Error)
        os.Exit(1)
    }
    raulPublicKey := &raulPrivateKey.PublicKey

    fmt.Println("Private Key : ", mariaPrivateKey)
    fmt.Println(reflect.TypeOf(mariaPrivateKey))
    fmt.Println("Public key ", mariaPublicKey)
    fmt.Println(reflect.TypeOf(mariaPublicKey))
    fmt.Println("Private Key : ", raulPrivateKey)
    fmt.Println("Public key ", raulPublicKey)


    message := []byte("the code must be like a piece of music")
    label := []byte("")
    hash := sha256.New()
    cipherText, err := rsa.EncryptOAEP(
        hash,
        rand.Reader,
        raulPublicKey,
        message,
        label,
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(message), cipherText)
}