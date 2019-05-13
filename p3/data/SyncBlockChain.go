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
Create a new SyncBlockChain
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
Get the latest blocks from this blockChain
 */
func (sbc *SyncBlockChain) SyncGetLatestBlocks() []p2.Block {
	sbc.mux.Lock()
	defer sbc.mux.Unlock()
	return sbc.bc.GetLatestBlocks()
}

/**
Get the parent of the input block
 */
 func (sbc *SyncBlockChain) SyncGetParentBlock(child p2.Block) p2.Block {
	 sbc.mux.Lock()
	 defer sbc.mux.Unlock()
	 return sbc.bc.GetParentBlock(child)
 }

/**
Get the specific block at the input height with the input hash
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
	//inserts each block into a newly created block chain and returns the block chain
	newBlockChain, err := p2.BlockChainDecodeFromJson(blockChainJson)
	fmt.Println("The block chain we downloaded:", newBlockChain)
	if err != nil {
		fmt.Println("UpdateEntireBlockChain error")
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
	parentHash := sbc.bc.GetLatestBlocks()[0].Header.Hash
	newBlock.NewBlock(height, parentHash, mpt)
	return newBlock
}

/**
Show the block chain
 */
func(sbc *SyncBlockChain) Show() string {
	return sbc.bc.Show()
}

/**
Show the Canonical chain
 */
func(sbc *SyncBlockChain) Canonical() string {

	fmt.Println()
	fmt.Println("Canonical!! -------------")
	canonical := ""
	for _, block := range sbc.SyncGetLatestBlocks() {
		dummy := block
		canonical += "\n"
		for dummy.Header.Height >= 1 {
			blockString := dummy.CompressBlock()
			fmt.Println("THE BLOCK HERE IS: ----------------- ", blockString)
			canonical += blockString + "\n"
			parent := sbc.SyncGetParentBlock(dummy)
			fmt.Println("PARENT BLOCK IS: ------------ ", parent.CompressBlock())
			dummy = parent
			fmt.Println("DUMMY BLOCK IS: ------------ ", dummy.CompressBlock())
		}
		fmt.Println("OUT OF LOOP")
	}
	return canonical
}


/**
Return all the book reviews in every block of the block chain as a Json
 */
func (sbc *SyncBlockChain) GetAllReviews() string {


	return "These are all the book reviews"
}