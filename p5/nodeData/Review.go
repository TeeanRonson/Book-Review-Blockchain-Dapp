package nodeData

import (
    "encoding/hex"
    "encoding/json"
    "fmt"
    "golang.org/x/crypto/sha3"
)

type ReviewData struct {
    Title string        `json:"title"`
    ReviewText string   `json:"reviewText"`
    ReviewRating int32  `json:"reviewRating"`
    TxFee float32       `json:"txFee"`
    PublicKey string    `json:"publicKey"`
    Stamp string        `json:"stamp"`
    BookId int32        `json:"bookId"`
}

type ReviewObject struct {
    Data string         `json:"data"`
    Signature string    `json:"signature"`
}

/**
Prepare a new ReviewObject
Review Object consist of the ReviewData and a signature of the user
 */
func PrepareReviewObject(title string, reviewText string, reviewRating int32, txFee float32, publicKey string, priKey string, bookId int32) ReviewObject {

    //Create a stamp
    senderStamp := sha3.Sum256([]byte(priKey))
    stamp := hex.EncodeToString(senderStamp[:])

    //Create the ReviewData
    data := ReviewData{title, reviewText, reviewRating, txFee, publicKey, stamp,bookId}
    jsonData, _ := data.EncodeToJson()

    fmt.Println("Showing the Json Data:" + jsonData)

    //create a signature
    hash := sha3.Sum256([]byte(jsonData))
    signature := hex.EncodeToString(hash[:])

    return ReviewObject{jsonData, signature}

}

/**
Encode Block struct into to Json string
 */
func (data *ReviewData) EncodeToJson() (string, error) {

    jsonFormatted, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error in EncodeToJson")
        return string(jsonFormatted), err
    }
    return string(jsonFormatted), nil
}



