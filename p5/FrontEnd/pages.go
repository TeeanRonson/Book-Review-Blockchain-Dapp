package FrontEnd

import "os"

//var bookReviewsHead =  "<h2>Reviews</h2>\n" +
//    "\n" +
//    "<table>\n" +
//    "<tr>\n" +
//    "<th></th>\n" +
//    "<th>Review</th>\n" +
//    "</tr>\n"
//
//var bookReviewsFoot = "</table> \n"


//func BookReviewsHead() string {
//    return bookReviewsHead
//}
//
//func BookReviewsFoot() string {
//    return bookReviewsFoot
//}


/**
Returns the interface for creating a new book review
 */
func CreateBookReview() string {

    var newBookReview = "<h2>New Book Review</h2> \n" + "<form action=\"/newBookReview\" method=\"post\">\n" +
        "Title:<br/>\n" + "<input type=\"text\" name=\"title\"/>\n" + "<br/>\n" +
        "Review:<br/>\n" + "<input type=\"text\" name=\"reviewText\"/>\n" + "<br/>\n" +
        "Rating:<br/>\n" + "<input type=\"text\" name=\"rating\"/>\n" + "<br/>\n" +
        "Transaction Fee:<br/>\n" + "<input type=\"text\" name=\"txFee\"/>\n" + "<br/>\n" +
        "Public Key:<br/>\n" + "<input type=\"text\" name=\"pubKey\" value=\"" + os.Args[1] + "\"/>\n" + "<br/>\n" +
        "Private Key:<br/>\n" + "<input type=\"text\" name=\"priKey\"/>\n" + "<br/>\n" +
        "<input type=\"submit\" value=\"Submit\"/>\n" +
        "</form>\n"
    return newBookReview
}

func Header() string {

    header := "<html>\n" + "<head>\n" + "<style>\n" +
        "table {\n" +
        "    font-family: arial, sans-serif;\n" +
        "    border-collapse: collapse;\n" +
        "    width: 100%;\n" +
        "}\n" +
        "\n" +
        "td, th {\n" +
        "    border: 1px solid #dddddd;\n" +
        "    text-align: left;\n" +
        "    padding: 8px;\n" +
        "}\n" +
        "\n" +
        "tr:nth-child(even) {\n" +
        "    background-color: #dddddd;\n" +
        "}\n" + "</style>\n" + "</head>\n + <body><center>"
    return header
}

/**
Return the end of the html page
 */
func End() string {
    end := "</body></html>"
    return end
}

/**
Footer of a page
 */
func Footer() string {
    footer := "<footer><center> Copyright &copy; The Book Review Company </center></footer>"
    return footer
}

/**
Book of the month
 */
func BookOfTheMonth() string {
    botm := "<img src=\"http://static1.squarespace.com/static/562a56dde4b026a61c55807c/57e38a1e5016e12f880dfeed/5b60003b2b6a28ea73ca1930/1533650606212/Made+to+Stick+Thumbnail.jpg?format=1500w\" height=\"200\" width=\"300\" style=\"float:middle;\"/>"
    return botm
}

/**
Confirmation of a new book review page
 */
func Confirmation(title string, review string, rating string, signature string) string {

    var page string
    var successfulPost string
    successfulPost = "<h2>Your book review was submitted successfully!</h2>\n <h4>Your confirmation</h4>\n"
    successfulPost += "Title: " + title + "<br/>\n"
    successfulPost += "Review: " + review+ "<br/>\n"
    successfulPost += "Your rating: " + rating + "<br/>\n"
    successfulPost += "Signed by: " + signature + "<br/>\n"
    page += Header() + BookOfTheMonth() + successfulPost + CreateBookReview() + Footer() + End()
    return page
}

/**
Write a new book review page
 */
func WriteANewBookReview() string {
    var page string
    page += Header() + BookOfTheMonth() + CreateBookReview() + Footer() + End()
    return page
}
