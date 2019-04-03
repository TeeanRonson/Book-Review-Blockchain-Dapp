package data

import (
	"fmt"
	"github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p1"
)

type HeartBeatData struct {
	IfNewBlock  bool   `json:"ifNewBlock"`
	Id          int32  `json:"id"`
	BlockJson   string `json:"blockJson"`
	PeerMapJson string `json:"peerMapJson"`
	Addr        string `json:"addr"`
	Hops        int32  `json:"hops"`
}

type HeartBeatDataMod struct {
	IfNewBlock  bool   				`json:"ifNewBlock"`
	Id          int32  				`json:"id"`
	BlockJson   string 				`json:"blockJson"`
	PeerMapJson map[string]int32 	`json:"peerMapJson"`
	Addr        string 				`json:"addr"`
	Hops        int32  				`json:"hops"`
}

//func NewHeartBeatData(ifNewBlock bool, id int32, blockJson string, peerMapJson string, addr string) HeartBeatData {
//
//	return HeartBeatData{ifNewBlock, id, blockJson, peerMapJson, addr, 3}
//}

/**
Randomly decide if you will generate the next block.
If no:
Return a HeartBeatData without new block;
if yes, do:
(1) Randomly create an MPT.
(2) Generate the next block.
(3) Create a HeartBeatData, add that new block, and return.
 */
//func PrepareHeartBeatData(sbc *SyncBlockChain, selfId int32, peerMapJsonString string, addr string) HeartBeatData {
//
//	//create a new block
//	if rand.Intn(100) < 50 {
//		mpt := p1.GetMPTrie()
//		block := sbc.GenBlock(mpt)
//		blockJson, err := block.EncodeToJson()
//		if err != nil {
//			fmt.Println("Error in PrepareHeartBeatData")
//			panic(err)
//		}
//		fmt.Println("We created a block!")
//		return HeartBeatData{true, selfId, blockJson, peerMapJsonString, addr, 3}
//	} else { //don't create a new block
//		fmt.Println("We are not creating a block!")
//		return HeartBeatData{false, selfId, "", peerMapJsonString, addr, 3}
//	}
//}

func PrepareHeartBeatData(sbc *SyncBlockChain, selfId int32, peerMapJsonString map[string]int32, addr string) HeartBeatDataMod {

	//create a new block
	//if rand.Intn(100) < 50 {
		mpt := p1.GetMPTrie()
		block := sbc.GenBlock(mpt)
		blockJson, err := block.EncodeToJson()
		if err != nil {
			fmt.Println("Error in PrepareHeartBeatData")
			panic(err)
		}
		fmt.Println("We created a block!")
		return HeartBeatDataMod{true, selfId, blockJson, peerMapJsonString, addr, 3}
	//} else { //don't create a new block
	//	fmt.Println("We are not creating a block!")
	//	return HeartBeatDataMod{false, selfId, "", peerMapJsonString, addr, 3}
	//}
}