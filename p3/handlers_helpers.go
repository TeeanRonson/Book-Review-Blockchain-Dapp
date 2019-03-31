package p3

import (
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p3/data"
    "strconv"
)

func convertToInt32(value string) int32 {

    i, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
       fmt.Println("Unable to convert string to int")
       panic(err)
    }
    return int32(i)
    //i1, err := strconv.Atoi(value)
    //if err != nil {
    //    fmt.Println(i1)
    //    panic(err)
    //}
    //return int32(i1)
}

func int32ToString(n int32) string {
    buf := [11]byte{}
    pos := len(buf)
    i := int64(n)
    signed := i < 0
    if signed {
        i = -i
    }
    for {
        pos--
        buf[pos], i = '0'+byte(i%10), i/10
        if i == 0 {
            if signed {
                pos--
                buf[pos] = '-'
            }
            return string(buf[pos:])
        }
    }
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


