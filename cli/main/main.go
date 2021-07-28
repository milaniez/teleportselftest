package main

import (
	"flag"
	"github.com/milaniez/teleportselftest/clientjob"
	"log"
)

func main () {
	var (
		certFile = flag.String(
			"cert",
			"/home/mehdi/mzzcerts/localcerts/server.crt",
			"Server certificate public key file",
		)
		keyFile  = flag.String(
			"key",
			"/home/mehdi/mzzcerts/localkeys/server.key",
			"Server certificate secret key file",
		)
		caDir    = flag.String(
			"cadir",
			"/home/mehdi/mzzcerts/cacerts/",
			"Certificate Authority directory",
		)
		addr     = flag.String(
			"address",
			"localhost:8443",
			"Server address",
		)
	)
	flag.Parse()
	log.SetFlags(log.Llongfile | log.LstdFlags)
	h, err := clientjob.HandleNew(*certFile, *keyFile, *caDir, *addr)
	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	}
	defer func(h *clientjob.Handle) {
		err := clientjob.HandleDel(h)
		if err != nil {
			log.Printf("issue in closure: %v\n", err)
		}
	}(h)
	if id, err := h.GetJobIDs(); err == nil {
		log.Printf("JobIDs: %s\n", id)
	} else {
		log.Printf("Error: %v", err)
	}
	if id, err := h.GetJobIDs(); err == nil {
		log.Printf("JobIDs: %s\n", id)
	} else {
		log.Printf("Error: %v", err)
	}
}