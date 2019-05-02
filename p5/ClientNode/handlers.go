package ClientNode

import (
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p2"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p5/nodeData"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3/data"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "reflect"
    "strings"
    "time"
)

var FIRST_NODE = "http://localhost:6686"
var GET_BC_FIRST_NODE = FIRST_NODE + "/upload"
var BookDatabase nodeData.BookDatabase
var HTTP = "http://"
var LOCAL_HOST = "localhost:"
var HTTPLOCALHOST = HTTP + LOCAL_HOST

var SBC data.SyncBlockChain
var Peers data.PeerList


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

    fmt.Println("Fetching the blockchain")
    Download()
    go StartHeartBeat()

}

/**
Create SyncBlockChain and PeerList instances
ifStarted indicates if the block can start sending and receiving blocks
i.e. after it has already downloaded the block chain
 */
func Init() {
    p3.SBC = data.NewBlockChain()
    p3.Peers = data.NewPeerList(p3.ConvertToInt32(os.Args[1]), 32)
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
Get all book reviews
 */
func GetAllBookReviews(w http.ResponseWriter, r *http.Request) {



}

/**
New Book Review
 */
func NewBookReview(w http.ResponseWriter, r *http.Request) {



}

/**
New Book Review
 */
func CreateBookReview(w http.ResponseWriter, r *http.Request) {



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
    p3.SBC.UpdateEntireBlockChain(string(currBlockChain))
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

    var theBlock p2.Block
    path := strings.Split(r.URL.Path, "/")
    height := path[2]
    targetHash := path[3]

    //Check for errors
    if path[1] != "block" || len(path) != 4 || path == nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }
    theBlock = SBC.GetBlock(p3.ConvertToInt32(height), targetHash)

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

    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    //Convert Json into HeartBeatData
    hbd := p3.DecodeJsonToHbd(string(body))
    Peers.InjectPeerMapJson(hbd.PeerMapJson, hbd.Addr, hbd.Id)

    //if HeartBeatData has a new block
    if hbd.IfNewBlock {
        //Get the new block
        recvBlock, _ := p2.DecodeFromJson(hbd.BlockJson)

        if p3.VerifyProofOfWork(recvBlock) {
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
            hbd.Id = p3.ConvertToInt32(os.Args[1])
            if hbd.Hops > 0 {
                ForwardHeartBeat(hbd)
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

    for addr, _ := range Peers.Copy() {
        endPoint := HTTPLOCALHOST + addr + "/block/" + p3.Int32ToString(height) + "/"+ hash
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

    return parent
}