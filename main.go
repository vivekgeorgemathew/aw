package main

import (
	"github.com/vivekgeorgemathew/aw/api"
	"github.com/vivekgeorgemathew/aw/db/store"
	"github.com/vivekgeorgemathew/aw/util"
	"log"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read the configuration : ", err)
	}
	dbStore := store.NewStore()
	server := api.NewServer(dbStore)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start the Server: ", err)
	}
}
