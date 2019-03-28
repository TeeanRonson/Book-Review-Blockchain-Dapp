package data

import (
	"fmt"
	"github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p1"
	"github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p2"
	"reflect"
	"sync"
)

type SyncBlockChain struct {
	bc p2.BlockChain
	mux sync.Mutex
}

/**
TODO: Could this be an issue?
 */
func NewBlockChain() SyncBlockChain {
	return SyncBlockChain{bc: p2.NewBlockChain(), mux: sync.Mutex{}}
}

/**
Get the blocks at the given height
 */
func(sbc *SyncBlockChain) Get(height int32) []p2.Block {
	sbc.mux.Lock()
	defer sbc.mux.Unlock()
	return sbc.bc.Get(height)
}

/**
Get the block at the input height with the input hash
 */
func(sbc *SyncBlockChain) GetBlock(height int32, hash string) p2.Block {
	sbc.mux.Lock()
	defer sbc.mux.Unlock()
	dummy := p2.Block{}
	for _, block := range sbc.bc.Chain[height] {
		if reflect.DeepEqual(block.Header.Hash, hash) {
			return block
		}
	}
	fmt.Println("Returning the dummy block")
	return dummy
}

/**
Insert a new block into the blockchain
 */
func(sbc *SyncBlockChain) Insert(block p2.Block) {
	sbc.mux.Lock()
	defer sbc.mux.Unlock()
	sbc.bc.Insert(block)
}

/**
Check the parent hash of the block
if parentHash of the block exists, return true
else return false

CheckParentHash() is used to check if the parent hash(or parent block) exists in the current blockchain before you insert a new block sent by others.
For example, if you received a block of height 7, you should check if that blocks parent block(of height 6) exist in your blockchain.
If not, you should ask others to download that parent block of height 6 before inserting the block of height 7.
 */
func(sbc *SyncBlockChain) CheckParentHash(insertBlock p2.Block) bool {
	sbc.mux.Lock()
	defer sbc.mux.Unlock()
	chain := sbc.bc.Get(insertBlock.Header.Height - 1)

	for _, block := range chain {
		if reflect.DeepEqual(block.Header.Hash, insertBlock.Header.ParentHash) {
			return true
		}
	}
	return false
}

/**
Replace existing blockchain with the new BlockChainJson
 */
func(sbc *SyncBlockChain) UpdateEntireBlockChain(blockChainJson string) {
	sbc.mux.Lock()
	defer sbc.mux.Unlock()
	newBlockChain, err := p2.BlockChainDecodeFromJson(blockChainJson)
	if err != nil {
		fmt.Println("Can't Update Entire BlockChain")
		panic(err)
	}
	sbc.bc = newBlockChain
}

/**
Convert the entire blockchain to Json
 */
func(sbc *SyncBlockChain) BlockChainToJson() (string, error) {
	sbc.mux.Lock()
	defer sbc.mux.Unlock()
	result, err := sbc.bc.BlockChainEncodeToJson()
	if err != nil {
		fmt.Println("BlockChainToJsonError in SyncBlockchain")
		return "", err
	}
	return result, nil
}

/**
Generate a new Block
 */
func(sbc *SyncBlockChain) GenBlock(mpt p1.MerklePatriciaTrie) p2.Block {

	sbc.mux.Lock()
	defer sbc.mux.Unlock()
	newBlock := p2.Block{}
	//generate a new block by passing in (height + 1, the parentHash, mpt)
	height := sbc.bc.Length + 1
	parentHash := sbc.bc.Chain[height][0].Header.ParentHash
	newBlock.NewBlock(height, parentHash, mpt)
	return newBlock
}

func(sbc *SyncBlockChain) Show() string {
	return sbc.bc.Show()
}