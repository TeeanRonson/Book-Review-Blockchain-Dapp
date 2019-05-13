package nodeData

import "strings"

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
func (bd *BookDatabase) Show() string {
    var result string
    for key, entry := range bd.db {
        result += "BookTitle: " + key + ", BookId = " + string(entry) + "\n"
    }
    return result
}

/**
Get the book id of the title
 */
func (bd *BookDatabase) GetBookId(title string) int32 {
    return bd.db[title]
}

func (bd *BookDatabase) AddBook(title string) {

    titleLower := strings.ToLower(title)
    if _, ok := bd.db[titleLower]; ok {
        if ok == false {
            bd.db[titleLower] = bd.total + 1
            bd.total++
        }
    }
}


