package ClientNode

import (
    "encoding/hex"
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p5/nodeData"
    "golang.org/x/crypto/sha3"
    "reflect"
    "strconv"
)

func ConvertToFloat32(value string) float32 {

    i, err := strconv.ParseFloat(value, 10)
    if err != nil {
        fmt.Println("Unable to convert string to int")
        panic(err)
    }
    return float32(i)
}

/**
Decode the JsonString into ReviewObject
*/
func DecodeJsonToReviewObject(hbd string) nodeData.ReviewObject {

    reviewObject := nodeData.ReviewObject{}
    if err := json.Unmarshal([]byte(hbd), &reviewObject); err != nil {
        fmt.Println("Can't Unmarshal in decodeReviewObject")
        return reviewObject
    }
    return reviewObject
}

/**
Decode the JsonString into ReviewData
 */
func DecodeJsonToReviewData(jsonString string) (nodeData.ReviewData, error) {

    reviewData := nodeData.ReviewData{}
    if err := json.Unmarshal([]byte(jsonString), &reviewData); err != nil {
        fmt.Println("Can't Unmarshal in decodeReviewData")
        panic(err)
        return reviewData, err
    }
    return reviewData, nil
}


/**
Verify that the SHA3(parentHash + nonce + mptRootHash) value precedes with 4 0's
 */
func VerifySignature(reviewObject nodeData.ReviewObject) bool {

    hash := sha3.Sum256([]byte(reviewObject.Data))
    return reflect.DeepEqual(reviewObject.Signature, hex.EncodeToString(hash[:]))
}