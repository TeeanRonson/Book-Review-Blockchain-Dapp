package data

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
)

type PeerList struct {
	selfId int32
	peerMap map[string]int32
	maxLength int32
	mux sync.Mutex
}

type SinglePeer struct {
	peerId int32
	peerAdd string
}

func NewPeerList(id int32, maxLength int32) PeerList {

	peerMap := make(map[string]int32)
	return PeerList{id, peerMap, maxLength, sync.Mutex{}}

}

/**
Add a new address and id to the peer list
 */
func(peers *PeerList) Add(addr string, id int32) {

	peers.peerMap[addr] = id
	peers.maxLength++
}

/**
Delete a peer from the PeerMap
 */
func(peers *PeerList) Delete(addr string) {

	delete(peers.peerMap, addr)
}

/**
Before sending HeartBeat, Rebalance the PeerList
by choosing 32 closest peers - 16 below and 16 above
1. Sort Map by Id
2. Insert selfId
3. Choose 16 nodes at each side of selfId
 */
func(peers *PeerList) Rebalance() {

	//Get the sorted PairList
	index := 0
	added := false
	peerListId := RebalanceHelper(peers.peerMap)
	tempList := make([]int32, 0)
	newPeerMap := make(map[string]int32, 0)

	//Add self to list
	for i, entry := range peerListId {
		if added == false && peers.selfId < entry.Value {
			tempList = append(tempList, peers.selfId)
			added = true
			index = i
		} else {
			tempList = append(tempList, entry.Value)
		}
	}

	//Get the closest top 16 peers
	j := 0
	topStart := index
	newPeersId := make(map[int32]bool, 0)
	for j < 16 {
		nextPeer := tempList[topStart % len(peers.peerMap)]
		newPeersId[nextPeer] = true
		topStart++
		j++
	}

	//Get the closes lower 16 peers
	k := 0
	botStart := index
	for k < 16 {
		nextPeer := tempList[botStart]
		newPeersId[nextPeer] = true
		botStart--
		k++
		if botStart < 0 {
			botStart += len(peers.peerMap)
		}
	}

	//create the new peer map
	for key, value := range peers.peerMap {
		if newPeersId[value] == true {
			newPeerMap[key] = value
		}
	}
}

/**
Show the peer list is a string
 */
func(peers *PeerList) Show() string {

	var result string
	for _, value := range peers.peerMap {
		result += string(value) + " "
	}
	return result
}

/**
Register?
 */
func(peers *PeerList) Register(id int32) {
	peers.selfId = id
	fmt.Printf("SelfId=%v\n", id)
}

/**

 */
func(peers *PeerList) Copy() map[string]int32 {}

/**

 */
func(peers *PeerList) GetSelfId() int32 {
	return peers.selfId
}

/**
Convert the PeerMap to Json format
 */
func(peers *PeerList) PeerMapToJson() (string, error) {

	peerList := make([]SinglePeer, 0)

	for key, value := range peers.peerMap {
		peerList = append(peerList, SinglePeer{value, key})
	}

	result, err := json.MarshalIndent(peerList, "", "")
	if err != nil {
		fmt.Println("Cannot Marshal Indent jsonList")
		log.Fatal(err)
	}

	return string(result), nil
}

/**
Inject NewPeerMap to existing PeerMap
 */
func(peers *PeerList) InjectPeerMapJson(peerMapJsonStr string, selfAddr string) {





}

func TestPeerListRebalance() {
	//test1
	peers := NewPeerList(5, 4)
	peers.Add("1111", 1)
	peers.Add("4444", 4)
	peers.Add("-1-1", -1)
	peers.Add("0000", 0)
	peers.Add("2121", 21)
	peers.Rebalance()
	expected := NewPeerList(5, 4)
	expected.Add("1111", 1)
	expected.Add("4444", 4)
	expected.Add("2121", 21)
	expected.Add("-1-1", -1)
	fmt.Println(reflect.DeepEqual(peers, expected))

	peers = NewPeerList(5, 2)
	peers.Add("1111", 1)
	peers.Add("4444", 4)
	peers.Add("-1-1", -1)
	peers.Add("0000", 0)
	peers.Add("2121", 21)
	peers.Rebalance()
	expected = NewPeerList(5, 2)
	expected.Add("4444", 4)
	expected.Add("2121", 21)
	fmt.Println(reflect.DeepEqual(peers, expected))

	peers = NewPeerList(5, 4)
	peers.Add("1111", 1)
	peers.Add("7777", 7)
	peers.Add("9999", 9)
	peers.Add("11111111", 11)
	peers.Add("2020", 20)
	peers.Rebalance()
	expected = NewPeerList(5, 4)
	expected.Add("1111", 1)
	expected.Add("7777", 7)
	expected.Add("9999", 9)
	expected.Add("2020", 20)
	fmt.Println(reflect.DeepEqual(peers, expected))
}