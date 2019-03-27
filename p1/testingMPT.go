package p1

import ("fmt"

)

func TestInsertAndGet() {

    mpt := GetMPTrie()
    mpt.Insert("p", "apple")
    mpt.Insert("aaaaa", "banana")
    mpt.Insert("aaaap", "orange")
    mpt.Insert("aa", "new")
    mpt.Insert("aaaab", "candle")
    mpt.Insert("king", "king")
    mpt.Insert("abc", "alphabet")


    fmt.Println("\nGet test")
    apple := mpt.Get("p")
    banana := mpt.Get("aaaaa")
    orange := mpt.Get("aaaap")
    newWord := mpt.Get("aa")
    candle := mpt.Get("aaaab")
    king := mpt.Get("king")
    alphabet := mpt.Get("abc")
    fmt.Println("Apple:", apple)
    fmt.Println("Banana:", banana)
    fmt.Println("Orange:", orange)
    fmt.Println("New:", newWord)
    fmt.Println("Candle:", candle)
    fmt.Println("King:", king)
    fmt.Println("alphabet:", alphabet)
}

func TestExt1() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("p", "apple")
    mpt.Insert("aa", "banana")
    mpt.Insert("ap", "orange")
    mpt.Insert("b", "new")

    fmt.Println("\nDeleting values")
    deleteIncorrect, err1 := mpt.Delete("c")
    fmt.Println(deleteIncorrect, err1)
    deleteNew, err2 := mpt.Delete("b")
    fmt.Println(deleteNew, err2)
    //deleteOrange, err3 := mpt.Delete("ap")
    //fmt.Println(deleteOrange, err3)
    //deleteBanana, err4 := mpt.Delete("aa")
    //fmt.Println(deleteBanana, err4)
    //deleteApple, err5 := mpt.Delete("p")
    //fmt.Println(deleteApple, err5)

    fmt.Println("\nGet Values")
    banana := mpt.Get("aa")
    fmt.Println("Get:", banana)
    orange := mpt.Get("ap")
    fmt.Println("Get:", orange)
    newWord := mpt.Get("b")
    fmt.Println("Get:", newWord)

    //fmt.Println(mpt.db[mpt.Root])
    //fmt.Println(mpt.db["HashStart_42a990655bffe188c9823a2f914641a32dcbb1b28e8586bd29af291db7dcd4e8_HashEnd"])
    //fmt.Println(mpt.db["HashStart_2fdf6310583baee09f440c41749fd03f2542d1bcb9cf24b78045caf56d77758c_HashEnd"])
    //fmt.Println(mpt.db["HashStart_23ca1c3a6072294f27e66941c8cd3531b5d5ed16d7bf05883b7e30fbf32cb59b_HashEnd"])
    //fmt.Println(mpt.db["HashStart_afb91e31b95ddfc4cc5b179ee86e4ed9d5d5681b0feeb15b21f9564c03749d01_HashEnd"])
    //fmt.Println(mpt.db["HashStart_3c255775632b05b1194107f9ac8b40f9d498720c70536a3f90be2686b31d1b67_HashEnd"])
}

func TestExt2() {
    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("p", "apple")
    mpt.Insert("aa", "banana")
    mpt.Insert("ap", "orange")
    mpt.Insert("ba", "new")

    fmt.Println("\nDeleting values")
    deleteIncorrect, err1 := mpt.Delete("c")
    fmt.Println(deleteIncorrect, err1)
    deleteNew, err2 := mpt.Delete("ba")
    fmt.Println(deleteNew, err2)

    banana := mpt.Get("aa")
    fmt.Println(banana)

    apple := mpt.Get("p")
    fmt.Println(apple)
}




func TestExt3() {
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

}


func TestExt4() {
    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("p", "apple")
    mpt.Insert("aaa", "banana")
    mpt.Insert("aap", "orange")
    mpt.Insert("b", "new")

    fmt.Println("\nDeleting values")
    //deleteNew, err1 := mpt.Delete("b")
    //fmt.Println(deleteNew, err1)

    banana := mpt.Get("aaa")
    fmt.Println("Banana:", banana)

    getNew := mpt.Get("b")
    fmt.Println(getNew)
}

func TestExt7() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("aa", "apple")
    mpt.Insert("ap", "banana")
    mpt.Insert("bc", "new")

    //fmt.Println("\nDeleting values")
    //deleteC, err1 := mpt.Delete("c")
    //fmt.Println(deleteC, err1)

    deleteNew, err2 := mpt.Delete("bc")
    fmt.Println(deleteNew, err2)

    banana := mpt.Get("ap")
    fmt.Println("Banana:", banana)
}

func TestExt8() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("p", "apple")
    mpt.Insert("aaaa", "banana")
    mpt.Insert("aaap", "orange")
    mpt.Insert("a", "new")

    fmt.Println("\nDeleting values")
    deleteC, err1 := mpt.Delete("c")
    fmt.Println(deleteC, err1)

    deleteNew, err2 := mpt.Delete("a")
    fmt.Println(deleteNew, err2)

    apple := mpt.Get("aaaa")
    fmt.Println("Banana:", apple)
}


func TestCharles6InsertAndDelete() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("aaa", "apple")
    mpt.Insert("aap", "banana")
    mpt.Insert("bc", "new")

    fmt.Println("\nDeleting values")
    deleteC, err1 := mpt.Delete("c")
    fmt.Println(deleteC, err1)

}

func TestLeaf1() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("a", "apple")
    mpt.Insert("b", "banana")
    mpt.Insert("a", "new")

    a2 := mpt.Get("a")
    fmt.Println("Get a = new:", a2)

    fmt.Println("\nDeleting values")
    deleteC, err1 := mpt.Delete("c")
    fmt.Println(deleteC, err1)

    deleteNew, err2 := mpt.Delete("a")
    fmt.Println(deleteNew, err2)

}

func TestLeaf2() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("a", "apple")
    mpt.Insert("b", "banana")
    mpt.Insert("ab", "new")

    fmt.Println("\nDeleting values")
    //deleteC, err1 := mpt.Delete("c")
    //fmt.Println(deleteC, err1)

    getNew := mpt.Get("ab")
    fmt.Println("New:", getNew)

    deleteNew, err2 := mpt.Delete("ab")
    fmt.Println(deleteNew, err2)

}

func TestLeaf3() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("a", "apple")
    mpt.Insert("p", "banana")
    mpt.Insert("b", "new")

    deleteNew, err2 := mpt.Delete("b")
    fmt.Println(deleteNew, err2)

    getNew := mpt.Get("b")
    fmt.Println("New:", getNew)

    getApple := mpt.Get("a")
    fmt.Println("Apple:", getApple)

    getBanana := mpt.Get("p")
    fmt.Println("Banana", getBanana)

}

func TestLeaf4() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("a", "apple")
    mpt.Insert("p", "banana")
    mpt.Insert("bc", "new")

    deleteNew, err2 := mpt.Delete("bc")
    fmt.Println(deleteNew, err2)

    getNew := mpt.Get("bc")
    fmt.Println("New:", getNew)

    getApple := mpt.Get("a")
    fmt.Println("Apple:", getApple)

    getBanana := mpt.Get("p")
    fmt.Println("Banana", getBanana)

}

func TestLeaf5() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("bab", "apple")
    mpt.Insert("aa", "banana")
    mpt.Insert("b", "new")

    deleteNew, err2 := mpt.Delete("b")
    fmt.Println(deleteNew, err2)

    getNew := mpt.Get("bc")
    fmt.Println("New:", getNew)
}

func TestLeaf6() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("aab", "apple")
    mpt.Insert("app", "banana")
    mpt.Insert("ac", "new")

    getNew := mpt.Get("ac")
    fmt.Println("New:", getNew)

    deleteNew, err2 := mpt.Delete("ac")
    fmt.Println(deleteNew, err2)

    getNew2 := mpt.Get("ac")
    fmt.Println("New:", getNew2)
}

func TestLeaf7() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("aab", "apple")
    mpt.Insert("app", "banana")
    mpt.Insert("ace", "new")

    getNew := mpt.Get("ace")
    fmt.Println("New:", getNew)

    deleteNew, err2 := mpt.Delete("ace")
    fmt.Println(deleteNew, err2)

    getNew2 := mpt.Get("ace")
    fmt.Println("New:", getNew2)
}

func TestLeaf8() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("p", "banana")
    mpt.Insert("a", "apple")
    mpt.Insert("a", "new")

    a := mpt.Get("a")
    fmt.Println("a:", a)

    deleteNew, err2 := mpt.Delete("ace")
    fmt.Println(deleteNew, err2)

}

func TestLeaf9() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("a", "apple")
    mpt.Insert("p", "banana")
    mpt.Insert("abc", "new")

    apple := mpt.Get("a")
    fmt.Println("apple:", apple)

    banana := mpt.Get("p")
    fmt.Println("banana:", banana)

    getNew := mpt.Get("abc")
    fmt.Println("abc:", getNew)

    deleteNew, err2 := mpt.Delete("abc")
    fmt.Println(deleteNew, err2)

    getNew2 := mpt.Get("abc")
    fmt.Println("abc:", getNew2)
}

func TestLeaf10() {

    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("a", "apple")
    mpt.Insert("b", "new")

    apple := mpt.Get("a")
    fmt.Println("apple:", apple)

    getNew := mpt.Get("b")
    fmt.Println("b:", getNew)

    deleteNew, err2 := mpt.Delete("b")
    fmt.Println(deleteNew, err2)

}
func TestLeaf11() {
    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("a", "apple")
    mpt.Insert("bc", "new")

    deleteDummy, err1 := mpt.Delete("c")
    fmt.Println(deleteDummy, err1)

    deleteNew, err2 := mpt.Delete("bc")
    fmt.Println(deleteNew, err2)

}

func TestLeaf12() {
    mpt := GetMPTrie()
    fmt.Println("Inserting values")
    mpt.Insert("ap", "apple")
    mpt.Insert("b", "new")

    deleteDummy, err1 := mpt.Delete("c")
    fmt.Println(deleteDummy, err1)

    deleteNew, err2 := mpt.Delete("b")
    fmt.Println(deleteNew, err2)

}

func TestBranch1() {

    mpt := GetMPTrie()
    mpt.Insert("aa", "apple")
    mpt.Insert("ap", "banana")
    mpt.Insert("a", "new")

    deleteNew, err2 := mpt.Delete("a")
    fmt.Println(deleteNew, err2)
}

func TestBranch2() {

    mpt := GetMPTrie()
    mpt.Insert("a", "old")
    mpt.Insert("aa", "apple")
    mpt.Insert("ap", "banana")
    mpt.Insert("a", "new")

    deleteNew, err2 := mpt.Delete("a")
    fmt.Println(deleteNew, err2)

    getNew := mpt.Get("aa")
    fmt.Println("a:", getNew)

}

func TestBranch3() {
    mpt := GetMPTrie()
    mpt.Insert("a", "apple")
    mpt.Insert("b", "banana")
    mpt.Insert("c", "new")


    deleteNew, err2 := mpt.Delete("cc")
    fmt.Println(deleteNew, err2)

    getNew := mpt.Get("b")
    fmt.Println("c:", getNew)
}

func TestBranch4() {
    mpt := GetMPTrie()
    mpt.Insert("aa", "apple")
    mpt.Insert("ap", "banana")
    mpt.Insert("a", "old")
    mpt.Insert("aA", "new")

    deleteNew, err2 := mpt.Delete("aA")
    fmt.Println(deleteNew, err2)

    getNew := mpt.Get("a")
    fmt.Println("c:", getNew)

}




