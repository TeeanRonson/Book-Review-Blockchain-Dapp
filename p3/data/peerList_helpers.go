package data

import (
    "sort"
)

// A data structure to hold key/value pairs
//type Pair struct {
//    Key   string
//    Value int32
//}

type PeerId struct {
    Value int32
}

// A slice of pairs that implements sort.Interface to sort by values
type PeerListId []PeerId

func (p PeerListId) Len() int           { return len(p) }
func (p PeerListId) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PeerListId) Less(i, j int) bool { return p[i].Value < p[j].Value }


func RebalanceHelper(peerList map[string]int32) PeerListId {

    list := make(PeerListId, len(peerList))

    i := 0
    for _, value := range peerList {
        list[i] = PeerId{value}
        i++
    }

    //fmt.Printf("Pre-sorted: ")
    //for _, k := range p {
    //    fmt.Printf("%s ", k.Key)
    //}
    //fmt.Println("")
    sort.Sort(list)

    //fmt.Printf("Post-sorted: ")
    //for _, k := range pairs {
    //    fmt.Printf("%s ", k.Key)
    //}
    //fmt.Println("")
    return list

}