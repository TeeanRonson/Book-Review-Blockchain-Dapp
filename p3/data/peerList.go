package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
)

type PeerList struct {
	selfId int32
	peerMap map[string]int32
	maxLength int32
	mux sync.Mutex
}

type SinglePeer struct {
	peerAdd string
	peerId int32
}

/**
Create New Peer list
 */
func NewPeerList(id int32, maxLength int32) PeerList {
	peerMap := make(map[string]int32)
	return PeerList{id, peerMap, maxLength, sync.Mutex{}}
}

/**
Add a new address and id to the peer list
 */
func(peers *PeerList) Add(addr string, id int32) {
	peers.mux.Lock()
	defer peers.mux.Unlock()
	peers.peerMap[addr] = id
	fmt.Println("After adding:", peers.peerMap)
}

/**
Delete a peer from the PeerMap
 */
func(peers *PeerList) Delete(addr string) {
	peers.mux.Lock()
	defer peers.mux.Unlock()
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
	peers.mux.Lock()
	defer peers.mux.Unlock()

	//if len(peers.peerMap) <= 32 {
	//	return
	//}
	index := 0
	added := false
	peerListId := RebalanceHelper(peers.peerMap)
	tempList := make([]int32, 0)
	newPeerMap := make(map[string]int32, 0)

	//Add self to list
	for i, entry := range peerListId {
		if added == false && peers.selfId < entry.Value {
			tempList = append(tempList, peers.selfId)
			tempList = append(tempList, entry.Value)
			added = true
			index = i
		} else {
			tempList = append(tempList, entry.Value)
		}
	}

	//Get the closest top 16 peers
	var j int32
	j = 0
	topStart := index + 1
	newPeersId := make(map[int32]bool, 0)
	for j < peers.maxLength/2 {
		nextPeer := tempList[topStart % len(tempList)]
		newPeersId[nextPeer] = true
		topStart++
		j++
	}
	//Get the closest lower 16 peers
	var k int32
	k = 0

	botStart := index - 1
	for k < peers.maxLength/2 {
		nextPeer := tempList[botStart]
		newPeersId[nextPeer] = true
		botStart--
		k++
		if botStart < 0 {
			botStart += len(tempList)
		}
	}

	//create the new peer map
	for key, value := range peers.peerMap {
		if newPeersId[value] == true {
			newPeerMap[key] = value
		}
	}
	peers.peerMap = newPeerMap
}

/**
Show() shows all addresses and their corresponding IDs.
For example, it returns "This is PeerMap: \n addr=127.0.0.1, id=1".
 */
func(peers *PeerList) Show() string {
	peers.mux.Lock()
	defer peers.mux.Unlock()
	var result string
	for key, entry := range peers.peerMap {
		result += "Address = " + string(key) + " Id = " + string(entry) + "\n"
	}
	return result
}

/**
Register an Id for self
 */
func(peers *PeerList) Register(id int32) {
	peers.mux.Lock()
	defer peers.mux.Unlock()
	peers.selfId = id
	fmt.Printf("SelfId=%v\n", id)
}

/**
Get a copy of self Peer Map
 */
func(peers *PeerList) Copy() map[string]int32 {
	peers.mux.Lock()
	defer peers.mux.Unlock()
	return peers.peerMap
}

/**
Get self Id
 */
func(peers *PeerList) GetSelfId() int32 {
	peers.mux.Lock()
	defer peers.mux.Unlock()
	return peers.selfId
}

/**
Convert the PeerMap to Json format
 */
func(peers *PeerList) PeerMapToJson() (string, error) {

	peers.mux.Lock()
	defer peers.mux.Unlock()

	result, err := json.MarshalIndent(peers.peerMap, "", "")
	if err != nil {
		fmt.Println("Cannot Marshal Indent jsonList")
		log.Fatal(err)
	}
	return string(result), nil
}

/**
Inject NewPeerMap to existing PeerMap
 */
func(peers *PeerList) InjectPeerMapJson(peerMap map[string]int32, senderAddr string, senderId int32) {

	peers.mux.Lock()
	defer peers.mux.Unlock()
	//newPeersList := make([]SinglePeer, 0)
	peers.peerMap[senderAddr] = senderId

	//add everything except yours
	for addr, id := range peerMap {
		if !reflect.DeepEqual(id, peers.selfId) {
			peers.peerMap[addr] = id
		}
	}
	fmt.Println("My PeerMap should not contain myself:", os.Args[1], peers.peerMap)
}

