package data

import (
    "encoding/json"
    "fmt"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p5/nodeData"
)

/**
Return all the book reviews in every block of the block chain as a Json
 */
func (sbc *SyncBlockChain) GetAllReviewsHelper() string {

    block := sbc.SyncGetLatestBlocks()[0]
    reviews := ""

    for block.Header.Height >= 1 {
        for key, value := range block.Mpt.Inputs {
            reviewData, _ := DecodeJsonToReviewData(value)
            fmt.Println(reviewData)
            reviews += "<tr>\n" +
                "<td>" + key + "</td>\n" +
                "<td>" +
                "BookId: " + Int32ToString(reviewData.BookId) + " Review: " + reviewData.ReviewText + " Rating: " + Int32ToString(reviewData.ReviewRating) +
                "</td>\n" +
                "</tr>"
        }
        parent := sbc.SyncGetParentBlock(block)
        block = parent
    }
    return reviews
}

/**
Decode the JsonString into ReviewData
 */
func DecodeJsonToReviewData(jsonString string) (nodeData.ReviewData, error) {

    reviewData := nodeData.ReviewData{}
    if err := json.Unmarshal([]byte(jsonString), &reviewData); err != nil {
        fmt.Println("Can't Unmarshal in decodeReviewData")
        panic(err)
        return reviewData, err
    }
    return reviewData, nil
}

func Int32ToString(n int32) string {
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
