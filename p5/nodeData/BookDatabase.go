package nodeData

import (
    "fmt"
    "strings"
    "sync"
)

type BookDatabase struct {
    db map[string]int32
    total int32
    mux sync.Mutex
}

/**
Create a new Book Database
 */
func NewBookDatabase() BookDatabase {
    db := make(map[string]int32)
    return BookDatabase{db, 0, sync.Mutex{}}
}

/**
Show the list of book titles and id's
 */
func (bd *BookDatabase) GetAllBooks() string {
    bd.mux.Lock()
    defer bd.mux.Unlock()
    allBooks := ""
    for key, entry := range bd.db {
        allBooks += "BookTitle: " + key + ", BookId = " + string(entry) + "\n"
    }
    fmt.Println(allBooks)
    return allBooks
}

/**
Get the book id of the title
 */
func (bd *BookDatabase) GetBookId(title string) int32 {
    bd.mux.Lock()
    defer bd.mux.Unlock()
    return bd.db[title]
}

func (bd *BookDatabase) Copy() map[string]int32 {
    bd.mux.Lock()
    defer bd.mux.Unlock()
    return bd.db
}

/**
Add a new book to the BookDatabase
 */
func (bd *BookDatabase) AddBook(title string) int32 {
    bd.mux.Lock()
    defer bd.mux.Unlock()
    titleLower := strings.ToLower(title)

    fmt.Println("book title:", titleLower)
    fmt.Println(bd.db[titleLower])
    id, ok := bd.db[titleLower]
    if ok {
        fmt.Println("Book Exists")
        fmt.Println(bd)
        return id
    } else {
        bd.db[titleLower] = bd.total
        bd.total++
        fmt.Println("Book doesn't exist")
        fmt.Println(bd)
        return bd.total
    }
}

/**
Add Peer book database to own
 */
func (bd *BookDatabase) InjectBookDatabase(books map[string]int32) {

    bd.mux.Lock()
    defer bd.mux.Unlock()

    for title, id := range books {
        bd.db[title] = id
    }
}


