package ClientNode
import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Show",
        "GET",
        "/show",
        Show,
    },
    Route{
        "UploadBlock",
        "GET",
        "/block/{height}/{hash}",
        UploadBlock,
    },
    Route{
        "HeartBeatReceive",
        "POST",
        "/heartbeat/receive",
        HeartBeatReceive,
    },
    Route{
        "Start",
        "GET",
        "/start",
        Start,
    },
    Route{
        "Canonical",
        "GET",
        "/canonical",
        Canonical,
    },
    Route{
        "AllBookReviews",
        "GET",
        "/getAllBookReviews",
        GetAllBookReviews,
    },
    Route{
        "NewBookReview",
        "GET",
        "/newBookReview",
        NewBookReview,
    },
    Route{
        "NewBookReview",
        "POST",
        "/newBookReview",
        NewBookReview,
    },
    Route{
        "ReviewObjectReceive",
        "POST",
        "/reviewData/receive",
        ReviewObjectReceive,
    },
}