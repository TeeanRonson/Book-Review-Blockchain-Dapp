package data

import (
    "fmt"
    "reflect"
    "testing"
)

func TestRebalance(t *testing.T) {

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
    if !reflect.DeepEqual(peers, expected) {
        t.Fail()
    }

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
    if !reflect.DeepEqual(peers, expected) {
       t.Fail()
    }

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
    if !reflect.DeepEqual(peers, expected) {
       t.Fail()
    }

}