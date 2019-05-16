package nodeData

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "sync"
)

var bookId int32
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
    bookId = ConvertToInt32(os.Args[3])
    return BookDatabase{db, 0, sync.Mutex{}}
}

func ConvertToInt32(value string) int32 {

    i, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
        fmt.Println("Unable to convert string to int")
        panic(err)
    }
    return int32(i)
}

/**
Show the list of book titles and id's
 */
func (bd *BookDatabase) GetAllBooks() string {
    bd.mux.Lock()
    defer bd.mux.Unlock()
    allBooks := ""
    for key, entry := range bd.db {
        //fmt.Println(entry)
        allBooks += "BookTitle: " + key + ", BookId = " + convertInt32ToString(entry) +  "\n"
    }
    return allBooks
}


func convertInt32ToString(n int32) string {
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
        fmt.Println("Book doesn't exist")
        bd.db[titleLower] = bookId
        bookId += 2
        bd.total += 1
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
        bd.total++
    }
}


