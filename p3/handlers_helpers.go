package p3

import (
    "crypto/rand"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p2"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3/data"
    "golang.org/x/crypto/sha3"
    "reflect"
    "strconv"
)

func convertToInt32(value string) int32 {

    i, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
       fmt.Println("Unable to convert string to int")
       panic(err)
    }
    return int32(i)
}

func int32ToString(n int32) string {
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
func decodeJsonToHbdMod(hbd string) data.HeartBeatDataMod {

    recvHeartBeatData := data.HeartBeatDataMod{}
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

    fmt.Println(value[:4])
    expected := "0000"
    if reflect.DeepEqual(value[:4], expected) {
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
func verifyProofOfWork(newBlock p2.Block) bool {

   result := sha3.Sum256([]byte(newBlock.Header.ParentHash + newBlock.Header.Nonce + newBlock.Mpt.Root))

   return CheckProofOfWork(hex.EncodeToString(result[:]))

}
