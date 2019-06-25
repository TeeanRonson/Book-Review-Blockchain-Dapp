package data

import (
	"fmt"
	"github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p1"
)


type HeartBeatData struct {
	IfNewBlock  bool   				`json:"ifNewBlock"`
	Id          int32  				`json:"id"`
	BlockJson   string 				`json:"blockJson"`
	PeerMapJson map[string]int32 	`json:"peerMapJson"`
	BookDatabase map[string]int32	`json:"bookDatabase"`
	Addr        string 				`json:"addr"`
	Hops        int32  				`json:"hops"`
}

/**
Randomly decide if you will generate the next block.
If no:
Return a HeartBeatData without new block;
if yes, do:
(1) Randomly create an MPT.
(2) Generate the next block.
(3) Create a HeartBeatData, add that new block, and return.
*/
func PrepareHeartBeatData(sbc *SyncBlockChain, selfId int32, peerMapJsonString map[string]int32, bookDatabase map[string]int32, addr string) HeartBeatData {

	fmt.Println("2a. We created HeartBeat!")
	fmt.Println("2b. We created HeartBeat!")
	return HeartBeatData{false, selfId, ""	, peerMapJsonString, bookDatabase, addr, 3}
}

/**
Use the mpt argument to create a new block
Update the block.Header.nonce to the input nonce argument
Create a HeartBeatData and return
 */
func PrepareHeartBeatDataWithBlock(sbc *SyncBlockChain, selfId int32, peerMapJsonString map[string]int32, books map[string]int32, addr string, mpt p1.MerklePatriciaTrie, nonce string) HeartBeatData {

	block := sbc.GenBlock(mpt)
	block.Header.Nonce = nonce
	blockJson, err := block.EncodeToJson()
	if err != nil {
		fmt.Println("Error in PrepareHeartBeatData")
		panic(err)
	}
	fmt.Println("We created a block!")
	return HeartBeatData{true, selfId, blockJson, peerMapJsonString, books,addr, 3}

}
