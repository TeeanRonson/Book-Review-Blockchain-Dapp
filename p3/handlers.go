package p3

import (
	"encoding/json"
	"fmt"
	"github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p2"
	"github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3/data"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

var TA_SERVER = "http://localhost:6688"
var REGISTER_SERVER = TA_SERVER + "/peer"
var BC_DOWNLOAD_SERVER = TA_SERVER + "/upload"
var LOCAL_HOST = "http://localhost:"
var FIRST_NODE = "http://localhost:6686"
var GET_BC_FIRST_NODE = FIRST_NODE + "/upload"

var SBC data.SyncBlockChain
var Peers data.PeerList
var ifStarted bool

//wait some time before creating blocks
/**
Create SyncBlockChain and PeerList instances
 */
func Init() {

	SBC = data.NewBlockChain()
	Peers = data.NewPeerList(convertToInt32(os.Args[1]), 32)
	ifStarted = false
}

/**
Register ID, download BlockChain, start HeartBeat
1. Get an ID from TA's server
2. Download the BlockChain from your own first node,
3. Use "go StartHeartBeat()" to start HeartBeat loop.
 */
func Start(w http.ResponseWriter, r *http.Request) {

	//Register yourself
	Init()
	if os.Args[2] == "on" {
		block := p2.Block{}
		block.CreateGenesisBlock()
		SBC.Insert(block)
		fmt.Println("Printing the block", block)
	} else {
		Download()
		ifStarted = true
	}
	//Peers.Register(convertToInt32(os.Args[0]))
	go StartHeartBeat()
}

/**
Register to TA's server, and get an ID
  */
func Register() {

	//Get an Id from TA's server
	//fmt.Println("Register Handler")
	//resp, err := http.Get(TA_SERVER)
	//if err != nil {
	//	fmt.Println("Unable to Register")
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//bytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Unable to read register Id")
	//	log.Fatal(err)
	//}
	//
	////Register the Id
	//id := binary.BigEndian.Uint32(bytes)
	//fmt.Println("This should be the id:", id)
	//Peers.Register(1)
}

/**
Download the current BlockChain from your the first node(can be hardcoded).
It's ok to use this function only after launching a new node.
You may not need it after node starts heartBeats.
 */
func Download() {

	//Get the block chain from first node
	resp, err := http.Get(GET_BC_FIRST_NODE)

	if err != nil {
		fmt.Println("Error in Download Handler")
		panic(err)
	}

	//decoder := json.NewDecoder(resp.Body)
	currBlockChain, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Can't read response body")
		log.Fatal(err)
	}
	fmt.Println("The block chain we downloaded:", currBlockChain)
	SBC.UpdateEntireBlockChain(string(currBlockChain))
}

/**
Shows the PeerMap and the BlockChain.
GET
Display the PeerList and the BlockChain. Use the helper function
BlockChain.show() in the starter code to display the BlockChain, and add your own function to display the PeerList.
 */
func Show(w http.ResponseWriter, r *http.Request) {

	_, _ = fmt.Fprintf(w, "%s\n%s", Peers.Show(), SBC.Show())

}

/**
/upload
Method: GET
Response: The JSON string of the BlockChain.
Description: Return JSON string of the entire blockchain to the downloader.
Return the BlockChain's JSON. And add the remote peer into the PeerMap.
 */
func Upload(w http.ResponseWriter, r *http.Request) {
	//senderAdd := r.RemoteAddr
	//senderId := r.URL.Query()["id"][0]

	//Add remote peer into PeerMap
	//TODO: Do we need to do this?
	//Peers.Add(senderAdd, convertToInt32(senderId))

	//Return the BlockChain's JSON
	blockChainJson, err := SBC.BlockChainToJson()
	if err != nil {
		fmt.Println("Upload Endpoint error")
		panic(err)
	}
	//TODO: Do we have to return the id of this sender to the request client?
	fmt.Fprint(w, blockChainJson)
}

/**
/block/{height}/{hash}
Method: GET
Response: If you have the block at that height, return the JSON string of the specific block;
if you don't have the block, return HTTP 204: StatusNoContent;
if there's an error, return HTTP 500: InternalServerError.
Description: Return JSON string of a specific block to the downloader.
 */
func UploadBlock(w http.ResponseWriter, r *http.Request) {

	path := strings.Split(r.URL.Path, "/")
	fmt.Println("Upload block path:", path)
	height := path[1]
	targetHash := path[2]
	blocksAtHeight := SBC.Get(convertToInt32(height))
	var theBlock p2.Block

	if path[0] != "block" || len(path) != 3 || path == nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	for _, entry := range blocksAtHeight {
		if reflect.DeepEqual(entry.Header.Hash, targetHash) {
			theBlock = entry
		}
	}

	if theBlock.Header.Size == 0 {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
	} else {
		theBlockJson, err := theBlock.EncodeToJson()
		if err != nil {
			fmt.Println("UploadBlock error")
			panic(err)
		}
		fmt.Fprint(w, theBlockJson)
	}
}

/**
/heartbeat/receive
Method: POST
Request: HeartBeatData
Description: Receive a heartbeat.
Add the remote address, and the PeerMapJSON into local PeerMap. Then check if the HeartBeatData contains a new block.
If so, do these:
(1) check if the parent block exists. If not, call AskForBlock() to download the parent block.
(2) insert the new block from HeartBeatData.
(3) HeartBeatData.hops minus one, and if it's still bigger than 0, call ForwardHeartBeat() to forward this heartBeat to all peers.
 */
func HeartBeatReceive(w http.ResponseWriter, r *http.Request) {

	//ifStarted wait till we finish downloading BC

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Convert Json into HeartBeatData
	hbd := decodeJsonToHbd(string(b))
	//Add sender PeerMap and sender into PeerMap
	Peers.InjectPeerMapJson(hbd.PeerMapJson, hbd.Addr, hbd.Id)
	//if HeartBeatData has a new block
	if hbd.IfNewBlock {
		//Get the new block
		recvBlock, _ := p2.DecodeFromJson(hbd.BlockJson)
		if SBC.CheckParentHash(recvBlock) {
			SBC.Insert(recvBlock)
		} else {
			//Asks for the parent block and inserts the parent block
			AskForBlock(recvBlock.Header.Height - 1, recvBlock.Header.ParentHash)
			SBC.Insert(recvBlock)
		}
		hbd.Hops = hbd.Hops - 1
		if hbd.Hops > 0 {
			ForwardHeartBeat(hbd)
		}
	}
}

/**
Loop through all peers in local PeerMap to download a block. As soon as one peer returns the block, stop the loop.
add the block to the block chain
 */
func AskForBlock(height int32, hash string) {

	for add, _ := range Peers.Copy() {
		endPoint := "LOCAL_HOST" + add + "/block/" + string(height) + "/"+ hash
		resp, err := http.Get(endPoint)
		if err != nil {
			fmt.Println("AskForBlock Error")
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Unable to read AskForBlock body")
			log.Fatal(err)
		}
		message := string(bytes)

		if message != http.StatusText(http.StatusInternalServerError) && message != http.StatusText(http.StatusNoContent) {
			//Then we shouldve received a blockJson
			block, _ := p2.DecodeFromJson(message)
			SBC.Insert(block)
			break
		}
	}

}

/**
Send the HeartBeatData to all peers in local PeerMap.
 */
func ForwardHeartBeat(heartBeatData data.HeartBeatData) {

	jsonFormatted, err := json.Marshal(heartBeatData)
	if err != nil {
		fmt.Println("Error in ForwardHeartBeatData")
	}
	for addr, _ := range Peers.Copy() {
		recipient := addr + "/heartbeat/receive"
		//send HeartBeatData to all peers in the local PeerMap
		_, err := http.Post(recipient, "application/json; charset=UTF-8", strings.NewReader(string(jsonFormatted)))
		if err != nil {
			fmt.Println("Can't send out HeartBeats")
			panic(err)
		}
	}
}

/**
Start a while loop. Inside the loop, sleep for randomly 5~10 seconds,
then use PrepareHeartBeatData() to create a HeartBeatData,
and send it to all peers in the local PeerMap.
 */
func StartHeartBeat() {

	for {
		time.Sleep(8 * time.Second)
		peerMapAsJson, _ := Peers.PeerMapToJson()
		selfAdd := LOCAL_HOST + os.Args[1]
		hbd := data.PrepareHeartBeatData(&SBC, Peers.GetSelfId(), peerMapAsJson, selfAdd)
		ForwardHeartBeat(hbd)
	}
}
