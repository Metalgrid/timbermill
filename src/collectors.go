package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	uuid "github.com/google/uuid"
	syslog "gopkg.in/mcuadros/go-syslog.v2"
)

// MakeSyslogCollector runs a syslog-compatible server that collects messages and stores them in a storage container
func MakeSyslogCollector(Name string, Protocol string, Port uint16) *syslog.Server {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	switch Protocol {
	case "tcp":
		server.ListenTCP(fmt.Sprintf("%s:%d", "0.0.0.0", Port))
		break
	case "udp":
		server.ListenUDP(fmt.Sprintf("%s:%d", "0.0.0.0", Port))
		break
	}
	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			logParts["id"] = uuid.New()
			logParts["timestamp"] = logParts["timestamp"].(time.Time).UTC().Unix()
			data, err := json.Marshal(logParts)
			if err != nil {
				log.Fatalln("Unable to convert syslog message to JSON:", err.Error())
			}

			res, err := http.Post("http://localhost:7700/indexes/syslog/documents", "application/json", strings.NewReader("["+string(data)+"]"))
			if err != nil {
				log.Fatalln("Unable to POST document to meilisearch:", err.Error())
			}
			defer res.Body.Close()

			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatalln("Meilisearch error:", err.Error())
			}
		}
	}(channel)

	go server.Wait()

	return server
}
