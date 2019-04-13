package p3

import (
	"encoding/json"
	"fmt"
	"github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p1"
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
Register ID, download BlockChain, start HeartBeat
If node is the primary node, it creates the genesis block
else it will get an ID, then Download() the block chain from the primary node
1. Get an ID from TA's server
2. Download the BlockChain from your own first node,
3. Use "go StartHeartBeat()" to start HeartBeat loop.
 */
func Start(w http.ResponseWriter, r *http.Request) {

	Init()
	if os.Args[2] == "on" {
		block := p2.Block{}
		block.CreateGenesisBlock()
		SBC.Insert(block)
	} else {
		fmt.Println("Fetching the blockchain")
		Download()
	}
	ifStarted = true
	go StartHeartBeat()
	go StartTryingNonces()
	_, err := fmt.Fprint(w, "ifStarted: ", ifStarted)
	if err != nil {
		fmt.Println("Could not start node")
		panic(err)
	}
}

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
Method: GET
Description: Shows the PeerMap and the BlockChain.
 */
func Show(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s\n%s", Peers.Show(), SBC.Show())
}

/**
This function prints the current canonical chain, and chains of all forks if there are forks.
Note that all forks should end at the same height (otherwise there wouldn't be a fork).
 */
func Canonical(w http.ResponseWriter, r *http.Request) {

	_, _ = fmt.Fprintf(w, "%s", SBC.Canonical())
}

/**
Download the current BlockChain from the primary(leader) node
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

		//Check for errors
		if path[1] != "block" || len(path) != 4 || path == nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		theBlock = SBC.GetBlock(convertToInt32(height), targetHash)

		//Check if block is valid
		if theBlock.Header.Size == 0 {
			fmt.Println("Empty Block ")
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
Decide once every interval (3secs) if this Node should create a block
If we decide to create a block:
1. Prepare a HeartBeatData
2. Forward onto other peers
 */
func StartHeartBeat() {

	for {
		time.Sleep(3 * time.Second)
		fmt.Println("1. HeartBeat of", os.Args[1])
		selfAddr := os.Args[1]
		hbd := data.PrepareHeartBeatData(&SBC, Peers.GetSelfId(), Peers.Copy(), selfAddr)
		ForwardHeartBeat(hbd)
	}
}

/**
Forward the HeartBeatData to all peers in local PeerMap.
 */
func ForwardHeartBeat(heartBeatData data.HeartBeatData) {

	fmt.Println("-- ForwardingHeartBeatData")
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
/heartbeat/receive
Method: POST
Description: Receive a heartbeat from another Node.
Add the sender address, Id, and PeerMap into local PeerMap. Then check if the HeartBeatData contains a new block.
If true:
(1) Verify that the sender has completed proof of work
(2) check if the parent block exists. If not, call AskForBlock() to download the parent block.
(2) insert the received block into the blockChain
(3) Subtract from HeartBeatData.hops. If hops > 0, call ForwardHeartBeat() to forward this heartBeat to all local peers.
 */
func HeartBeatReceive(w http.ResponseWriter, r *http.Request) {

	if ifStarted {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		//Convert Json into HeartBeatData
		hbd := decodeJsonToHbd(string(body))
		Peers.InjectPeerMapJson(hbd.PeerMapJson, hbd.Addr, hbd.Id)

		//TODO: Remove
		if hbd.IfNewBlock {
			fmt.Println("HeartBeatData has a Block")
			fmt.Println(hbd)
		} else {
			fmt.Println("Just a HeartBeat, no Block")
		}
		//if HeartBeatData has a new block
		if hbd.IfNewBlock {
			//Get the new block
			recvBlock, _ := p2.DecodeFromJson(hbd.BlockJson)

			if verifyProofOfWork(recvBlock) {
				fmt.Println("Verified Proof Of Work of Received block")
				currBlock := recvBlock
				//keep asking for the parent block until we have them all
				for !SBC.CheckParentHash(currBlock) {
					fmt.Println("ASKING FOR A PARENT BLOCK!!!! -----------------------")
					parent := askForBlock(currBlock.Header.Height - 1, currBlock.Header.ParentHash)
					SBC.Insert(parent)
					//If we hit the genesis block, break out of the loop
					if reflect.DeepEqual(parent.Header.Hash, "") {
						break
					}
					currBlock = parent
				}
				//Insert the latest block received
				SBC.Insert(recvBlock)
				fmt.Println("Latest block inserted successfully:", recvBlock)
				hbd.Hops = hbd.Hops - 1
				hbd.Addr = os.Args[1]
				hbd.Id = convertToInt32(os.Args[1])
				if hbd.Hops > 0 {
					ForwardHeartBeat(hbd)
				}
			}
		}
	}
}

/**
Loop through all peers in local PeerMap requesting the block with the input hash
As soon as one peer returns the block, return the block
 */
func askForBlock(height int32, hash string) p2.Block {
	var parent p2.Block

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
			//fmt.Println("the received message should be a json block:", message)
			if resp.StatusCode == http.StatusOK {
				//Then we should've received a blockJson
				block, _ := p2.DecodeFromJson(message)
				parent = block
				//SBC.Insert(block)
				return parent
			}
		}
	}
	return parent
}


/**
TODO: Multi-threading, Canonical print, Initalize rand with something unique?
This function starts a new thread that tries different nonces to generate new blocks.
Nonce is a string of 16 hexes such as "1f7b169c846f218a".
Initialize the rand when you start a new node with something unique about each node,
such as the current time or the port number. Here's the workflow of generating blocks:
(1) Start a while loop.
(2) Get the latest block or one of the latest blocks to use as a parent block.
(3) Create an MPT.
(4) Randomly generate the first nonce, verify it with simple PoW algorithm to see if SHA3(parentHash + nonce + mptRootHash)
starts with 10 0's (or the number you modified into).
Since we use one laptop to try different nonces, six to seven 0's could be enough.
If the nonce failed the verification, increment it by 1 and try the next nonce.
(6) If a nonce is found and the next block is generated, forward that block to all peers with a HeartBeatData;
(7) If someone else found a nonce first, and you received the new block through your function ReceiveHeartBeat(), stop trying nonce on the current block, continue to the while loop by jumping to the step(2).
*/
func StartTryingNonces() {

	var checkHeight int32
	selfAddr := os.Args[1]

	for true {
		fmt.Println("TryingNoncesForNewBlock")
		parentHash := SBC.SyncGetLatestBlocks()[0].Header.Hash
		currentHeight := GetHeight()
		checkHeight = currentHeight
		nonce, _ := RandomHex()
		//Create a new MPT
		mpt := FetchMptData()
		y := GetY(parentHash, nonce, mpt.Root)

		for !CheckProofOfWork(y) && checkHeight == currentHeight {
			nonce, _ = RandomHex()
			y = GetY(parentHash, nonce, mpt.Root)
			checkHeight = GetHeight()
		}

		//TODO: ERASE
		if checkHeight != currentHeight {
			fmt.Println("\nSomeone else found the block!!")
			fmt.Println("Current Height:", currentHeight)
			fmt.Println("New Height:", checkHeight)
			fmt.Println()
		}
		//If we solved PoW and no one else has added a new block
		//return the nonce
		if CheckProofOfWork(y) && checkHeight == currentHeight {
			fmt.Println("\nWe solved the block!")
			fmt.Println("Found:", y)
			fmt.Println("Nonce:", nonce)
			ReadyData(nonce, selfAddr, mpt)
		}
	}
}

/**
Create a block with new nonce and new MPT
Form a new HeartBeatData
 */
func ReadyData(nonce string, selfAddr string, mpt p1.MerklePatriciaTrie) {

	if reflect.DeepEqual(nonce, "") {
		fmt.Println("NO NONCE!!!!! ")
	}
	fmt.Println("We found the nonce! Lets prepare the HeartBeatData: Nonce = ", nonce)
	hbd := data.PrepareHeartBeatDataWithBlock(&SBC, Peers.GetSelfId(), Peers.Copy(), selfAddr, mpt, nonce)
	ForwardHeartBeat(hbd)
}

