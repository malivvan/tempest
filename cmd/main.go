package main

import (
	"fmt"
	"time"

	"github.com/malivvan/tempest"
)

func main() {
	db, err := tempest.Open("store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m1, err := db.GetMeta([]byte("foo"))
	if err == nil {
		fmt.Println("m1:", string(m1))

	}
	err = db.SetMeta([]byte("foo"), []byte(time.Now().String()))
	if err != nil {
		panic(err)
	}
	m2, err := db.GetMeta([]byte("foo"))
	if err != nil {
		panic(err)
	}
	fmt.Println("m2:", string(m2))

}
