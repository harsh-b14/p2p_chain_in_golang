package storage

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

func StartDataBase() {
  // Open the Badger database located in the /tmp/badger directory.
  // It is created if it doesn't exist.
  db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
  if err != nil {
    log.Fatal(err)
  }

	fmt.Println("✅ Database started successfully")

  defer db.Close()

  // your code here
}