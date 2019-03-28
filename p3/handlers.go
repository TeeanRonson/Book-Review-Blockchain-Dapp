package p3

import (
	"./data"
	"fmt"
	"net/http"
)

var TA_SERVER = "http://localhost:6688"
var REGISTER_SERVER = TA_SERVER + "/peer"
var BC_DOWNLOAD_SERVER = TA_SERVER + "/upload"
var SELF_ADDR = "http://localhost:6686"

var SBC data.SyncBlockChain
var Peers data.PeerList
var ifStarted bool

/**
Create SyncBlockChain and PeerList instances
 */
func init() {
	data.NewBlockChain()
	data.NewPeerList(1, 32)
}

/**
Register ID, download BlockChain, start HeartBeat
You can start the program by calling this route(be careful to start only once),
or start the program during bootstrap.

Get an ID from TA's server, download the BlockChain from your own first node,
use "go StartHeartBeat()" to start HeartBeat loop.
 */
func Start(w http.ResponseWriter, r *http.Request) {


}

/**
Shows the PeerMap and the BlockChain.
GET
Display the PeerList and the BlockChain. Use the helper function
BlockChain.show() in the starter code to display the BlockChain, and add your own function to display the PeerList.
 */
func Show(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s\n%s", Peers.Show(), SBC.Show())
}

// Register to TA's server, and get an ID
func Register() {}

/**
Download the current BlockChain from your own first node(can be hardcoded).
It's ok to use this function only after launching a new node.
You may not need it after node starts heartBeats.
 */
func Download() {}

/**
/upload
Method: GET
Response: The JSON string of the BlockChain.
Description: Return JSON string of the entire blockchain to the downloader.
Upload blockchain to whoever called this method, return jsonStr

Return the BlockChain's JSON. And add the remote peer into the PeerMap.
 */
func Upload(w http.ResponseWriter, r *http.Request) {
	blockChainJson, err := SBC.BlockChainToJson()
	if err != nil {
		//data.PrintError(err, "Upload")
	}
	fmt.Fprint(w, blockChainJson)
}

/**
/block/{height}/{hash}
Method: GET
Response: If you have the block, return the JSON string of the specific block;
if you don't have the block, return HTTP 204: StatusNoContent;
if there's an error, return HTTP 500: InternalServerError.
Description: Return JSON string of a specific block to the downloader.
Upload a block to whoever called this method, return jsonStr

Return the Block's JSON.
 */
func UploadBlock(w http.ResponseWriter, r *http.Request) {


}

/**
/heartbeat/receive
Method: POST
Request: HeartBeatData(see the data structure below)
Description: Receive a heartbeat.
Add the remote address, and the PeerMapJSON into local PeerMap. Then check if the HeartBeatData contains a new block.
If so, do these: (1) check if the parent block exists. If not, call AskForBlock() to download the parent block.
(2) insert the new block from HeartBeatData.
(3) HeartBeatData.hops minus one, and if it's still bigger than 0, call ForwardHeartBeat() to forward this heartBeat to all peers.
 */
func HeartBeatReceive(w http.ResponseWriter, r *http.Request) {




}
/**
Loop through all peers in local PeerMap to download a block. As soon as one peer returns the block, stop the loop.
 */
func AskForBlock(height int32, hash string) {}

/**
Send the HeartBeatData to all peers in local PeerMap.
 */
func ForwardHeartBeat(heartBeatData data.HeartBeatData) {}

/**
Start a while loop. Inside the loop, sleep for randomly 5~10 seconds,
then use PrepareHeartBeatData() to create a HeartBeatData,
and send it to all peers in the local PeerMap.
 */
func StartHeartBeat() {}