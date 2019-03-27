package p2

import (
    "bytes"
    "encoding/gob"
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686_blockchain_P1_Go_skeleton/p1"
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