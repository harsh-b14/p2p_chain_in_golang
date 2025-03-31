package initialize

// import (
// 	"log"

// 	db "github.com/harsh-b14/p2p-chain/storage"
// )

// func GlobalDbVar() {
// 	err := db.BadgerDB.View(db.Get([]byte("stateRootHash"), &t.StateRootHash))
// 	if err != nil {
// 		log.Default().Printf("Error in init StateRootHash\n")
// 		log.Default().Println(err.Error())
// 	} else {
// 		log.Default().Printf("StateRootHash: %x\n", t.StateRootHash)
// 	}
// 	err = db.BadgerDB.View(db.Get([]byte("latestBlock"), &t.LatestBlock))
// 	if err != nil {
// 		log.Default().Printf("Error in init latestBlock: \n")
// 		log.Default().Println(err.Error())
// 	} else {
// 		log.Default().Printf("LatestBlock: %x\n", t.LatestBlock)
// 	}
// }
