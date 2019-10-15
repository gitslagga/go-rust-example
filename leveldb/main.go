package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	levelDB()
}
func levelDB() {
	key := []byte("hello")
	value := []byte("hi i'm levelDB by go")

	k1 := "foo"
	v1 := "barbarbarbarbar"

	// The returned DB instance is safe for concurrent use.
	// The DB must be closed after use, by calling Close method.
	db, err := leveldb.OpenFile("./leveldb/leveldb-data", nil)

	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	e := db.Put(key, value, nil)
	e = db.Put([]byte(k1), []byte(v1), nil)
	fmt.Println(e) //<nil>

	data, _ := db.Get(key, nil)
	fmt.Println(data)        //[104 105 32 105 39 109 32 108 101 118 101 108 68 66 32 98 121 32 103 111]
	fmt.Printf("%s\n", data) //hi i'm levelDB by go
	data, _ = db.Get([]byte(k1), nil)
	fmt.Println(data)        // [98 97 114 98 97 114 98 97 114 98 97 114 98 97 114]
	fmt.Printf("%s\n", data) //barbarbarbarbar

	i := db.Delete(key, nil)
	fmt.Println(i) // <nil>
	data, _ = db.Get(key, nil)
	fmt.Println(data) // []

	snapshot, i := db.GetSnapshot()
	fmt.Println(snapshot) // leveldb.Snapshot{22}
	fmt.Println(i)        // <nil>

	//The snapshot must be released after use, by calling Release method.
	snapshot.Release()
}
