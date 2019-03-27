package p3

import (
	"../p2"
	"./data"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var TA_SERVER = "http://localhost:6688"
var REGISTER_SERVER = TA_SERVER + "/peer"
var BC_DOWNLOAD_SERVER = TA_SERVER + "/upload"
var SELF_ADDR = "http://localhost:6686"

var SBC data.SyncBlockChain
var Peers data.PeerList
var ifStarted bool

func init() {
	// This function will be executed before everything else.
	// Do some initialization here.
}

/**
Register ID, download BlockChain, start HeartBeat
You can start the program by calling this route(be careful to start only once),
or start the program during bootstrap.
 */
func Start(w http.ResponseWriter, r *http.Request) {

}

/**
GET
Display peerList and sbc
Display the PeerList and the BlockChain. Use the helper function
BlockChain.show() in the starter code to display the BlockChain, and add your own function to display the PeerList.
 */
func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n%s", Peers.Show(), SBC.Show())
}

// Register to TA's server, get an ID
func Register() {}

// Download blockchain from TA server
func Download() {}

/**
/upload
Method: GET
Response: The JSON string of the BlockChain.
Description: Return JSON string of the entire blockchain to the downloader.
Upload blockchain to whoever called this method, return jsonStr
 */
func Upload(w http.ResponseWriter, r *http.Request) {
	blockChainJson, err := SBC.BlockChainToJson()
	if err != nil {
		data.PrintError(err, "Upload")
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
 */
func UploadBlock(w http.ResponseWriter, r *http.Request) {}

/**
/heartbeat/receive
Method: POST
Request: HeartBeatData(see the data structure below)
Description: Receive a heartbeat.
 */
func HeartBeatReceive(w http.ResponseWriter, r *http.Request) {}

// Ask another server to return a block of certain height and hash
func AskForBlock(height int32, hash string) {}

func ForwardHeartBeat(heartBeatData data.HeartBeatData) {}

func StartHeartBeat() {}