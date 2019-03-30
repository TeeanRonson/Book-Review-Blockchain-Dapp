package p2

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
	"log"
	"reflect"
	"sort"
	"errors"
)

/**
Chain = map: key = block height; value = array of Block objects.
Length = Length equals to the highest block height.
 */
type BlockChain struct {
	Length int32
	Chain map[int32][]Block
}

func NewBlockChain() BlockChain {
	//var genesis Block
	//genesis.CreateGenesisBlock()
	//chain[0] = []Block{genesis}
	chain := make(map[int32][]Block, 0)
	return BlockChain{0, chain}
}
/**
This function takes a height as the argument,
returns the list of blocks stored in that height or None if the height doesn't exist.
 */
func (bc *BlockChain) Get(height int32) []Block {

	currChain := bc.Chain[height]
	if currChain != nil {
		return currChain
	}
	return nil
}

/**
This function takes a block as the argument,
uses its height to find the corresponding list in the Blockchain's Chain map.
If the list already contains that block's hash,
we ignore it because we don't store duplicate blocks; if not, insert the block into the list at that height
 */
func (bc *BlockChain) Insert(block Block) {

	currChain := bc.Get(block.Header.Height)
	fmt.Println("\nInsert New Block")

	if currChain == nil {
		//fmt.Println("No []Block at that height: append to Block height:", block.Header.Height)
		newChain := make([]Block, 0)
		newChain = append(newChain, block)
		//add a new chain at the height
		fmt.Println("block header height", block.Header.Height)
		bc.Chain[block.Header.Height] = newChain

	} else {
		for _, currBlock := range currChain {
			if reflect.DeepEqual(block.Header.Hash, currBlock.Header.Hash) {
				return
			}
		}
		bc.Chain[block.Header.Height] = append(bc.Chain[block.Header.Height], block)
	}
	if bc.Length < block.Header.Height {
		bc.Length = block.Header.Height
	}
}

/**
This function iterates over all the blocks,
generate blocks' JsonString by the function you implemented previously,
and return the list of those JsonStrings
 */
func (bc *BlockChain) BlockChainEncodeToJson() (string, error) {

	jsonList := make([]BlockJson, 0)
	for _, chain := range bc.Chain {
		for _, block := range chain {
			jsonList = append(jsonList, blockToBlockJson(block))
		}
	}

	result, err := json.MarshalIndent(jsonList, "", "")
	if err != nil {
		fmt.Println("Cannot Marshal Indent jsonList")
		log.Fatal(err)
	}
	return string(result), nil
}
/**
This function is called upon a blockchain instance.
It takes a blockchain JSON string as input,
decodes the JSON string back to a list of block JSON strings,
decodes each block JSON string back to a block instance, and inserts every block into the blockchain.
 */
func BlockChainDecodeFromJson(jsonString string) (BlockChain, error) {

	newBlockChain := NewBlockChain()
	blockJsonList := make([]BlockJson, 0)

	if err := json.Unmarshal([]byte(jsonString), &blockJsonList); err != nil {
		panic(err)
		return newBlockChain, errors.New("Blockchain DecodeFromJson error")
	}

	for _, item := range blockJsonList {
		createBlock := blockJsonToBlock(item)
		newBlockChain.Insert(createBlock)
	}
	return newBlockChain, nil
}

func (bc *BlockChain) Show() string {
	rs := ""
	var idList []int
	for id := range bc.Chain {
		idList = append(idList, int(id))
	}
	sort.Ints(idList)
	for _, id := range idList {
		var hashs []string
		for _, block := range bc.Chain[int32(id)] {
			hashs = append(hashs, block.Header.Hash + "<=" + block.Header.ParentHash)
		}
		sort.Strings(hashs)
		rs += fmt.Sprintf("%v: ", id)
		for _, h := range hashs {
			rs += fmt.Sprintf("%s, ", h)
		}
		rs += "\n"
	}
	sum := sha3.Sum256([]byte(rs))
	rs = fmt.Sprintf("This is the BlockChain: %s\n", hex.EncodeToString(sum[:])) + rs
	return rs
}