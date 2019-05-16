package p3

import (
    "crypto/rand"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p1"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p2"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3/data"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p5/nodeData"
    "golang.org/x/crypto/sha3"
    "reflect"
    "strconv"
)

func ConvertToInt32(value string) int32 {

    i, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
       fmt.Println("Unable to convert string to int")
       panic(err)
    }
    return int32(i)
}

func Int32ToString(n int32) string {
    buf := [11]byte{}
    pos := len(buf)
    i := int64(n)
    signed := i < 0
    if signed {
        i = -i
    }
    for {
        pos--
        buf[pos], i = '0'+byte(i%10), i/10
        if i == 0 {
            if signed {
                pos--
                buf[pos] = '-'
            }
            return string(buf[pos:])
        }
    }
}

func RandomHex() (string, error) {
    bytes := make([]byte, 8)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}


/**
Decode the JsonString into HeartBeatData
 */
func DecodeJsonToHbd(hbd string) data.HeartBeatData {

    recvHeartBeatData := data.HeartBeatData{}
    if err := json.Unmarshal([]byte(hbd), &recvHeartBeatData); err != nil {
        fmt.Println("Can't Unmarshal in decodeHBD")
        return recvHeartBeatData
    }
    return recvHeartBeatData
}

/**
Check proof of work
 */
func CheckProofOfWork(value string) bool {

    expected := "00000"
    if reflect.DeepEqual(value[:5], expected) {
        return true
    }
    return false
}

/**
A function to test the ProofOfWork Method
Only difference between this and the original ProofOfWork is the dummy input variables here
 */
func ProofOfWorkTest(parentHash string, mptRootHash string) bool {

    sum := sha3.Sum256([]byte("firstNonce"))
    y := hex.EncodeToString(sum[:])

    for true {

        for !CheckProofOfWork(y) {
            nonce, _ := RandomHex()
            result := sha3.Sum256([]byte(parentHash + nonce + mptRootHash))
            y = hex.EncodeToString(result[:])
        }
        return true
    }
    return false
}

/**
Verify that the SHA3(parentHash + nonce + mptRootHash) value precedes with 4 0's
 */
func VerifyProofOfWork(newBlock p2.Block) bool {

   result := sha3.Sum256([]byte(newBlock.Header.ParentHash + newBlock.Header.Nonce + newBlock.Mpt.Root))

   return CheckProofOfWork(hex.EncodeToString(result[:]))
}

func GetY(parentHash string, nonce string, mptRootHash string) string {
    y := sha3.Sum256([]byte(parentHash + nonce + mptRootHash))
    pow := hex.EncodeToString(y[:])
    return pow
}

/**
Get the height of the blockChain
 */
func GetHeight() int32 {
    return SBC.SyncGetLatestBlocks()[0].Header.Height
}

/**
Get all the new nodeData for this block

Should we add a block size limit?
 */
 func FetchMptData() (p1.MerklePatriciaTrie, float32, bool) {

     mpt := p1.GetMPTrie()
     txFees := float32(0)

     //If TxPool is empty return true
     if TxPool.IsEmpty() {
         return mpt, txFees, true
     }

     for !TxPool.IsEmpty() {
         reviewData, err := TxPool.Poll()
         if err != nil {
             fmt.Println(err)
             break
         }
         jsonReviewData, _ := reviewData.EncodeToJson()
         mpt.Insert(reviewData.Title, jsonReviewData)
         txFees += reviewData.TxFee
     }
     return mpt, txFees, false
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
Verify the signature of the incoming ReviewObject

We verify by decrypting the Object with the sender public key

 */
func VerifySignature(reviewObject nodeData.ReviewObject) bool {

    hash := sha3.Sum256([]byte(reviewObject.Data))
    return reflect.DeepEqual(reviewObject.Signature, hex.EncodeToString(hash[:]))
}
