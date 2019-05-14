package nodeData

import (
    "fmt"
    "strings"
)

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

func (bd *BookDatabase) AddBook(title string) int32 {

    titleLower := strings.ToLower(title)

    fmt.Println("book title:", titleLower)
    fmt.Println(bd.db[titleLower])
    id, ok := bd.db[titleLower]
    if ok {
        fmt.Println("Book Database0")
        fmt.Println(bd)
        fmt.Println(bd.total)
        return id
    } else {
        bd.db[titleLower] = bd.total
        bd.total++
        fmt.Println("Book Database")
        fmt.Println(bd)
        fmt.Println(bd.total)
        return bd.total
    }


    return 0
}


