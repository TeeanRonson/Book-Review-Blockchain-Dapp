package p2

import (
    "encoding/hex"
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p1"
    "golang.org/x/crypto/sha3"
    "time"
)

type Block struct {
    Header Header
    Mpt p1.MerklePatriciaTrie
}

type Header struct {
    Height int32                `json:"height"`
    Timestamp int64             `json:"timeStamp"`
    Hash string                 `json:"hash"`
    ParentHash string           `json:"parentHash"`
    Size int32                  `json:"size"`
    Nonce string                `json:"nonce"`
}

type BlockJson struct {
    Height     int32             `json:"height"`
    Timestamp  int64             `json:"timeStamp"`
    Hash       string            `json:"hash"`
    ParentHash string            `json:"parentHash"`
    Size       int32             `json:"size"`
    Nonce      string            `json:"nonce"`
    MPT        map[string]string `json:"mpt"`
}

/**
This function takes arguments(such as height, parentHash, and value of MPT type) and forms a block.
This is a method of the block struct.
 */
func (b *Block) NewBlock(height int32, parentHash string, value p1.MerklePatriciaTrie) {

    var header Header
    mptAsBytes := getBytes(value)

    header.Height = height
    header.Timestamp = int64(time.Now().Unix())
    header.ParentHash = parentHash
    header.Size = int32(len(mptAsBytes))
    header.Nonce = ""
    hashString := string(header.Height) + string(header.Timestamp) + header.ParentHash + value.Root + string(header.Size)
    sum := sha3.Sum256([]byte(hashString))
    header.Hash = hex.EncodeToString(sum[:])

    b.Header = header
    b.Mpt = value
}

/**
Method creates the Genesis block
 */
func (b *Block) CreateGenesisBlock() {

    header := Header{0, int64(time.Now().Unix()), "GenesisBlock", "", 0, ""}
    b.Mpt = p1.GetMPTrie()
    b.Header = header
}

/**
Decode the JsonString into a Block struct object
Note that you have to reconstruct an MPT from the JSON string, and use that MPT as the block's value.
 */
func DecodeFromJson(jsonString string) (Block, error) {

  var header Header
  newBlock := Block{}
  blockJson, err := jsonToBlockJson(jsonString)
  if err != nil {
      return newBlock, err
  }

  mpt := NewTrie(blockJson.MPT)
  header.Height = blockJson.Height
  header.Timestamp = blockJson.Timestamp
  header.Hash = blockJson.Hash
  header.ParentHash = blockJson.ParentHash
  header.Size = blockJson.Size
  header.Nonce = blockJson.Nonce

  newBlock.Header = header
  newBlock.Mpt = mpt
  return newBlock, nil
}

/**
Encode Block struct into to Json string
 */
func (b *Block) EncodeToJson() (string, error) {

    toJson := BlockJson{
        b.Header.Height,
        b.Header.Timestamp,
        b.Header.Hash,
        b.Header.ParentHash,
        b.Header.Size,
        b.Header.Nonce,
        b.Mpt.Inputs,
    }

    jsonFormatted, err := json.Marshal(toJson)
    if err != nil {
        fmt.Println("Error in EncodeToJson")
        return string(jsonFormatted), err
    }
    return string(jsonFormatted), nil
}

/**
Compress the block to show subset of information
 */
func (b *Block) CompressBlock() string {
    return "height=" + int32ToString(b.Header.Height) + ", timestamp=" + int64ToString(b.Header.Timestamp) +  ", hash=" + b.Header.Hash + ", parentHash=" + b.Header.ParentHash + ", size=" + int32ToString(b.Header.Size)
}














