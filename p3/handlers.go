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
var LOCAL_HOST = "localhost:"
var FIRST_NODE = "http://localhost:6686"
var GET_BC_FIRST_NODE = FIRST_NODE + "/upload"
var HTTP = "http://"
var HTTPLOCALHOST = HTTP + LOCAL_HOST

var SBC data.SyncBlockChain
var Peers data.PeerList
var ifStarted bool

/**
Create SyncBlockChain and PeerList instances
ifStarted indicates if the block can start sending and receiving blocks
i.e. after it has already downloaded the block chain
 */
func Init() {

	SBC = data.NewBlockChain()
	Peers = data.NewPeerList(convertToInt32(os.Args[1]), 32)
	ifStarted = false
}

/**
Register ID, download BlockChain, start HeartBeat
If node is the primary node, it creates the genesis block
else it will get an ID, then Download() the block chain from the primary node
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
	} else {
		fmt.Println("Fetching the blockchain")
		Download()
		time.Sleep(1 * time.Second)
	}
	ifStarted = true
	go StartHeartBeat()
	_, err :=fmt.Fprint(w, "ifStarted: ", ifStarted)
	if err != nil {
		fmt.Fprint(w, "ifStarted: ", false)
	}

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
Download the current BlockChain from the primary node
 */
func Download() {

	resp, _ := http.Get(GET_BC_FIRST_NODE + "?id=" + os.Args[1])

	currBlockChain, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Can't read response body")
		log.Fatal(err)
	}
	SBC.UpdateEntireBlockChain(string(currBlockChain))
}

/**
Method: GET
Description: Shows the PeerMap and the BlockChain.
 */
func Show(w http.ResponseWriter, r *http.Request) {

	_, _ = fmt.Fprintf(w, "%s\n%s", Peers.Show(), SBC.Show())

}

/**
/upload
Method: GET
Response: The JSON string of the BlockChain.
Description: Return a JSON string representation of the entire blockchain to the downloader.
Return the BlockChain's JSON. And add the remote peer into the PeerMap.
 */
func Upload(w http.ResponseWriter, r *http.Request) {
	if ifStarted {
		senderId := r.URL.Query()["id"][0]
		//Add remote peer into PeerMap
		Peers.Add(senderId, convertToInt32(senderId))

		//Return the BlockChain's JSON
		blockChainJson, err := SBC.BlockChainToJson()
		if err != nil {
			fmt.Println("Upload Endpoint error")
			panic(err)
		}
		fmt.Fprint(w, blockChainJson)
	}
}

/**
/block/{height}/{hash}
Method: GET
Response: JSON string representation of the specific block;
if you don't have the block, return HTTP 204: StatusNoContent;
if there's an error, return HTTP 500: InternalServerError.
Description: Return JSON string representation of a specific block to the requester.
 */
func UploadBlock(w http.ResponseWriter, r *http.Request) {

	if ifStarted {
		var theBlock p2.Block
		path := strings.Split(r.URL.Path, "/")
		height := path[2]
		targetHash := path[3]
		blocksAtHeight := SBC.Get(convertToInt32(height))

		//Check for errors
		if path[1] != "block" || len(path) != 4 || path == nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		//Look for the specific block
		for _, entry := range blocksAtHeight {
			if reflect.DeepEqual(entry.Header.Hash, targetHash) {
				theBlock = entry
			}
		}

		//Check if block is valid
		if theBlock.Header.Size == 0 {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
			return
		} else {
			//Convert the block to Json format
			theBlockJson, err := theBlock.EncodeToJson()
			if err != nil {
				fmt.Println("UploadBlock error")
				panic(err)
			}

			fmt.Println("Uploading this block to you:", theBlockJson)
			fmt.Fprint(w, theBlockJson)
		}
	}
}

/**
/heartbeat/receive
Method: POST
Description: Receive a heartbeat from another Node.
Add the sender address, Id, and PeerMap into local PeerMap. Then check if the HeartBeatData contains a new block.
If true:
(1) check if the parent block exists. If not, call AskForBlock() to download the parent block.
(2) insert the received block into the blockChain
(3) Subtract from HeartBeatData.hops. If hops > 0, call ForwardHeartBeat() to forward this heartBeat to all local peers.
 */
func HeartBeatReceive(w http.ResponseWriter, r *http.Request) {

	//ifStarted wait till we finish downloading BC
	if ifStarted {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		//Convert Json into HeartBeatData
		hbd := decodeJsonToHbdMod(string(body))
		//Add sender PeerMap and sender into PeerMap
		Peers.InjectPeerMapJson(hbd.PeerMapJson, hbd.Addr, hbd.Id)
		//if HeartBeatData has a new block
		if hbd.IfNewBlock {
			//Get the new block
			recvBlock, _ := p2.DecodeFromJson(hbd.BlockJson)
			if !SBC.CheckParentHash(recvBlock) {
				//Asks for the parent block and inserts the parent block
				fmt.Println("ASKING FOR A BLOCK!!!! -----------------------")
				askForBlock(recvBlock.Header.Height - 1, recvBlock.Header.ParentHash)
			}
			SBC.Insert(recvBlock)
			fmt.Println("Inserted successfully:", recvBlock)
			hbd.Hops = hbd.Hops - 1
			hbd.Addr = os.Args[1]
			hbd.Id = convertToInt32(os.Args[1])
			if hbd.Hops > 0 {
				ForwardHeartBeat(hbd)
			}
		}
	}
}

/**
Loop through all peers in local PeerMap to download the requested block.
As soon as one peer returns the block, add the block to the block chain and stop the loop.

 */
func askForBlock(height int32, hash string) {

	if ifStarted {
		for addr, _ := range Peers.Copy() {
			endPoint := HTTPLOCALHOST + addr + "/block/" + int32ToString(height) + "/"+ hash
			fmt.Println("EndPoint in AskforBlock:", endPoint)
			resp, err := http.Get(endPoint)
			if err != nil {
				fmt.Println("AskForBlock Error")
				log.Fatal(err)
			}

			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Unable to read AskForBlock body")
				log.Fatal(err)
			}
			message := string(bytes)
			fmt.Println("the received message should be a json block:", message)

			if resp.StatusCode == http.StatusOK {
				//Then we should've received a blockJson
				block, _ := p2.DecodeFromJson(message)
				SBC.Insert(block)
				break
			}
		}
	}
}

/**
Forward the HeartBeatData to all peers in local PeerMap.
 */
func ForwardHeartBeat(heartBeatData data.HeartBeatDataMod) {

	if !heartBeatData.IfNewBlock {
		return
	}
	jsonFormatted, err := json.MarshalIndent(heartBeatData, "", "")
	if err != nil {
		fmt.Println("Error in ForwardHeartBeatData")
	}
	for addr, _ := range Peers.Copy() {
		recipient := HTTPLOCALHOST + addr + "/heartbeat/receive"
		fmt.Println("The recipient is:", recipient)
		//send HeartBeatData to all peers in the local PeerMap
		_, err := http.Post(recipient, "application/json; charset=UTF-8", strings.NewReader(string(jsonFormatted)))
		if err != nil {
			fmt.Println("Can't send out HeartBeats")
		}
	}
}


/**
Decide once every interval (3secs) if this Node should create a block
If we decide to create a block:
1. Prepare a HeartBeatData
2. Forward onto other peers
 */
func StartHeartBeat() {

	for {
		time.Sleep(3 * time.Second)
		fmt.Println("HeartBeat", os.Args[1])
		selfAdd := os.Args[1]
		hbd := data.PrepareHeartBeatData(&SBC, Peers.GetSelfId(), Peers.Copy(), selfAdd)
		ForwardHeartBeat(hbd)
	}
}
