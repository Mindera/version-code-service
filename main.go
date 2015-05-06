package main

import (
	"flag"
	"fmt"
	"github.com/mindera/version-code-service/vccounter"
	"log"
	"net/http"
)

const (
	mgoDatabaseName = "versionCodesDB"
)

func main() {

	serverPortPtr := flag.Int("port", 8080, "Service port number to listen for requests")
	flag.Parse()

	avcDataStore := vccounter.AppVersionCodeMGODataStore{}
	avcDataStore.OpenConnection(fmt.Sprintf("localhost/%s", mgoDatabaseName))
	defer avcDataStore.CloseConnection()

	vc := vccounter.VersionCodeService{&avcDataStore}
	vc.Register()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *serverPortPtr), nil))
}
