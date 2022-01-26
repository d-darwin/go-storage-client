package main

import (
	"fmt"
	"log"

	"github.com/d-darwin/go-storage/pkg/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Go modules are awesome!"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)
}
