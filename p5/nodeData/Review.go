package nodeData

type ReviewData struct {
    Title string        `json:"title"`
    ReviewText string   `json:"reviewText"`
    ReviewRating int32  `json:"reviewRating"`
    TxFee float32       `json:"txFee"`
    PublicKey string    `json:"publicKey"`
    Signature string    `json:"signature"`
    BookId int32        `json:"bookId"`
}

type ReviewObject struct {
    Title string `json:"title"`
    ReviewText string `json:"reviewText"`
    ReviewRating string `json:"reviewRating"`
    BookId int32 `json:"bookId"`
}

/**
Prepare a new Review Data object
 */
func prepareReviewData(title string, reviewText string, reviewRating int32, txFee float32, publicKey string, signature string) ReviewData {

    return ReviewData{title, reviewText, reviewRating, txFee, publicKey, signature, 1}

}



