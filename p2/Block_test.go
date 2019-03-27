package p2

import (
    "fmt"
    "reflect"
    "testing"
)

func TestDecodeEncode(t *testing.T) {

    message := "TestDecodeFromJson error"
    test1 := "{\"height\":1,\"timeStamp\":1234567890,\"hash\":\"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\",\"parentHash\":\"genesis\",\"size\":1174,\"mpt\":{\"charles\":\"ge\",\"hello\":\"world\"}}"
    block, err := DecodeFromJson(test1)
    if err != nil {
        fmt.Println(err)
        t.Errorf("Error at DecodeFromJson: %s", message)

    }

    compare, err := block.EncodeToJson()

    if err != nil {
        fmt.Println(err)
        t.Errorf("Error at EncodeToJson")
    }

    if !reflect.DeepEqual(compare, test1) {
        t.Errorf("Incorrect output match: %s", message)
    }
}


