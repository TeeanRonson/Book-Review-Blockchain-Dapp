package data

import (

)

type HeartBeatData struct {
	IfNewBlock  bool   `json:"ifNewBlock"`
	Id          int32  `json:"id"`
	BlockJson   string `json:"blockJson"`
	PeerMapJson string `json:"peerMapJson"`
	Addr        string `json:"addr"`
	Hops        int32  `json:"hops"`
}

func NewHeartBeatData(ifNewBlock bool, id int32, blockJson string, peerMapJson string, addr string) HeartBeatData {

	return HeartBeatData{ifNewBlock, id, blockJson, peerMapJson, addr, 3}
}

func PrepareHeartBeatData(sbc *SyncBlockChain, selfId int32, peerMapBase64 string, addr string) HeartBeatData {


	return HeartBeatData{}
}