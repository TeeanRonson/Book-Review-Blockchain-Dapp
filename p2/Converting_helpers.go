package p2

import (
    "bytes"
    "encoding/gob"
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p1"

    "log"
)

/**
Convert Block to BlockJson
 */
func blockToBlockJson(b Block) BlockJson {

    blockJson := BlockJson{
        b.Header.Height,
        b.Header.Timestamp,
        b.Header.Hash,
        b.Header.ParentHash,
        b.Header.Size,
        b.Header.Nonce,
        b.Mpt.Inputs,
    }
    return blockJson
}

/**
Convert BlockJson To BLock
 */
func blockJsonToBlock(blockJson BlockJson) Block {
    var header Header
    newBlock := Block{}
    mpt := NewTrie(blockJson.MPT)
    header.Height = blockJson.Height
    header.Timestamp = blockJson.Timestamp
    header.Hash = blockJson.Hash
    header.ParentHash = blockJson.ParentHash
    header.Size = blockJson.Size

    newBlock.Header = header
    newBlock.Mpt = mpt
    return newBlock
}

/**
Convert Json string to BlockJson
 */
func jsonToBlockJson(jsonString string) (BlockJson, error) {

    blockJson := BlockJson{}
    if err := json.Unmarshal([]byte(jsonString), &blockJson); err != nil {
        fmt.Println("Can't convert Json to BlockJson")
        panic(err)
        return blockJson, err
    }
    return blockJson, nil
}

/**
Convert the MPT struct into bytes
 */
func getBytes(value p1.MerklePatriciaTrie) []byte {

    var network bytes.Buffer        // Stand-in for a network connection
    enc := gob.NewEncoder(&network) // Will write to network.
    //dec := gob.NewDecoder(&network) // Will read from network.
    // Encode (send) the value.
    err := enc.Encode(value)
    if err != nil {
        fmt.Println("error")
        log.Fatal("encode error:", err)
    }
    // HERE ARE YOUR BYTES!!!!
    return network.Bytes()
}

/**
Convert map[string]string to MPT
 */
func NewTrie(values map[string]string) p1.MerklePatriciaTrie {

    mpt := p1.GetMPTrie()
    for key, value := range values {
        mpt.Insert(key, value)
    }
    return mpt
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

func int64ToString(n int64) string {
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