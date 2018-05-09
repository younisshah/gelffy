package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/younisshah/gelffy/service"
	"github.com/younisshah/gelffy/upload"
	"github.com/younisshah/gelffy/listener"
	"os"
)

const GelfPort = "12123"

func main() {

	customerToken := os.Getenv("LOGGLY_CUSTOMER_TOKEN")

	if len(customerToken) == 0 {
		log.Errorln("failed to read Loggly customer token")
		log.Fatalln("make sure to specify the token as an ENV variable")
	}
	log.Infoln("starting gelffy server on port: " + GelfPort)

	udpConn := service.StartUDPServer(GelfPort)
	defer udpConn.Close()

	logChan := make(chan []byte)
	forever := make(chan struct{})

	go upload.StartUpload(customerToken, logChan)

	go listener.ListenForLogs(udpConn, logChan)

	log.Infoln("gelffy server started")

	<-forever // run forever
}
