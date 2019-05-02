package nodeData



type BookDatabase struct {
    db map[string]int32
    total int32
}

/**
Create a new Book Database
 */
func NewBookDatabase() BookDatabase {
    db := make(map[string]int32)

    return BookDatabase{db, 0}
}

/**
Show the list of book titles and id's
 */
func (reviews *BookDatabase) Show() string {
    var result string
    for key, entry := range reviews.db {
        result += "BookTitle: " + key + ", BookId = " + string(entry) + "\n"
    }
    return result
}

/**
Get the book id of the title
 */
func (reviews *BookDatabase) getBookId(title string) int32 {
    return reviews.db[title]
}

