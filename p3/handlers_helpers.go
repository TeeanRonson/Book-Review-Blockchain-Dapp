package p3

import (
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3/data"
    "strconv"
)

func convertToInt32(value string) int32 {

    i, err := strconv.ParseInt(value, 10, 32)
    if err != nil {
        fmt.Println("Unable to convert string to int")
        panic(err)
    }

    return int32(i)
}

/**
Decode the JsonString into HeartBeatData
 */
func decodeJsonToHbd(hbd string) data.HeartBeatData {

    recvHeartBeatData := data.HeartBeatData{}
    if err := json.Unmarshal([]byte(hbd), &recvHeartBeatData); err != nil {
        fmt.Println("Can't Unmarshal in decodeHBD")
        return recvHeartBeatData
    }
    return recvHeartBeatData
}

/**
Decode the JsonString into HeartBeatData
 */
func decodeJsonToHbdMod(hbd string) data.HeartBeatDataMod {

    recvHeartBeatData := data.HeartBeatDataMod{}
    if err := json.Unmarshal([]byte(hbd), &recvHeartBeatData); err != nil {
        fmt.Println("Can't Unmarshal in decodeHBD")
        return recvHeartBeatData
    }
    return recvHeartBeatData
}


