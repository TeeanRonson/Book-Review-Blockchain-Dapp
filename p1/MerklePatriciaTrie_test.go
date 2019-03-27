package p1

import (
    "fmt"
    "testing"
)

func TestMerklePatriciaTrie_Get1(t *testing.T) {

    mpt := GetMPTrie()
    mpt.Insert("p", "apple")
    mpt.Insert("aaaaa", "banana")
    mpt.Insert("aaaap", "orange")
    mpt.Insert("aa", "new")
    mpt.Insert("aaaab", "candle")
    mpt.Insert("king", "king")
    mpt.Insert("abc", "alphabet")

    apple := mpt.Get("p")
    banana := mpt.Get("aaaaa")
    orange := mpt.Get("aaaap")
    newWord := mpt.Get("aa")
    candle := mpt.Get("aaaab")
    king := mpt.Get("king")
    alphabet := mpt.Get("abc")
    if  apple != "apple" || banana != "banana" || orange != "orange" || newWord != "new" || candle != "candle" || king != "king" || alphabet != "alphabet"{
        t.Errorf("Result is %s", apple)
        t.Errorf("Result is %s", banana)
        t.Errorf("Result is %s", orange)
        t.Errorf("Result is %s", newWord)
        t.Errorf("Result is %s", candle)
        t.Errorf("Result is %s", king)
        t.Errorf("Result is %s", alphabet)
    }
}

func TestMerklePatriciaTrie_Delete(t *testing.T) {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("aaa", "apple")
    mpt.Insert("aap", "banana")
    mpt.Insert("bb", "right leaf")
    mpt.Insert("aa", "new")

    fmt.Println("\nDeleting values")
    deleteNew, err1 := mpt.Delete("aa")
    fmt.Println(deleteNew, err1)

    banana := mpt.Get("aap")
    fmt.Println("Banana:", banana)

    getApple := mpt.Get("aaa")
    fmt.Println("Apple:", getApple)

    if deleteNew != "Successful Deletion" || banana != "banana" || getApple != "apple" {
        t.Errorf("deleteNew %s", deleteNew)
        t.Errorf("banana %s", banana)
        t.Errorf("apple %s", getApple)
    }

}
